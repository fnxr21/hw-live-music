package router

import (
	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/handlers"
	repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	// ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/labstack/echo/v4"
)

func LivePlaylist(e *echo.Group) {
	repo := repositories.Repository(postgres.DB)
	h := handlers.HandlerLivePlaylist(repo)
	// e.POST("/playlist", h.CreateLivePlaylist)           
	e.GET("/playlist/:id", h.GetLivePlaylistByID)      
	e.GET("/playlists", h.ListLivePlaylists)          
	e.PUT("/playlist/:id", h.UpdateLivePlaylist)      
	e.DELETE("/playlist/:id", h.DeleteLivePlaylist)  

	// e.GET("/ws", func(c echo.Context) error {
	// 	ws.HandleWS(c.Response(), c.Request())
	// 	return nil
	// })
}
