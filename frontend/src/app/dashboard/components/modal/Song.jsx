"use client";

import { useState } from "react";

export default function SongModal({ isOpen, onClose, onSubmit, initialData }) {
  const [title, setTitle] = useState(initialData?.title || "");
  const [artist, setArtist] = useState(initialData?.artist || "");
  const [duration, setDuration] = useState(initialData?.duration || "");

  if (!isOpen) return null;

  const handleSubmit = () => {
    onSubmit({
      ...initialData,
      title,
      artist,
      duration: duration ? parseInt(duration) : undefined,
    });
    onClose();
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-neutral-900 p-6 rounded-lg w-[400px]">
        <h2 className="text-lg font-semibold mb-4">
          {initialData ? "Update Song" : "Add New Song"}
        </h2>

        <div className="space-y-2">
          <input
            type="text"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full rounded bg-neutral-800 px-3 py-2 text-sm"
          />
          <input
            type="text"
            placeholder="Artist"
            value={artist}
            onChange={(e) => setArtist(e.target.value)}
            className="w-full rounded bg-neutral-800 px-3 py-2 text-sm"
          />
          <input
            type="number"
            placeholder="Duration (seconds)"
            value={duration}
            onChange={(e) => setDuration(e.target.value)}
            className="w-full rounded bg-neutral-800 px-3 py-2 text-sm"
          />
        </div>

        <div className="flex justify-end gap-2 mt-4">
          <button
            onClick={onClose}
            className="bg-gray-600 px-4 py-2 rounded text-sm"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            className="bg-green-500 px-4 py-2 rounded text-sm"
          >
            {initialData ? "Update" : "Add"}
          </button>
        </div>
      </div>
    </div>
  );
}
