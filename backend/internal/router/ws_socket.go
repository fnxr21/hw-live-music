package router

import (
	"encoding/json"
	// "strconv"

	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func WebSocketRoute(e *echo.Group) {

	repo := repositories.Repository(postgres.DB)

	// public / everyone /
	e.GET("/ws/playlists", func(c echo.Context) error {
		ws.HandleWS(c.Response(), c.Request(),  func(client *ws.Client) {
			playlists, err := repo.RealTimeListLivePlaylists()
			if err != nil {
				return
			}
			payload, _ := json.Marshal(playlists)
			client.Conn.WriteMessage(websocket.TextMessage, payload)
		}) 
		return nil
	})
}
