package postgres

import (
	"fmt"
	"log"

	// "live_music_api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Postgres connected with GORM")


	// i dont use auto migrate from gorm
	// Auto migrate tables
	// err = DB.AutoMigrate(
		// &models.Song{},
		// &models.Table{},
		// &models.SongRequest{},
		// &models.LivePlaylist{},
	// )
	// if err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }
}
