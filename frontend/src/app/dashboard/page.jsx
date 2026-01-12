"use client";

import { useState } from "react";
import Sidebar from "./components/Sidebar";
import AlbumPanel from "./components/AlbumPanel";
import PlaylistPanel from "./components/PlaylistPanel";
import RequestPanel from "./components/RequestPanel";
import RightRequestPanel from "./components/RightRequestPanel";
import { useSongRequests } from "@/hooks/useSongRequest";
import { useLivePlaylist } from "@/hooks/usePlaylist";
import LivePlaylistView from "./components/LivePlaylistView";

export default function Dashboard() {
  const [menu, setMenu] = useState("playlist");

  const { requests, page, setPage, total, loading, updateStatus } = useSongRequests();

  const approveRequest = (req) =>
    updateStatus(req, "d5b943d1-4bfa-436c-bb70-cb96cdcc3491");
  const rejectRequest = (req) =>
    updateStatus(req, "46702771-272a-4bc3-814f-d1b9373ae4fa");

  const {
    playlist,
    setPlaylist,
    updatePlaylist,
  } = useLivePlaylist();

  const nextSong = () => {
    console.log("Next song triggered!");
    // implement your next song logic
  };

  const deleteSong = (id) => {
    setPlaylist((prev) => prev.filter((song) => song.id !== id));
    // optionally call API to delete
  };

  // Handle drag-and-drop reorder
  const handleReorder = async (updatedPlaylist) => {
    setPlaylist(updatedPlaylist);

    // Hit API for each song's new order_number
    for (const song of updatedPlaylist) {
      await updatePlaylist(song);
    }
  };

  return (
    <>
      <Sidebar menu={menu} setMenu={setMenu} />
      <main className="flex-1 p-6">
        {menu === "album" && <AlbumPanel />}

        {menu === "playlist" && (
          <PlaylistPanel
            playlist={playlist}
            setPlaylist={setPlaylist}
            onReorder={handleReorder}
            next={nextSong}
            deleteSong={deleteSong}
          />
        )}

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
        )}
      </main>

      <aside className="w-80 border-l border-neutral-800 p-4">
        {menu === "playlist" && (
          <RightRequestPanel
            requests={requests}
            approve={approveRequest}
            reject={rejectRequest}
            compact
          />
        )}
        {menu === "album" &&    <LivePlaylistView
            playlist={playlist}
           
          />}
        {menu === "request" &&    <LivePlaylistView
            playlist={playlist}
        
          />}
      </aside>
    </>
  );
}
