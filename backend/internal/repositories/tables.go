package repositories

import (
	"errors"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"gorm.io/gorm"
)

type Table interface {
	CreateTable(table models.RefTable) (*models.RefTable, error)
	GetTableByID(id string) (*models.RefTable, error)
	ListTables() ([]*models.RefTable, error)
	UpdateTable(table models.RefTable) (*models.RefTable, error)
	DeleteTable(id string) error
}

// CreateTable creates a new table, ensuring table number is unique
func (r *repository) CreateTable(table models.RefTable) (*models.RefTable, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Check if table number exists
	var existing models.RefTable
	if err := tx.Where("table_number = ?", table.TableNumber).Take(&existing).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("table number already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&table).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &table, nil
}

// GetTableByID fetches a table by UUID
func (r *repository) GetTableByID(id string) (*models.RefTable, error) {
	var table models.RefTable
	err := r.db.Where("table_id = ? AND is_active = ?", id, true).First(&table).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &table, err
}

// ListTables returns all active tables
func (r *repository) ListTables() ([]*models.RefTable, error) {
	var tables []*models.RefTable
	if err := r.db.Where("is_active = ?", true).Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// UpdateTable updates an existing table
func (r *repository) UpdateTable(table models.RefTable) (*models.RefTable, error) {
	table.UpdatedAt = time.Now()
	if err := r.db.Save(&table).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

// DeleteTable performs a soft delete
func (r *repository) DeleteTable(id string) error {
	return r.db.Model(&models.RefTable{}).
		Where("table_id = ?", id).
		Update("is_active", false).Error
}
