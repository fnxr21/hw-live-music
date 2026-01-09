package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(user models.RefUser) (*models.RefUser, error)
	GetUserByID(id string) (*models.RefUser, error)
	ListUsers() ([]*models.RefUser, error)
	UpdateUser(user models.RefUser) (*models.RefUser, error)
	DeleteUser(id string) error
}

// CreateUser creates a new user, returns error if name exists
func (r *repository) CreateUser(user models.RefUser) (*models.RefUser, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var existing models.RefUser
	if err := tx.Where("name = ?", user.Name).Take(&existing).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("user with this name already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// GetUserByID fetches a user by UUID
func (r *repository) GetUserByID(id string) (*models.RefUser, error) {
	var user models.RefUser
	err := r.db.Where("user_id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// ListUsers returns all users
func (r *repository) ListUsers() ([]*models.RefUser, error) {
	var users []*models.RefUser
	if err := r.db.Where("is_active = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user
func (r *repository) UpdateUser(user models.RefUser) (*models.RefUser, error) {
	user.UpdatedAt = time.Now()

	fmt.Println("Updating user with ID:", user)
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser deletes a user by UUID
func (r *repository) DeleteUser(id string) error {
	// return r.db.Delete(&models.RefUser{}, "user_id = ?", id).Error
	return r.db.Model(&models.RefUser{}).
		Where("user_id = ?", id).
		Update("is_active", false).Error
}
