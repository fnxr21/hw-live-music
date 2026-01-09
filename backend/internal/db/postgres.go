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


	// i dont use auto migrate for now because i prefer manual migration with golang-migrate and sql files
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

// Listen to pg_notify channel and call callback
// func ListenNotify(ctx context.Context, channel string, callback func(payload string)) {
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	conn, err := sqlDB.Conn(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	_, err = conn.ExecContext(ctx, fmt.Sprintf("LISTEN %s;", channel))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Listening to channel:", channel)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			// Wait for notifications
// 			_, err := conn.ExecContext(ctx, "SELECT 1") // simple ping to keep connection alive
// 			if err != nil {
// 				continue
// 			}
// 		}
// 	}
// }
