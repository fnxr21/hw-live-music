package router

import (
	// postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	// "github.com/fnxr21/hw-live-music/backend/internal/handlers"
	// repositories "github.com/fnxr21/hw-live-music/backend/internal/repositories"
	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/labstack/echo/v4"
)

func WebSocketRoute(e *echo.Group) {
	
	// e.GET("/ws", func(c echo.Context) error {
	// 	ws.HandleWS(c.Response(), c.Request())
	// 	return nil
	// })
		// public / everyone
	e.GET("/ws/playlists", func(c echo.Context) error {
		ws.HandleWS(c.Response(), c.Request(), "") // empty tableID = public
		return nil
	})

	// private / table-specific
	e.GET("/ws/table/:id", func(c echo.Context) error {
		tableID := c.Param("id")
		ws.HandleWS(c.Response(), c.Request(), tableID)
		return nil
	})
}
