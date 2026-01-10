"use client";

import { useState } from "react";
import Sidebar from "./components/Sidebar";
import AlbumPanel from "./components/AlbumPanel";
import PlaylistPanel from "./components/PlaylistPanel";
import RequestPanel from "./components/RequestPanel";
import LivePlaylistView from "./components/LivePlaylistView";

export default function Dashboard() {
  const [menu, setMenu] = useState("album");

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

 const [requests, setRequests] = useState([
  {
    id: 101,
    title: "Neon City Lights",
    table: "Table 3",
  },
  {
    id: 102,
    title: "Lo-Fi Dreams",
    table: "Table 1",
  },
  {
    id: 103,
    title: "Rainy Tokyo",
    table: "VIP 2",
  },
]);


const [playlist, setPlaylist] = useState([
  {
    id: "p1",
    songId: "1",
    title: "Midnight Drive",
    source: "admin",
  },
  {
    id: "p2",
    songId: "4",
    title: "Jazz After Dark",
    source: "admin",
  },
  {
    id: "p3",
    songId: "2",
    title: "Neon City Lights",
    source: "request",
  },
]);

  /* ACTIONS */

  const addAlbumSongToPlaylist = (song) => {
    setPlaylist((prev) => [
      ...prev,
      {
        id: Date.now().toString(),
        songId: song.id,
        title: song.title,
        source: "admin",
      },
    ]);
  };

  const approveRequest = (req) => {
    setPlaylist((prev) => [
      ...prev,
      {
        id: Date.now().toString(),
        title: req.title,
        source: "request",
      },
    ]);
    setRequests((prev) => prev.filter((r) => r.id !== req.id));
  };

  const rejectRequest = (id) => {
    setRequests((prev) => prev.filter((r) => r.id !== id));
  };

  return (
    <>
      <Sidebar menu={menu} setMenu={setMenu} />

      <main className="flex-1 p-6">
        {menu === "album" && (
          <AlbumPanel
            albums={albums}
            setAlbums={setAlbums}
            addToPlaylist={addAlbumSongToPlaylist}
          />
        )}

        {menu === "playlist" && (
          <PlaylistPanel
            playlist={playlist}
            setPlaylist={setPlaylist}
          />
        )}

        {menu === "request" && (
          <RequestPanel
            requests={requests}
            approve={approveRequest}
            reject={rejectRequest}
          />
        )}
      </main>

      {/* RIGHT PANEL */}
      <aside className="w-80 border-l border-neutral-800 p-4">
        {menu === "playlist" && (
          <RequestPanel
            requests={requests}
            approve={approveRequest}
            reject={rejectRequest}
            compact
          />
        )}

        {menu === "request" && (
          <LivePlaylistView playlist={playlist} />
        )}
      </aside>
    </>
  );
}
