package models

import (
	"time"

	"github.com/google/uuid"
)

type TrxSongRequest struct {
	SongRequestID uuid.UUID  `json:"song_request_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TableID       *uuid.UUID `json:"table_id,omitempty"`
	SongID        uuid.UUID  `json:"song_id"`
	Status        uuid.UUID  `json:"status"` // Could map to status table
	RequestedAt   time.Time  `json:"requested_at" gorm:"autoCreateTime"`
	ApprovedAt    *time.Time `json:"approved_at,omitempty"`
	ApprovedBy    *uuid.UUID `json:"approved_by,omitempty"`
	IsActive      bool       `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy     *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy     *uuid.UUID `json:"updated_by,omitempty"`
}
