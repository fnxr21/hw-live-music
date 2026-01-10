"use client";

import { useState } from "react";
import Pagination from "./Pagination";

export default function AlbumPanel({ albums, setAlbums, addToPlaylist }) {
  const [page, setPage] = useState(1);
  const [newSong, setNewSong] = useState("");

  const perPage = 5;
  const total = Math.ceil(albums.length / perPage);
  const view = albums.slice(
    (page - 1) * perPage,
    page * perPage
  );

  const addSong = () => {
    if (!newSong.trim()) return;
    setAlbums((prev) => [
      ...prev,
      { id: Date.now().toString(), title: newSong },
    ]);
    setNewSong("");
  };

  const deleteSong = (id) => {
    setAlbums((prev) => prev.filter((s) => s.id !== id));
  };

  return (
    <>
      <h2 className="text-lg font-semibold mb-4">
        Song Library
      </h2>

      {/* ADD SONG */}
      <div className="flex gap-2 mb-4">
        <input
          value={newSong}
          onChange={(e) => setNewSong(e.target.value)}
          placeholder="New song title"
          className="flex-1 rounded bg-neutral-800 px-3 py-2 text-sm"
        />
        <button
          onClick={addSong}
          className="bg-green-500 text-black px-4 rounded text-sm"
        >
          Add
        </button>
      </div>

      {/* SONG LIST */}
      <div className="space-y-2">
        {view.map((song) => (
          <div
            key={song.id}
            className="flex justify-between items-center bg-neutral-900 rounded px-4 py-2 text-sm"
          >
            <span>{song.title}</span>

            <div className="flex gap-2">
              <button
                onClick={() => addToPlaylist(song)}
                className="bg-green-500 text-black px-2 py-1 rounded text-xs"
              >
                Add to Live
              </button>

              <button
                onClick={() => deleteSong(song.id)}
                className="text-red-400 text-xs"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>

      <Pagination page={page} total={total} onChange={setPage} />
    </>
  );
}
