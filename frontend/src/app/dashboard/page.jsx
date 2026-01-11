"use client";

import { useState, useEffect } from "react";
import Sidebar from "./components/Sidebar";
import AlbumPanel from "./components/AlbumPanel";
import PlaylistPanel from "./components/PlaylistPanel";
import RequestPanel from "./components/RequestPanel";
import LivePlaylistView from "./components/LivePlaylistView";

import { ApiListSongsRequest, ApiUpdateSongRequest } from "@/config/api";
import RightRequestPanel from "./components/RightRequestPanel";

export default function Dashboard() {
  const [menu, setMenu] = useState("album");

  const [page, setPage] = useState(1);
  const [perPage] = useState(5);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  const [albums, setAlbums] = useState([
    { id: "1", title: "Midnight Drive" },
    { id: "2", title: "Neon City Lights" },
    { id: "3", title: "Lo-Fi Dreams" },
    { id: "4", title: "Jazz After Dark" },
    { id: "5", title: "Rainy Tokyo" },
    { id: "6", title: "Analog Love" },
    { id: "7", title: "City Pop Sunset" },
    { id: "8", title: "Late Night Coding" },
    { id: "9", title: "Smooth Espresso" },
    { id: "10", title: "Vinyl Memories" },
  ]);

  const [requests, setRequests] = useState([]);
  const [playlist, setPlaylist] = useState([
    { id: "p1", songId: "1", title: "Midnight Drive", source: "admin" },
    { id: "p2", songId: "4", title: "Jazz After Dark", source: "admin" },
    { id: "p3", songId: "2", title: "Neon City Lights", source: "request" },
  ]);

  /* Fetch paginated song requests from backend */
  const fetchSongRequest = async () => {
    setLoading(true);
    try {
      const response = await ApiListSongsRequest({ page, limit: perPage });
      setRequests(response.data.data || []);
      setTotal(response.data.total || 0);
    } catch (err) {
      console.error("Failed to fetch requests:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSongRequest();
  }, [page]);

  /* ACTIONS */

  const addAlbumSongToPlaylist = (song) => {
    setPlaylist((prev) => [
      ...prev,
      { id: Date.now().toString(), songId: song.id, title: song.title, source: "admin" },
    ]);
  };

  /* Approve or Reject a request */
  const updateStatus = async (req, statusUuid) => {
    try {
      await ApiUpdateSongRequest({ id: req.song_request_id, status: statusUuid });

      // If approved, add to live playlist
      if (statusUuid === "d5b943d1-4bfa-436c-bb70-cb96cdcc3491") {
        setPlaylist((prev) => [
          ...prev,
          { id: Date.now().toString(), songId: req.song_id, title: req.title, source: "request" },
        ]);
      }

      // Refresh requests list
      fetchSongRequest();
    } catch (err) {
      console.error("Failed to update request status:", err);
    }
  };

  const approveRequest = (req) =>
    updateStatus(req, "d5b943d1-4bfa-436c-bb70-cb96cdcc3491");
  const rejectRequest = (req) =>
    updateStatus(req, "46702771-272a-4bc3-814f-d1b9373ae4fa");

  return (
    <>
      <Sidebar menu={menu} setMenu={setMenu} />

      <main className="flex-1 p-6">
        {menu === "album" && (
          <AlbumPanel albums={albums} setAlbums={setAlbums} addToPlaylist={addAlbumSongToPlaylist} />
        )}

        {menu === "playlist" && <PlaylistPanel playlist={playlist} setPlaylist={setPlaylist} />}

        {menu === "request" && (
          <RequestPanel
  requests={requests}
  approve={approveRequest}
  reject={rejectRequest}
  page={page}
  total={total}
  onPageChange={setPage}
  loading={loading}
/>
          // <RequestPanel requests={requests} approve={approveRequest} reject={rejectRequest} />
        )}
      </main>

      <aside className="w-80 border-l border-neutral-800 p-4">
        {menu === "playlist" && (

          <RightRequestPanel requests={requests} approve={approveRequest} reject={rejectRequest} compact />
        )}

        {menu === "request" && <LivePlaylistView playlist={playlist} />}
      </aside>
    </>
  );
}
