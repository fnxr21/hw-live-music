package handlers

import (
	"net/http"
	"strconv"

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
		page := 1
	limit := 5

	if pages := c.QueryParam("page"); pages != "" {
		if parsed, err := strconv.Atoi(pages); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if limits := c.QueryParam("limit"); limits != "" {
		if parsed, err := strconv.Atoi(limits); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	offset := (page - 1) * limit

	songs,total, err := h.SongRepo.ListSongs(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// return c.JSON(http.StatusOK, songs)
		return c.JSON(http.StatusOK,  map[string]interface{}{
		"data":  songs,
		"page":  page,
		"limit": limit,
		"total": total,
	})
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
