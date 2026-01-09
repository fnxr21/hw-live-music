package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/labstack/echo/v4"
)

type handlerLivePlaylist struct {
	PlaylistRepo repositories.LivePlaylist
}

func HandlerLivePlaylist(repo repositories.LivePlaylist) *handlerLivePlaylist {
	return &handlerLivePlaylist{repo}
}

// CreateLivePlaylist creates a new playlist entry
func (h *handlerLivePlaylist) CreateLivePlaylist(c echo.Context) error {
	var playlist models.TrxLivePlaylist
	if err := c.Bind(&playlist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	created, err := h.PlaylistRepo.CreateLivePlaylist(playlist)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, created)
}

// ListLivePlaylists returns all active playlist entries
func (h *handlerLivePlaylist) ListLivePlaylists(c echo.Context) error {
	playlists, err := h.PlaylistRepo.ListLivePlaylists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, playlists)
}

// GetLivePlaylistByID returns a playlist entry by UUID
func (h *handlerLivePlaylist) GetLivePlaylistByID(c echo.Context) error {
	id := c.Param("id")
	playlist, err := h.PlaylistRepo.GetLivePlaylistByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if playlist == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Playlist entry not found"})
	}
	return c.JSON(http.StatusOK, playlist)
}

// UpdateLivePlaylist updates a playlist entry
func (h *handlerLivePlaylist) UpdateLivePlaylist(c echo.Context) error {
	idParam := c.Param("id")
	playlistID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist ID"})
	}

	var playlist models.TrxLivePlaylist
	if err := c.Bind(&playlist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	playlist.LivePlaylistID = playlistID
	playlist.UpdatedAt = time.Now()

	updated, err := h.PlaylistRepo.UpdateLivePlaylist(playlist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteLivePlaylist performs a soft delete
func (h *handlerLivePlaylist) DeleteLivePlaylist(c echo.Context) error {
	id := c.Param("id")
	if err := h.PlaylistRepo.DeleteLivePlaylist(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Playlist entry deleted successfully"})
}
