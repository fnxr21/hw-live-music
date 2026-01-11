


  "use client";

  import { useEffect, useState } from "react";
  import { useParams } from "next/navigation";

  export default function TablePage() {
    const params = useParams();
    const tableId = params.id; // dynamic slug from /table/[id]

    const [playlists, setPlaylists] = useState([]);
    const [requests, setRequests] = useState([]);
    const [wsStatus, setWsStatus] = useState({
      playlists: false,
      table: false,
    });

    useEffect(() => {
      // if (!tableId) return;

      // --------------------------
      // 1️⃣ Public playlist socket
      // --------------------------
      const wsPlaylist = new WebSocket("ws://localhost:8080/api/v1/ws/playlists");

      wsPlaylist.onopen = () => {
        console.log("✅ Connected to playlist WebSocket");
        setWsStatus((s) => ({ ...s, playlists: true }));
      };

      wsPlaylist.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          setPlaylists(data); // replace state with latest playlists
        } catch (err) {
          console.error("Failed to parse playlist WS data:", err);
        }
      };

      wsPlaylist.onclose = () => {
        console.log("⚠ Playlist WebSocket disconnected");
        setWsStatus((s) => ({ ...s, playlists: false }));
      };

      wsPlaylist.onerror = (err) => {
        console.error("❌ Playlist WebSocket error:", err);
      };

      // --------------------------
      // 2️⃣ Table-specific socket
      // --------------------------
      const wsTable = new WebSocket(`ws://localhost:8080/api/v1/ws/table/1`);

      wsTable.onopen = () => {
        console.log(`✅ Connected to table 10 WebSocket`);
        setWsStatus((s) => ({ ...s, table: true }));
      };

      wsTable.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          setRequests(data); // replace state with latest table-specific requests
        } catch (err) {
          console.error("Failed to parse table WS data:", err);
        }
      };
      
      wsTable.onclose = () => {
        console.log(`⚠ Table ${tableId} WebSocket disconnected`);
        setWsStatus((s) => ({ ...s, table: false }));
      };

      wsTable.onerror = (err) => {
        console.error(`❌ Table ${tableId} WebSocket error:`, err);
      };

      // Cleanup
      return () => {
        wsPlaylist.close();
        wsTable.close();
      };
    }, [tableId]);

    return (
      <div className="p-6 bg-zinc-50 dark:bg-black min-h-screen">
        <h1 className="mb-4 text-2xl font-bold text-center text-neutral-900 dark:text-white">
          Table {tableId} - Live Updates 🎵
        </h1>

        <p className="text-center mb-4">
          Playlist WS: {wsStatus.playlists ? "🟢 Connected" : "🔴 Disconnected"} |{" "}
          Table WS: {wsStatus.table ? "🟢 Connected" : "🔴 Disconnected"}
        </p>

        <div className="mb-6">
          <h2 className="text-xl font-semibold mb-2 text-neutral-900 dark:text-white">Live Playlist</h2>
          {playlists.length === 0 ? (
            <p className="text-neutral-500 dark:text-neutral-400 text-center">Waiting for playlists...</p>
          ) : (
            playlists.map((pl, idx) => (
              <div
                key={idx}
                className="rounded-md bg-blue-100 p-3 text-sm dark:bg-blue-900 dark:text-blue-200"
              >
                <pre className="whitespace-pre-wrap">{JSON.stringify(pl, null, 2)}</pre>
              </div>
            ))
          )}
        </div>

        <div>
          <h2 className="text-xl font-semibold mb-2 text-neutral-900 dark:text-white">Table Song Requests</h2>
          {requests.length === 0 ? (
            <p className="text-neutral-500 dark:text-neutral-400 text-center">Waiting for song requests...</p>
          ) : (
            requests.map((req, idx) => (
              <div
                key={idx}
                className="rounded-md bg-green-100 p-3 text-sm dark:bg-green-900 dark:text-green-200"
              >
                <pre className="whitespace-pre-wrap">{JSON.stringify(req, null, 2)}</pre>
              </div>
            ))
          )}
        </div>
      </div>
    );
  }
