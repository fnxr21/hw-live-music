package router

import (
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/handlers"
	repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

func SongRequest(e *echo.Group) {
	// Initialize repository
	repo := repositories.Repository(postgres.DB)

	// Initialize handler
	h := handlers.HandlerSongRequest(repo)

	// Define routes
	e.POST("/song-request", h.CreateSongRequest)          // Create
	e.GET("/song-request/:id", h.GetSongRequestByID)     // Get by UUID
	e.GET("/song-requests", h.ListSongRequests)         // List all active
	e.PUT("/song-request/:id", h.UpdateSongRequest)     // Update
	e.DELETE("/song-request/:id", h.DeleteSongRequest)  // Soft delete
}
