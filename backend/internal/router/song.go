package router

import (
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/handlers"
	repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

func Song(e *echo.Group) {
	repo := repositories.Repository(postgres.DB)
	h := handlers.HandlerSong(repo)

	e.POST("/song", h.CreateSong)
	e.GET("/song/:id", h.GetSongByID)
	e.GET("/songs", h.ListSongs)
	e.PUT("/song/:id", h.UpdateSong)
	e.DELETE("/song/:id", h.DeleteSong)
}