package main

import (
	"log"

	"github.com/fnxr21/hw-live-music/backend/internal/config"
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/fnxr21/hw-live-music/backend/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	postgres.Connect(cfg.DbURL)
	// Start the realtime listener
		// Create repository instance
	repo := repositories.Repository(postgres.DB) // your repo struct that implements SongRequest

	postgres.StartRealtimeListener(cfg.DbURL, repo, repo)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))
	router.RouterInit(e.Group("/api/v1"))

	log.Printf("Server running at :%s\n", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
