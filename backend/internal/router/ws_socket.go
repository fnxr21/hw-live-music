package router

import (
	"encoding/json"
	"strconv"

	postgres "github.com/fnxr21/hw-live-music/backend/internal/db"
	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func WebSocketRoute(e *echo.Group) {

	repo := repositories.Repository(postgres.DB)

	// public / everyone
	e.GET("/ws/playlists", func(c echo.Context) error {
		ws.HandleWS(c.Response(), c.Request(), "", func(client *ws.Client) {
			playlists, err := repo.ListLivePlaylists()
			if err != nil {
				return
			}
			payload, _ := json.Marshal(playlists)
			client.Conn.WriteMessage(websocket.TextMessage, payload)
		}) // empty tableID = public
		return nil
	})

	// private / table-specific
	e.GET("/ws/table/:id", func(c echo.Context) error {
		tableID := c.Param("id")
		ws.HandleWS(c.Response(), c.Request(), tableID, func(client *ws.Client) {
			id, err := strconv.Atoi(tableID)
			if err != nil {
				return
			}

			reqs, err := repo.GetSongRequestByIDTable(id)
			if err != nil {
				return
			}

			payload, _ := json.Marshal(reqs)
			client.Conn.WriteMessage(websocket.TextMessage, payload)
		})
		return nil
	})

	// example test

	// private / table-specific
	e.GET("/ws/client/:id", func(c echo.Context) error {
		tableID := c.Param("id")
		ws.HandleWS(c.Response(), c.Request(), tableID, func(client *ws.Client) {

			playlists, err := repo.ListLivePlaylists()
			if err != nil {
				return
			}
			id, err := strconv.Atoi(tableID)
			if err != nil {
				return
			}

			reqs, err := repo.GetSongRequestByIDTable(id)
			if err != nil {
				return
			}

			payload := map[string]interface{}{
				"playlists": playlists,
				"requests":  reqs,
			}
			data, err := json.Marshal(payload)
			if err != nil {
				return
			}

			// payload, _ := json.Marshal(reqs)
			client.Conn.WriteMessage(websocket.TextMessage, data)
		})
		return nil
	})
}
