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
	ListLivePlaylists(limit, offset int) ([]*LivePlaylistWithDetails, int64, error)

	UpdateLivePlaylist(playlist models.TrxLivePlaylist) (*models.TrxLivePlaylist, error)
	DeleteLivePlaylist(id string) error

	RealTimeListLivePlaylists() ([]*models.TrxLivePlaylist, error)
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
// ListLivePlaylists returns all active playlist entries with pagination
// func (r *repository) RealTimeListLivePlaylists() ([]*models.TrxLivePlaylist, error) {
// 	var playlists []*models.TrxLivePlaylist

// 	query := `
// 		SELECT *
// 		FROM trx_live_playlists
// 		WHERE is_active = true
// 		LIMIT 20
// 	`
// 	if err := r.db.Raw(query).Scan(&playlists).Error; err != nil {
// 		return nil, err
// 	}

// 	return playlists, nil
// }

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


func (r *repository) RealTimeListLivePlaylists() ([]*models.TrxLivePlaylist, error) {
	var playlists []*models.TrxLivePlaylist

	query := `
		SELECT *
		FROM trx_live_playlists
		WHERE is_active = true
		ORDER BY order_number ASC
		LIMIT 20
	`
	if err := r.db.Raw(query).Scan(&playlists).Error; err != nil {
		return nil, err
	}

	return playlists, nil
}

func (r *repository) ListLivePlaylists(limit, offset int) ([]*LivePlaylistWithDetails, int64, error) {
	var playlists []*LivePlaylistWithDetails
	var total int64

	query := `
		SELECT
			tlp.*,
			rs.title,
			rs.artist,
			rs.duration,
			rs.header_image_url,
			rs.release_song_date,
			rs.url,
			rt.table_number
		FROM live_music.trx_live_playlists tlp
		LEFT JOIN live_music.trx_song_requests tsr
			ON tsr.song_request_id = tlp.song_request_id
			AND tsr.is_active = true
		LEFT JOIN live_music.ref_songs rs
			ON rs.song_id = tsr.song_id
			AND rs.is_active = true
		LEFT JOIN live_music.ref_tables rt
			ON rt.table_id = tlp.table_id
			AND rt.is_active = true
		WHERE tlp.is_active = true
		ORDER BY tlp.order_number ASC
		LIMIT ? OFFSET ?;
	`

	if err := r.db.Raw(query, limit, offset).Scan(&playlists).Error; err != nil {
		return nil, 0, err
	}

	// count total active playlists
	if err := r.db.Model(&models.TrxLivePlaylist{}).
		Where("is_active = ?", true).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return playlists, total, nil
}

type LivePlaylistWithDetails struct {
	models.TrxLivePlaylist            // embed your original playlist struct
	Title                  *string    `gorm:"column:title"`
	Artist                 *string    `gorm:"column:artist"`
	Duration               *int       `gorm:"column:duration"`
	HeaderImageURL         *string    `gorm:"column:header_image_url"`
	ReleaseSongDate        *time.Time `gorm:"column:release_song_date"`
	URL                    *string    `gorm:"column:url"`
	TableNumber            *int       `gorm:"column:table_number"`
}
