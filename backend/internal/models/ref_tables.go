package models

import (
	"time"

	"github.com/google/uuid"
)

type RefTable struct {
	TableID   uuid.UUID  `json:"table_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TableNumber int      `json:"table_number"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
}
