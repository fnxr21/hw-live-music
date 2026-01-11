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

// store connected clients
// store *Client, not *websocket.Conn
var clients = make(map[*Client]bool) 


var mu sync.Mutex




type InitialStateFunc func(*Client)

// Clients structure
type Client struct {
	Conn    *websocket.Conn
	TableID string // empty for public/global
}


// HandleWS upgrades HTTP connection to WebSocket
func HandleWS(w http.ResponseWriter, r *http.Request, tableID string, initialState InitialStateFunc) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{Conn: conn, TableID: tableID}
	mu.Lock()
	clients[client] = true
	mu.Unlock()
	log.Printf("New WebSocket client connected (tableID=%s)", tableID)

	// 🎯 SEND CURRENT STATE IMMEDIATELY
	if initialState != nil {
		initialState(client)
	}
	// Keep connection alive
	for {
		if _, _, err := conn.NextReader(); err != nil {
			mu.Lock()
			delete(clients, client)
			mu.Unlock()
			conn.Close()
			log.Printf("WebSocket client disconnected (tableID=%s)", tableID)
			break
		}
	}
}






// Broadcast to all clients
func BroadcastAll(message string) {
	mu.Lock()
	defer mu.Unlock()
	for c := range clients {
		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("WebSocket write error:", err)
			c.Conn.Close()
			delete(clients, c)
		}
	}
}

// Broadcast to specific tableID only
func BroadcastTable(tableID string, message string) {
	mu.Lock()
	defer mu.Unlock()
	for c := range clients {
		if c.TableID == tableID {
			if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("WebSocket write error:", err)
				c.Conn.Close()
				delete(clients, c)
			}
		}
	}
}


