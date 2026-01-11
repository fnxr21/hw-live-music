export default function Sidebar({ menu, setMenu }) {
  const items = [
    { key: "album", label: "Album" },
    { key: "request", label: "Song Request" },
    { key: "playlist", label: "Live Playlist" },
  ];

  return (
    <aside className="w-56 border-r border-neutral-800 p-4 space-y-2">
      {items.map((i) => (
        <button
          key={i.key}
          onClick={() => setMenu(i.key)}
          className={`w-full text-left px-4 py-2 rounded ${
            menu === i.key
              ? "bg-green-500 text-black"
              : "hover:bg-neutral-800"
          }`}
        >
          {i.label}
        </button>
      ))}
    </aside>
  );
}
