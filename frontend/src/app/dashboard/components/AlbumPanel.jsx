"use client";

import { useState, useEffect } from "react";
import Pagination from "./Pagination";
import SongModal from "./modal/Song";

import { ApiListSongs, ApiUpdateSong, ApiDeleteSong, ApiCreateSong } from "@/config/api";

export default function AlbumPanel({ addToPlaylist }) {
  const [songs, setSongs] = useState([]);
  const [page, setPage] = useState(1);
  const [perPage] = useState(5);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  
const [modalOpen, setModalOpen] = useState(false);
  const [modalData, setModalData] = useState(null); // null = add, object = update

  // const [newSong, setNewSong] = useState("");

  const fetchSongs = async () => {
    setLoading(true);
    try {
      const response = await ApiListSongs({ page, limit: perPage });
      setSongs(response.data.data || []);
      setTotal(response.data.total || 0);
    } catch (err) {
      console.error("Failed to fetch songs:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSongs();
  }, [page]);

  const handleAddOrUpdate = async (form) => {
    setLoading(true);
    try {
      if (form.song_id) {
        await ApiUpdateSong(form);
      } else {
        await ApiCreateSong(form);
      }
      fetchSongs();
    } catch (err) {
      console.error("Failed to save song:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    setLoading(true);
    try {
      await ApiDeleteSong(id);
      fetchSongs();
    } catch (err) {
      console.error("Failed to delete song:", err);
    } finally {
      setLoading(false);
    }
  };


  return (
    <>
      <h2 className="text-lg font-semibold mb-4">Song Library</h2>

      {/* ADD SONG */}
      <div className="flex gap-2 mb-4">
        <button
          onClick={() => {
            setModalData(null);
            setModalOpen(true);
          }}
          className="bg-green-500 text-black px-4 rounded text-sm"
          disabled={loading}
        >
          Add Song
        </button>
        {/* <input
          value={newSong}
          // onChange={(e) => setNewSong(e.target.value)}
          placeholder="New song title"
          className="flex-1 rounded bg-neutral-800 px-3 py-2 text-sm"
        />
        <button
          onClick={()=>addSong()}
          className="bg-green-500 text-black px-4 rounded text-sm"
          disabled={loading}
        >
          Add
        </button> */}
      </div>
      <div className="space-y-2">
        {loading ? (
          <p>Loading...</p>
        ) : songs.length === 0 ? (
          <p>No songs found.</p>
        ) : (
          songs.map((song) => (
            <div
              key={song.song_id}
              className="flex justify-between items-center bg-neutral-900 rounded px-4 py-2 text-sm"
            >
              <span>{song.title}</span>
              <div className="flex gap-2">
                <button
                  onClick={() => addToPlaylist(song)}
                  className="bg-green-500 text-black px-2 py-1 rounded text-xs"
                >
                  Add to PlayList
                </button>

                <button
                  onClick={() => {
                    setModalData(song);
                    setModalOpen(true);
                  }}
                  className="bg-yellow-500 px-2 py-1 rounded text-xs"
                >
                  Edit
                </button>

                <button
                  onClick={() => handleDelete(song.song_id)}
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
        <SongModal
              isOpen={modalOpen}
              onClose={() => setModalOpen(false)}
              onSubmit={handleAddOrUpdate}
              initialData={modalData}
            />
    </>
  );
}
