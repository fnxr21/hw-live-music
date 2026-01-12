package repositories

import (
	"errors"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"gorm.io/gorm"
)

type SongRequest interface {
	CreateSongRequest(req models.TrxSongRequest) (*models.TrxSongRequest, error)
	GetSongRequestByID(id string) (*models.TrxSongRequest, error)
	ListSongRequests(limit, offset int) ([]*SongRequestWithDetails, int64, error)
	UpdateSongRequest(req models.TrxSongRequest) (*models.TrxSongRequest, error)
	DeleteSongRequest(id string) error
	GetSongRequestByIDTable(id int) ([]*models.TrxSongRequest, error)
}

// CreateSongRequest inserts a new song request
func (r *repository) CreateSongRequest(req models.TrxSongRequest) (*models.TrxSongRequest, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(&req).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &req, nil
}

// GetSongRequestByID fetches a song request by UUID
func (r *repository) GetSongRequestByID(id string) (*models.TrxSongRequest, error) {
	var req models.TrxSongRequest
	err := r.db.Where("song_request_id = ? AND is_active = ?", id, true).First(&req).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &req, err
}

// UpdateSongRequest updates an existing song request
func (r *repository) UpdateSongRequest(req models.TrxSongRequest) (*models.TrxSongRequest, error) {
	req.UpdatedAt = time.Now()
	if err := r.db.Save(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

// DeleteSongRequest soft deletes a song request
func (r *repository) DeleteSongRequest(id string) error {
	return r.db.Model(&models.TrxSongRequest{}).
		Where("song_request_id = ?", id).
		Update("is_active", false).Error
}


func (r *repository) GetSongRequestByIDTable(tableNumber int) ([]*models.TrxSongRequest, error) {
	var reqs []*models.TrxSongRequest

	query := `
		SELECT tsr.*, rss.status_name, rt.table_number
		FROM live_music.trx_song_requests tsr
		LEFT JOIN live_music.ref_song_status rss 
			ON rss.status_id = tsr.status
		LEFT JOIN live_music.ref_tables rt 
			ON rt.table_id = tsr.table_id
		WHERE tsr.is_active = TRUE
		  AND rt.table_number = ?
		ORDER BY tsr.requested_at ASC
	`

	if err := r.db.Raw(query, tableNumber).Scan(&reqs).Error; err != nil {
		return nil, err
	}
	return reqs, nil
}



type SongRequestWithDetails struct {
	*models.TrxSongRequest
	TableNumber     string `json:"table_number"`
	Title           string `json:"title"`
	Artist          string `json:"artist"`
	Duration        int    `json:"duration"`
	HeaderImageURL  string `json:"header_image_url"`
	ReleaseSongDate string `json:"release_song_date"`
}

func (r *repository) ListSongRequests(limit, offset int) ([]*SongRequestWithDetails, int64, error) {
	var reqs []*SongRequestWithDetails
	var total int64

	// Paginated query for all active requests
	query := `
		SELECT tsr.*, rt.table_number, rs.title, rs.artist, rs.duration, rs.header_image_url, rs.release_song_date
		FROM trx_song_requests tsr
		LEFT JOIN ref_songs rs ON rs.song_id = tsr.song_id
		LEFT JOIN ref_tables rt ON rt.table_id = tsr.table_id
		WHERE tsr.is_active = true
		  AND rt.is_active = true
		  AND rs.is_active = true
		ORDER BY tsr.requested_at ASC
		LIMIT ? OFFSET ?
	`

	if err := r.db.Raw(query, limit, offset).Scan(&reqs).Error; err != nil {
		return nil, 0, err
	}

	// Count total rows matching the same conditions
	countQuery := `
		SELECT COUNT(*)
		FROM trx_song_requests tsr
		LEFT JOIN ref_songs rs ON rs.song_id = tsr.song_id
		LEFT JOIN ref_tables rt ON rt.table_id = tsr.table_id
		WHERE tsr.is_active = true
		  AND rt.is_active = true
		  AND rs.is_active = true
	`
	if err := r.db.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	return reqs, total, nil
}
