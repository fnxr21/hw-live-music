package models

import (
	"time"

	"github.com/google/uuid"
)

type RefSong struct {
	SongID          uuid.UUID  `json:"song_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title           string     `json:"title"`
	Artist          string     `json:"artist"`
	Duration        *int       `json:"duration,omitempty"`          
	HeaderImageURL  *string    `json:"header_image_url,omitempty"`  
	URL             *string    `json:"url,omitempty"`               
	ReleaseSongDate *time.Time `json:"release_song_date,omitempty"` 
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy       *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy       *uuid.UUID `json:"updated_by,omitempty"`
}
