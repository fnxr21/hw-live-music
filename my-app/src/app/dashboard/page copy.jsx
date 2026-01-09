"use client";

import { useState, useMemo } from "react";

/* =======================
   PAGINATION COMPONENT
======================= */
function Pagination({ page, total, onChange }) {
  return (
    <div className="flex gap-2 mt-4">
      {Array.from({ length: total }).map((_, i) => (
        <button
          key={i}
          onClick={() => onChange(i + 1)}
          className={`px-3 py-1 rounded text-sm ${
            page === i + 1
              ? "bg-green-500 text-black"
              : "bg-neutral-800"
          }`}
        >
          {i + 1}
        </button>
      ))}
    </div>
  );
}

/* =======================
   RIGHT PANEL – LIVE VIEW
======================= */
function LivePlaylistView({ playlist }) {
  return (
    <>
      <h3 className="font-semibold text-green-500 mb-4">
        Live Playlist (User View)
      </h3>

      {playlist.map((song, i) => (
        <div
          key={song.id}
          className="flex justify-between rounded bg-neutral-800 px-3 py-2 mb-2 text-sm"
        >
          <span>
            #{song.order} {song.title}
          </span>
          {i === 0 ? (
            <span className="text-green-500">PLAYING</span>
          ) : (
            <span className="text-neutral-400">QUEUED</span>
          )}
        </div>
      ))}
    </>
  );
}

/* =======================
   RIGHT PANEL – REQUESTS
======================= */
function RequestPanel({ requests, onApprove, onReject }) {
  return (
    <>
      <h3 className="font-semibold text-green-500 mb-4">
        Song Requests
      </h3>

      {requests.length === 0 && (
        <p className="text-sm text-neutral-400">
          No pending requests
        </p>
      )}

      {requests.map((req) => (
        <div
          key={req.id}
          className="flex justify-between items-center rounded bg-neutral-800 px-3 py-2 mb-2 text-sm"
        >
          <div>
            <p>{req.title}</p>
            <p className="text-xs text-neutral-400">
              {req.table}
            </p>
          </div>

          <div className="flex gap-2">
            <button
              onClick={() => onApprove(req)}
              className="bg-green-500 text-black px-2 py-1 rounded text-xs"
            >
              Approve
            </button>
            <button
              onClick={() => onReject(req.id)}
              className="bg-red-500 text-white px-2 py-1 rounded text-xs"
            >
              Reject
            </button>
          </div>
        </div>
      ))}
    </>
  );
}

/* =======================
   MAIN DASHBOARD
======================= */
export default function Dashboard() {
  const [menu, setMenu] = useState("album");
  const [page, setPage] = useState(1);

  /* DATA */
  const albums = Array.from({ length: 12 }).map((_, i) => ({
    id: i,
    title: `Album ${i + 1}`,
  }));

  const [requests, setRequests] = useState([
    { id: 1, title: "Neon City", table: "Table 2" },
    { id: 2, title: "Lo-Fi Dreams", table: "Table 5" },
  ]);

  const [playlist, setPlaylist] = useState([
    { id: 1, title: "Midnight Drive", order: 1 },
    { id: 2, title: "Jazz After Dark", order: 2 },
  ]);

  /* PAGINATION */
  const perPage = 4;
  const totalPage = Math.ceil(albums.length / perPage);
  const pagedAlbums = useMemo(() => {
    const start = (page - 1) * perPage;
    return albums.slice(start, start + perPage);
  }, [page, albums]);

  /* ACTIONS */
  const approveRequest = (req) => {
    setPlaylist((prev) => [
      ...prev,
      {
        id: Date.now(),
        title: req.title,
        order: prev.length + 1,
      },
    ]);
    setRequests((prev) => prev.filter((r) => r.id !== req.id));
  };

  const rejectRequest = (id) => {
    setRequests((prev) => prev.filter((r) => r.id !== id));
  };

  return (
    <div className="min-h-screen bg-neutral-950 text-white flex">
      {/* LEFT MENU */}
      <aside className="w-56 border-r border-neutral-800 p-4 space-y-2">
        {["album", "playlist", "request"].map((m) => (
          <button
            key={m}
            onClick={() => setMenu(m)}
            className={`w-full text-left px-4 py-2 rounded ${
              menu === m
                ? "bg-green-500 text-black"
                : "hover:bg-neutral-800"
            }`}
          >
            {m === "album" && "Album"}
            {m === "playlist" && "Live Playlist"}
            {m === "request" && "Song Request"}
          </button>
        ))}
      </aside>

      {/* MAIN CONTENT */}
      <main className="flex-1 p-6">
        {menu === "album" && (
          <>
            <h2 className="text-lg font-semibold mb-4">
              Album
            </h2>

            <div className="grid grid-cols-2 gap-3">
              {pagedAlbums.map((a) => (
                <div
                  key={a.id}
                  className="rounded bg-neutral-900 px-4 py-3"
                >
                  {a.title}
                </div>
              ))}
            </div>

            <Pagination
              page={page}
              total={totalPage}
              onChange={setPage}
            />
          </>
        )}

        {menu === "playlist" && (
          <>
            <h2 className="text-lg font-semibold mb-4">
              Live Playlist (Admin)
            </h2>

            {playlist.map((song, i) => (
              <div
                key={song.id}
                className="flex justify-between bg-neutral-900 rounded px-4 py-3 mb-2"
              >
                <span>
                  #{song.order} {song.title}
                </span>
                <span className="text-xs text-neutral-400">
                  drag later
                </span>
              </div>
            ))}
          </>
        )}

        {menu === "request" && (
          <>
            <h2 className="text-lg font-semibold mb-4">
              Requests from Tables
            </h2>

            {requests.map((r) => (
              <div
                key={r.id}
                className="bg-neutral-900 rounded px-4 py-3 mb-2"
              >
                {r.title} –{" "}
                <span className="text-xs text-neutral-400">
                  {r.table}
                </span>
              </div>
            ))}
          </>
        )}
      </main>

      {/* RIGHT PANEL (CONDITIONAL) */}
      <aside className="w-80 border-l border-neutral-800 p-4">
        {menu === "playlist" && (
          <RequestPanel
            requests={requests}
            onApprove={approveRequest}
            onReject={rejectRequest}
          />
        )}

        {menu === "request" && (
          <LivePlaylistView playlist={playlist} />
        )}
      </aside>
    </div>
  );
}

