package models

import (
	"time"

	"github.com/google/uuid"
)

type RefUser struct {
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Token     *string    `json:"token,omitempty"`
	Role      string     `json:"role"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
}
