package repositories

import (
	"errors"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"gorm.io/gorm"
)

type LivePlaylist interface {
	CreateLivePlaylist(playlist models.TrxLivePlaylist) (*models.TrxLivePlaylist, error)
	GetLivePlaylistByID(id string) (*models.TrxLivePlaylist, error)
	ListLivePlaylists() ([]*models.TrxLivePlaylist, error)
	UpdateLivePlaylist(playlist models.TrxLivePlaylist) (*models.TrxLivePlaylist, error)
	DeleteLivePlaylist(id string) error
}

// CreateLivePlaylist creates a new playlist entry
func (r *repository) CreateLivePlaylist(playlist models.TrxLivePlaylist) (*models.TrxLivePlaylist, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(&playlist).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &playlist, nil
}

// GetLivePlaylistByID fetches a playlist entry by UUID
func (r *repository) GetLivePlaylistByID(id string) (*models.TrxLivePlaylist, error) {
	var playlist models.TrxLivePlaylist
	err := r.db.Where("live_playlist_id = ? AND is_active = ?", id, true).First(&playlist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &playlist, err
}

// ListLivePlaylists returns all active playlist entries
func (r *repository) ListLivePlaylists() ([]*models.TrxLivePlaylist, error) {
	var playlists []*models.TrxLivePlaylist
	if err := r.db.Where("is_active = ?", true).Find(&playlists).Error; err != nil {
		return nil, err
	}
	return playlists, nil
}

// UpdateLivePlaylist updates an existing playlist entry
func (r *repository) UpdateLivePlaylist(playlist models.TrxLivePlaylist) (*models.TrxLivePlaylist, error) {
	playlist.UpdatedAt = time.Now()
	if err := r.db.Save(&playlist).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

// DeleteLivePlaylist performs a soft delete
func (r *repository) DeleteLivePlaylist(id string) error {
	return r.db.Model(&models.TrxLivePlaylist{}).
		Where("live_playlist_id = ?", id).
		Update("is_active", false).Error
}
