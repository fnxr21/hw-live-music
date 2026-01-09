package repositories

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func Repository(db *gorm.DB) *repository {
	return &repository{db}
}