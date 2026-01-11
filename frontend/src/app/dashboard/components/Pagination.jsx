export default function Pagination({ page, total, onChange }) {
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
