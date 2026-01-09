package handlers

import (
	"net/http"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerSong struct {
	SongRepo repositories.Song
}

func HandlerSong(songRepo repositories.Song) *handlerSong {
	return &handlerSong{songRepo}
}

// CreateSong handles creating a new song
func (h *handlerSong) CreateSong(c echo.Context) error {
	var song models.RefSong
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	createdSong, err := h.SongRepo.CreateSong(song)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, createdSong)
}

// ListSongs returns all songs
func (h *handlerSong) ListSongs(c echo.Context) error {
	songs, err := h.SongRepo.ListSongs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, songs)
}

// GetSongByID returns a single song by UUID
func (h *handlerSong) GetSongByID(c echo.Context) error {
	id := c.Param("id")
	song, err := h.SongRepo.GetSongByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if song == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Song not found",
		})
	}

	return c.JSON(http.StatusOK, song)
}

// UpdateSong updates an existing song
func (h *handlerSong) UpdateSong(c echo.Context) error {
	idParam := c.Param("id")

	// Parse UUID
	songID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid song ID",
		})
	}

	var song models.RefSong
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	song.SongID = songID

	updatedSong, err := h.SongRepo.UpdateSong(song)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, updatedSong)
}

// DeleteSong deletes a song by UUID
func (h *handlerSong) DeleteSong(c echo.Context) error {
	id := c.Param("id")
	if err := h.SongRepo.DeleteSong(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Song deleted successfully",
	})
}
