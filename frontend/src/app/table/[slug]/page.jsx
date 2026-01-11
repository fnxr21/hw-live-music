


"use client";

import { API } from "@/config/axios";
import { useState, useEffect, useMemo } from "react";
// import { API } from "../path/to/API"; // import your Axios instance

export default function Page() {
  // UI state
  const [showSearch, setShowSearch] = useState(false);
  const [showLive, setShowLive] = useState(true);
  const [searchQuery, setSearchQuery] = useState("");

  // Data state
  const [albums, setAlbums] = useState([]);
  const [songRequests, setSongRequests] = useState([]);
  const [livePlaylist, setLivePlaylist] = useState([]);

  // =========================
  // Fetch albums from API (optional)
  // =========================
  // useEffect(() => {
  //   API.get("/spn") // adjust endpoint if needed
  //     .then((res) => setAlbums(res.data))
  //     .catch((err) => console.error("Error fetching albums:", err));
  // }, []);

  // =========================
  // Fetch live playlist
  // =========================
  useEffect(() => {
    API.get("/playlists")
      .then((res) => setLivePlaylist(res.data))
      .catch((err) => console.error("Error fetching live playlist:", err));
  }, []);

  // =========================
  // Fetch user song requests
  // =========================
  useEffect(() => {
    API.get("/song-requests")
      .then((res) => setSongRequests(res.data))
      .catch((err) => console.error("Error fetching song requests:", err));
  }, []);

  // =========================
  // Sort live playlist by order
  // =========================
  const sortedPlaylist = useMemo(() => {
    return [...livePlaylist].sort((a, b) => a.order - b.order);
  }, [livePlaylist]);

  const nowPlaying = sortedPlaylist.find((s) => s.status === "playing");
  const upcoming = sortedPlaylist.filter((s) => s.status === "queued");

  // =========================
  // Filter albums by search
  // =========================
  const filteredAlbums = albums.filter((album) =>
    album.title.toLowerCase().includes(searchQuery.toLowerCase())
  );

  // =========================
  // Request a song
  // =========================
  const requestSong = async (title) => {
    try {
      const res = await API.post("/song-request", { title });
      setSongRequests((prev) => [...prev, res.data]);
    } catch (err) {
      console.error("Error requesting song:", err);
    }
  };

  // =========================
  // Status color helper
  // =========================
  const statusColor = (status) => {
    switch (status) {
      case "pending":
        return "bg-yellow-500 text-black";
      case "rejected":
        return "bg-red-500 text-white";
      case "approved":
        return "bg-green-500 text-black";
      default:
        return "bg-neutral-700";
    }
  };

  return (
    <div className="min-h-screen bg-neutral-950 text-white">
      {/* HEADER */}
      <header className="sticky top-0 z-30 border-b border-neutral-800 bg-neutral-900/80 backdrop-blur">
        <div className="flex items-center justify-between px-6 py-4">
          <h1 className="text-xl font-bold text-green-500">Live Music</h1>

          <div className="flex items-center gap-3">
            <button
              onClick={() => setShowLive(!showLive)}
              className="rounded-full border border-neutral-700 px-4 py-2 text-sm"
            >
              {showLive ? "Hide Live Playlist" : "Open Live Playlist"}
            </button>

            <input
              onFocus={() => setShowSearch(true)}
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="What do you want to listen to?"
              className="w-64 rounded-full bg-neutral-800 px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>
        </div>
      </header>

      {/* MAIN */}
      <main className="px-6 py-8 space-y-12">
        {/* USER REQUEST STATUS */}
        <section>
          <h2 className="mb-4 text-lg font-semibold">My Song Requests</h2>

          <div className="space-y-3">
            {songRequests.map((req) => (
              <div
                key={req.id}
                className="flex justify-between rounded bg-neutral-900 px-4 py-3"
              >
                <span>{req.title}</span>
                <span
                  className={`rounded-full px-3 py-1 text-xs font-semibold ${statusColor(
                    req.status
                  )}`}
                >
                  {req.status.toUpperCase()}
                </span>
              </div>
            ))}
          </div>
        </section>
      </main>

      {/* SEARCH / REQUEST MODAL */}
      {showSearch && (
        <div className="fixed inset-0 z-50 bg-black/60 backdrop-blur-sm">
          <div className="mx-auto mt-24 w-full max-w-xl rounded bg-neutral-900 p-6">
            <div className="mb-4 flex justify-between">
              <h3 className="text-lg font-semibold">Request Song</h3>
              <button
                onClick={() => {
                  setShowSearch(false);
                  setSearchQuery("");
                }}
              >
                ✕
              </button>
            </div>

            <div className="space-y-3">
              {filteredAlbums.map((album) => (
                <div
                  key={album.id}
                  className="flex justify-between rounded bg-neutral-800 px-4 py-3"
                >
                  <span>{album.title}</span>
                  <button
                    onClick={() => requestSong(album.title)}
                    className="rounded-full bg-green-500 px-4 py-1 text-sm text-black"
                  >
                    Add
                  </button>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* LIVE PLAYLIST (READ ONLY) */}
      {showLive && (
        <div className="fixed right-0 top-20 z-20 h-[calc(100vh-5rem)] w-80 border-l border-neutral-800 bg-neutral-900">
          <div className="border-b border-neutral-800 p-4 flex justify-between">
            <h3 className="font-semibold text-green-500">
              Live Playlist 🎧
            </h3>
            <button onClick={() => setShowLive(false)}>✕</button>
          </div>

          {/* NOW PLAYING */}
          {nowPlaying && (
            <div className="border-b border-neutral-800 p-4">
              <p className="text-xs text-neutral-400 mb-2">Now Playing</p>
              <div className="flex gap-4 items-center">
                <div className="h-14 w-14 rounded-full bg-neutral-700 animate-spin-slow relative">
                  <div className="absolute inset-4 rounded-full bg-neutral-900" />
                </div>
                <div>
                  <p className="font-medium">{nowPlaying.title}</p>
                  <p className="text-xs text-green-500">Playing</p>
                </div>
              </div>
            </div>
          )}

          {/* UPCOMING */}
          <div className="p-4 space-y-3 text-sm">
            {upcoming.map((song, index) => (
              <div
                key={song.id}
                className="flex justify-between rounded bg-neutral-800 px-3 py-2"
              >
                <span>
                  {index + 2}. {song.title}
                </span>
                <span className="text-xs text-neutral-400">QUEUED</span>
              </div>
            ))}
          </div>

          <div className="border-t border-neutral-800 p-3 text-center text-xs text-neutral-500">
            Playlist is managed by the host
          </div>
        </div>
      )}

      {/* ROTATION */}
      <style jsx global>{`
        @keyframes spinSlow {
          from {
            transform: rotate(0deg);
          }
          to {
            transform: rotate(360deg);
          }
        }
        .animate-spin-slow {
          animation: spinSlow 4s linear infinite;
        }
      `}</style>
    </div>
  );
}
