package ws

// import (
// 	"log"
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/websocket"
// )

// // Upgrade HTTP connection to WebSocket
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// // Store connected clients
// var clients = make(map[*websocket.Conn]bool)
// var mu sync.Mutex

// func HandleWS(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade error:", err)
// 		return
// 	}

// 	mu.Lock()
// 	clients[conn] = true
// 	mu.Unlock()
// 	log.Println("New WebSocket client connected")

// 	// Keep connection alive
// 	for {
// 		if _, _, err := conn.NextReader(); err != nil {
// 			mu.Lock()
// 			delete(clients, conn)
// 			mu.Unlock()
// 			conn.Close()
// 			log.Println("WebSocket client disconnected")
// 			break
// 		}
// 	}
// }

// // Broadcast message to all clients
// func BroadcastMessage(message string) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	for conn := range clients {
// 		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
// 		if err != nil {
// 			log.Println("WebSocket write error:", err)
// 			conn.Close()
// 			delete(clients, conn)
// 		}
// 	}
// }
