package handlers

import (
	"net/http"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/models"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerSongRequest struct {
	Repo repositories.SongRequest
}

func HandlerSongRequest(repo repositories.SongRequest) *handlerSongRequest {
	return &handlerSongRequest{repo}
}

// CreateSongRequest creates a new song request
func (h *handlerSongRequest) CreateSongRequest(c echo.Context) error {
	var req models.TrxSongRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	created, err := h.Repo.CreateSongRequest(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, created)
}

// ListSongRequests returns all active song requests
func (h *handlerSongRequest) ListSongRequests(c echo.Context) error {
	reqs, err := h.Repo.ListSongRequests()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, reqs)
}

// GetSongRequestByID fetches a single song request by UUID
func (h *handlerSongRequest) GetSongRequestByID(c echo.Context) error {
	id := c.Param("id")
	req, err := h.Repo.GetSongRequestByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if req == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Song request not found"})
	}
	return c.JSON(http.StatusOK, req)
}

// UpdateSongRequest updates an existing song request
func (h *handlerSongRequest) UpdateSongRequest(c echo.Context) error {
	idParam := c.Param("id")
	requestID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid song request ID"})
	}

	var req models.TrxSongRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	req.SongRequestID = requestID
	req.UpdatedAt = time.Now()

	updated, err := h.Repo.UpdateSongRequest(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// func (h *handlerSongRequest) UpdateSongRequest(c echo.Context) error {
// 	// 1. Parse UUID from path
// 	idParam := c.Param("song_request_id")
// 	songRequestID, err := uuid.Parse(idParam)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid song request ID"})
// 	}

// 	// 2. Bind JSON payload to model
// 	var req models.TrxSongRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
// 	}

// 	// 3. Set the ID from path
// 	req.SongRequestID = songRequestID

// 	// 4. Call repository update
// 	updated, err := h.Repo.UpdateSongRequest(req)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, updated)
// }

// DeleteSongRequest performs a soft delete
func (h *handlerSongRequest) DeleteSongRequest(c echo.Context) error {
	id := c.Param("id")
	if err := h.Repo.DeleteSongRequest(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Song request deleted successfully"})
}





// // ListSongRequests returns all active song requests
// func (h *handlerSongRequest) ListSongRequests(c echo.Context) error {
// 	page := 1
// 	limit := 5
// 	tableID := 0

// 	if pages := c.QueryParam("page"); pages != "" {
// 		if parsed, err := strconv.Atoi(pages); err == nil && parsed > 0 {
// 			page = parsed
// 		}
// 	}

// 	if limits := c.QueryParam("limit"); limits != "" {
// 		if parsed, err := strconv.Atoi(limits); err == nil && parsed > 0 {
// 			limit = parsed
// 		}
// 	}

// 	if tableid := c.QueryParam("tableId"); tableid != "" {
// 		if parsed, err := strconv.Atoi(tableid); err == nil && parsed > 0 {
// 			tableID = parsed
// 		}
// 	}

// 	offset := (page - 1) * limit

// 	reqs,total, err := h.Repo.ListSongRequests(limit, offset,tableID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}
	
// 	return c.JSON(http.StatusOK,  map[string]interface{}{
// 		"data":  reqs,
// 		"page":  page,
// 		"limit": limit,
// 		"total": total,
// 	})
// }