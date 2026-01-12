package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// clients holds all active WebSocket connections.
//
// IMPORTANT:
// This map is accessed by multiple goroutines:
//   - each HandleWS call runs in its own goroutine
//   - BroadcastAll can be called at any time
//
// Go maps are NOT thread-safe, so access must be synchronized.
var clients = make(map[*Client]bool)

// mu protects the clients map.
//
// Any read, write, delete, or iteration over `clients`
// MUST hold this mutex, otherwise the program may panic
// with "concurrent map read and map write".
var mu sync.Mutex

type InitialStateFunc func(*Client)

// Client represents a single WebSocket connection

type Client struct {
	Conn *websocket.Conn
	// TableID string // empty for public/global
}

// HandleWS upgrades HTTP connection to WebSocket
func HandleWS(w http.ResponseWriter, r *http.Request, initialState InitialStateFunc) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{Conn: conn}
	mu.Lock()
	clients[client] = true
	mu.Unlock()

	// 🎯 SEND CURRENT STATE IMMEDIATELY
	if initialState != nil {
		initialState(client)
	}
	// Keep connection alive
	// Each WebSocket connection runs in its own goroutine.
	for {
		if _, _, err := conn.NextReader(); err != nil {
			// Client disconnected.
			// Lock before removing from the shared map.
			mu.Lock()
			delete(clients, client)
			mu.Unlock()
			conn.Close()
			log.Printf("WebSocket client disconnected)")
			break
		}
	}
}

// Broadcast to all clients
func BroadcastAll(message string) {
	// Lock is required because:
	// - we are iterating over the `clients` map
	// - other goroutines may be adding/removing clients at the same time
	mu.Lock()
	defer mu.Unlock()
	for c := range clients {
		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("WebSocket write error:", err)

			// On write failure, clean up the dead connection.
			c.Conn.Close()
			delete(clients, c)
		}
	}
}
