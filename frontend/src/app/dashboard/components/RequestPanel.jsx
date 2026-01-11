"use client";

import Pagination from "./Pagination";

export default function RequestPanel({ requests, approve, reject, compact, page, total, onPageChange, loading }) {
  return (
    <>
      <h3 className="text-green-500 font-semibold mb-4">
        Song Requests
      </h3>

      {loading ? (
        <p className="text-sm text-neutral-400">Loading...</p>
      ) : requests.length === 0 ? (
        <p className="text-sm text-neutral-400">No requests</p>
      ) : (
        requests.map((r) => (
          <div
            key={r.song_request_id}
            className="bg-neutral-900 rounded px-3 py-2 mb-2 text-sm"
          >
            <p>{r.title}</p>
            {!compact && r.table_number && (
              <p className="text-xs text-neutral-400">{r.table_number}</p>
            )}

            <div className="flex gap-2 mt-2">
              <button
                onClick={() => approve(r)}
                className="bg-green-500 text-black px-2 py-1 rounded text-xs"
              >
                Approve
              </button>
              <button
                onClick={() => reject(r)}
                className="bg-red-500 px-2 py-1 rounded text-xs"
              >
                Reject
              </button>
            </div>
          </div>
        ))
      )}

      {/* Pagination */}
      {!compact && total > 1 && (
        <Pagination page={page} total={Math.ceil(total / 5)} onChange={onPageChange} />
      )}
    </>
  );
}
