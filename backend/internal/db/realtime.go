package postgres

import (
	// "fmt"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/fnxr21/hw-live-music/backend/internal/repositories"
	ws "github.com/fnxr21/hw-live-music/backend/internal/websocket"
	"github.com/lib/pq"
)

type NotifyPayload struct {
	Table     string `json:"table"`
	Operation string `json:"operation"`
	ID        string `json:"id"`
	TableID   string `json:"table_id,omitempty"`
}

func StartRealtimeListener(dsn string, repoPlaylist repositories.LivePlaylist, repoSongRequest repositories.SongRequest) {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("Realtime listener error:", err)
		}
	}

	listener := pq.NewListener(dsn, 10*time.Second, 60*time.Second, reportProblem)
	if err := listener.Listen("realtime_channel"); err != nil {
		log.Fatal("Failed to listen on realtime_channel:", err)
	}

	log.Println("Realtime listener started on 'realtime_channel'")

	go func() {
		for {
			select {
			case n := <-listener.Notify:
				if n == nil {
					continue
				}

				log.Println("📣 Notification received:", n.Extra)

				var notify NotifyPayload
				if err := json.Unmarshal([]byte(n.Extra), &notify); err != nil {
					log.Println("Failed to parse NOTIFY payload:", err)
					continue
				}

				switch notify.Table {
				case "trx_live_playlists":
					playlists, err := repoPlaylist.ListLivePlaylists()

					if err != nil {
						log.Println("Failed to fetch playlists:", err)
						continue
					}
					payload, _ := json.Marshal(playlists)
					ws.BroadcastAll(string(payload)) // everyone

					// broadcast to clients
				// ws.BroadcastClientPlaylists(string(payload))
				// broadcast to admins
				// ws.BroadcastAdminPlaylists(string(payload))

				case "trx_song_requests":
					// All requests for admin
					// allRequests, err := repoSongRequest.ListSongRequests()
					// if err != nil {
					// 	log.Println("Failed to fetch all song requests:", err)
					// 	continue
					// }
					// adminPayload, _ := json.Marshal(allRequests)
					// ws.BroadcastAdminSongRequests(string(adminPayload))

					// if notify.TableID == "" {
					//     log.Println("No tableID in notify, skipping")
					//     continue
					// }
					tableID, err := strconv.Atoi(notify.TableID)
					requests, err := repoSongRequest.GetSongRequestByIDTable(tableID)
					// requests, err := repoPlaylist.ListLivePlaylists()

					if err != nil {
						log.Println("Failed to fetch song requests:", err)
						continue
					}
					payload, _ := json.Marshal(requests)
					ws.BroadcastTable(notify.TableID, string(payload)) // only table
					//  ws.BroadcastClientSongRequests(notify.TableID, string(payload))
				}

			case <-time.After(90 * time.Second):
				go listener.Ping()
			}
		}
	}()

}
