package postgres

import (
	// "fmt"
	"encoding/json"
	"log"
	// "strconv"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/lib/pq"
)

// StartRealtimeListener listens to PostgreSQL NOTIFY events
// and broadcasts updated state to all WebSocket clients.
//
// This function:
//   1. Connects to Postgres LISTEN/NOTIFY
//   2. Waits for database events
//   3. Fetches fresh data from repositories
//   4. Broadcasts the new state over WebSockets

type NotifyPayload struct {
	Table     string `json:"table"`
	Operation string `json:"operation"`
	ID        string `json:"id"`
	TableID   string `json:"table_id,omitempty"`
}

func StartRealtimeListener(dsn string, repoPlaylist repositories.LivePlaylist, repoSongRequest repositories.SongRequest) {
	// reportProblem is a callback used by pq.Listener
	// It gets invoked when the listener encounters connection issues,
	// reconnections, or other internal errors.
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("Realtime listener error:", err)
		}
	}

	// Create a new PostgreSQL listener.
	//
	// 10s  -> minimum reconnect delay
	// 60s  -> maximum reconnect delay
	// reportProblem -> error callback
	listener := pq.NewListener(dsn, 10*time.Second, 60*time.Second, reportProblem)

	// Subscribe to the Postgres channel.
	// This must match the channel name used in `NOTIFY realtime_channel, payload`.
	if err := listener.Listen("realtime_channel"); err != nil {
		log.Fatal("Failed to listen on realtime_channel:", err)
	}

	log.Println("Realtime listener started on 'realtime_channel'")

	// Run the listener loop in its own goroutine.
	// This allows the rest of the application to continue running.
	go func() {
		for {
			select {
			// Case 1: We received a NOTIFY event from Postgres
			case n := <-listener.Notify:
				if n == nil {
					// This can happen during reconnects
					continue
				}

				log.Println("📣 Notification received:", n.Extra)

				var notify NotifyPayload
				if err := json.Unmarshal([]byte(n.Extra), &notify); err != nil {
					log.Println("Failed to parse NOTIFY payload:", err)

					continue
				}

				// Decide what to do based on which table triggered the event
				switch notify.Table {

				case "trx_live_playlists":
					// Fetch the latest playlist state from the repository.
					// We do NOT trust partial updates from the database;
					// instead, we broadcast a full fresh snapshot.
					playlists, err := repoPlaylist.RealTimeListLivePlaylists()

					if err != nil {
						log.Println("Failed to fetch playlists:", err)
						continue
					}

					payload, _ := json.Marshal(playlists)

					// Broadcast to all connected WebSocket clients
					ws.BroadcastAll(string(payload)) // everyone

				case "trx_song_requests":

					playlists, err := repoPlaylist.RealTimeListLivePlaylists()

					if err != nil {
						log.Println("Failed to fetch playlists:", err)
						continue
					}
					payload, _ := json.Marshal(playlists)

					ws.BroadcastAll(string(payload)) // everyone

				}

			// Safety heartbeat:
			// If no NOTIFY events arrive for 90s,
			// ping the database to keep the connection alive.
			case <-time.After(90 * time.Second):

				// Ping keeps the connection alive and detects stale connections.
				// This prevents the listener from silently dying.
				go listener.Ping()
			}
		}
	}()

}
