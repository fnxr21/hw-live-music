package postgres

// import (
// 	// "fmt"
// 	"encoding/json"
// 	"log"
// 	"time"

// 	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
// 	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
// 	"github.com/lib/pq"
// )


// func StartRealtimeListener(dsn string, repo repositories.LivePlaylist) {
// 	reportProblem := func(ev pq.ListenerEventType, err error) {
// 		if err != nil {
// 			log.Println("Realtime listener error:", err)
// 		}
// 	}

// 	listener := pq.NewListener(dsn, 10*time.Second, 60*time.Second, reportProblem)
// 	if err := listener.Listen("realtime_channel"); err != nil {
// 		log.Fatal("Failed to listen on realtime_channel:", err)
// 	}

// 	log.Println("🔔 Realtime listener started on 'realtime_channel'")

// 	go func() {
// 		for {
// 			select {
// 			case n := <-listener.Notify:
// 				if n != nil {
// 					log.Println("📣 Notification received:", n.Extra)

// 					// Example: send all table-specific song requests
// 					// tableID := 2 // optionally parse from n.Extra if you encode table info
// 					data, err := repo.ListLivePlaylists()
// 					if err != nil {
// 						log.Println("Failed to fetch song requests:", err)
// 						continue
// 					}

// 					payload, _ := json.Marshal(data)
// 					ws.BroadcastMessage(string(payload))
// 				}
// 			case <-time.After(90 * time.Second):
// 				go listener.Ping()
// 			}
// 		}
// 	}()
// }
