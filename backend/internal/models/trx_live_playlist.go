package models

import (
	"time"

	"github.com/google/uuid"
)

type TrxLivePlaylist struct {
	LivePlaylistID uuid.UUID  `json:"live_playlist_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SongRequestID  uuid.UUID  `json:"song_request_id"`
	OrderNumber    *int       `json:"order_number,omitempty"`
	IsCurrent      bool       `json:"is_current" gorm:"default:false"`
	TableID        *uuid.UUID `json:"table_id,omitempty"`
	IsActive       bool       `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy      *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy      *uuid.UUID `json:"updated_by,omitempty"`
}
