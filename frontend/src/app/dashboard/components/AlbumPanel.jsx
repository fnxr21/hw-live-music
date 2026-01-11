"use client";

import { useState, useEffect } from "react";
import Pagination from "./Pagination";
import { ApiListSongs, ApiCreateSong, ApiDeleteSong } from "@/config/api";

export default function AlbumPanel({ addToPlaylist }) {
  const [albums, setAlbums] = useState([]);
  const [page, setPage] = useState(1);
  const [perPage] = useState(5);
  const [total, setTotal] = useState(0);
  const [newSong, setNewSong] = useState("");
  const [loading, setLoading] = useState(false);

  // Fetch songs from backend
  const fetchSongs = async () => {
    setLoading(true);
    try {
      const response = await ApiListSongs({ page, limit: perPage,tableId:1 });
      setAlbums(response.data.data || []);
      setTotal(response.data.total || 0);
    } catch (error) {
      console.error("Failed to fetch songs:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSongs();
  }, [page]);

  // Add new song via backend
  // const addSong = async () => {
  //   if (!newSong.trim()) return;
  //   try {
  //     await ApiCreateSong({ title: newSong });
  //     setNewSong("");
  //     fetchSongs(); // refresh list
  //   } catch (error) {
  //     console.error("Failed to create song:", error);
  //   }
  // };

  // Delete song via backend
  // const deleteSong = async (id) => {
  //   try {
  //     await ApiDeleteSong(id);
  //     fetchSongs(); // refresh list
  //   } catch (error) {
  //     console.error("Failed to delete song:", error);
  //   }
  // };

  return (
    <>
      <h2 className="text-lg font-semibold mb-4">Song Library</h2>

      {/* ADD SONG */}
      <div className="flex gap-2 mb-4">
        <input
          value={newSong}
          onChange={(e) => setNewSong(e.target.value)}
          placeholder="New song title"
          className="flex-1 rounded bg-neutral-800 px-3 py-2 text-sm"
        />
        <button
          // onClick={addSong}
          className="bg-green-500 text-black px-4 rounded text-sm"
          disabled={loading}
        >
          Add
        </button>
      </div>

      {/* SONG LIST */}
      <div className="space-y-2">
        {loading ? (
          <p>Loading...</p>
        ) : albums.length === 0 ? (
          <p>No songs found.</p>
        ) : (
          albums.map((song,index) => (
            <div
              key={index}
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
                  // onClick={() => deleteSong(song.id)}
                  className="text-red-400 text-xs"
                  disabled={loading}
                >
                  Delete
                </button>
              </div>
            </div>
          ))
        )}
      </div>

      <Pagination
        page={page}
        total={Math.ceil(total / perPage)}
        onChange={setPage}
      />
    </>
  );
}
