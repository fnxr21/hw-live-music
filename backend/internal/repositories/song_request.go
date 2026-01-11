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
	ListSongRequests(limit, offset int) ([]*models.TrxSongRequest,int64, error)
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

// // ListSongRequests returns all active song requests

func (r *repository) ListSongRequests(limit, offset int) ([]*models.TrxSongRequest,int64, error) {
	var reqs []*models.TrxSongRequest
	var total int64

	// if err := r.db.Where("is_active = ?", true).Find(&reqs).Error; err != nil {
	// 	return nil,0, err
	// }
	// Fetch paginated results
	if err := r.db.
		Where("is_active = ?", true).
		Limit(limit).
		Offset(offset).
		Find(&reqs).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.TrxSongRequest{}).
		Where("is_active = ?", true).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return reqs, total,nil
}

// update this for make limit  on next patch ,focus on mvp:
// func (r *repository) ListSongRequests(limit, offset int) ([]*models.TrxSongRequest, error) {
// 	var reqs []*models.TrxSongRequest

// 	query := `
// 		SELECT tsr.*,rss.status_name ,rt.table_number 
// 		FROM live_music.trx_song_requests tsr 
// 		left join live_music.ref_song_status rss on rss.status_id  =tsr.status 
// 		left join live_music.ref_tables rt on rt.table_id  =tsr.table_id 
// 		WHERE tsr.is_active = TRUE
// 		ORDER BY tsr.requested_at ASC
// 		LIMIT ? OFFSET ?

// 	`

// 	if err := r.db.Raw(query, limit, offset).Scan(&reqs).Error; err != nil {
// 		return nil, err
// 	}

// 	return reqs, nil
// }


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
