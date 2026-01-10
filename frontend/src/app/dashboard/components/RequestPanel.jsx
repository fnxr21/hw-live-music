export default function RequestPanel({
  requests,
  approve,
  reject,
  compact,
}) {
  return (
    <>
      <h3 className="text-green-500 font-semibold mb-4">
        Song Requests
      </h3>

      {requests.length === 0 && (
        <p className="text-sm text-neutral-400">
          No requests
        </p>
      )}

      {requests.map((r) => (
        <div
          key={r.id}
          className="bg-neutral-900 rounded px-3 py-2 mb-2 text-sm"
        >
          <p>{r.title}</p>
          {!compact && (
            <p className="text-xs text-neutral-400">
              {r.table}
            </p>
          )}

          <div className="flex gap-2 mt-2">
            <button
              onClick={() => approve(r)}
              className="bg-green-500 text-black px-2 py-1 rounded text-xs"
            >
              Approve
            </button>
            <button
              onClick={() => reject(r.id)}
              className="bg-red-500 px-2 py-1 rounded text-xs"
            >
              Reject
            </button>
          </div>
        </div>
      ))}
    </>
  );
}
