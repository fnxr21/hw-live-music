package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"gorm.io/gorm"
)

type Song interface {
	CreateSong(song models.RefSong) (*models.RefSong, error)
	GetSongByID(id string) (*models.RefSong, error)
	ListSongs() ([]*models.RefSong, error)
	UpdateSong(song models.RefSong) (*models.RefSong, error)
	DeleteSong(id string) error
}

// CreateSong creates a new song, returns error if a song with the same title and artist exists
func (r *repository) CreateSong(song models.RefSong) (*models.RefSong, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Check if a song with the same title and artist already exists
	var existing models.RefSong
	if err := tx.Where("title = ? AND artist = ?", song.Title, song.Artist).Take(&existing).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("song with this title and artist already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// Create the new song
	if err := tx.Create(&song).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &song, nil
}

// GetSongByID fetches a song by UUID, only if active
func (r *repository) GetSongByID(id string) (*models.RefSong, error) {
	var song models.RefSong
	err := r.db.Where("song_id = ? AND is_active = ?", id, true).First(&song).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &song, err
}

// ListSongs returns all active songs
func (r *repository) ListSongs() ([]*models.RefSong, error) {
	var songs []*models.RefSong
	if err := r.db.Where("is_active = ?", true).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

// UpdateSong updates an existing song
func (r *repository) UpdateSong(song models.RefSong) (*models.RefSong, error) {
	song.UpdatedAt = time.Now()

	fmt.Println("Updating song with ID:", song.SongID)
	if err := r.db.Save(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

// DeleteSong performs a soft delete by setting is_active = false
func (r *repository) DeleteSong(id string) error {
	result := r.db.Model(&models.RefSong{}).
		Where("song_id = ?", id).
		Update("is_active", false)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
