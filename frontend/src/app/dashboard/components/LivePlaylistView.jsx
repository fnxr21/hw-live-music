export default function LivePlaylistView({ playlist }) {
  return (
    <>
      <h3 className="text-green-500 font-semibold mb-4">
        Live Playlist
      </h3>

      {playlist.map((s, i) => (
        <div
          key={s.id}
          className="bg-neutral-900 rounded px-3 py-2 mb-2 text-sm flex justify-between"
        >
          <span>
            #{i + 1} {s.title}
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
