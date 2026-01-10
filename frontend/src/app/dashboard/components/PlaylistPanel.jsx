import {
  DndContext,
  closestCenter,
} from "@dnd-kit/core";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
  arrayMove,
} from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";

function SortableItem({ song }) {
  const { attributes, listeners, setNodeRef, transform, transition } =
    useSortable({ id: song.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className="bg-neutral-900 rounded px-4 py-3 mb-2 cursor-grab"
    >
      {song.title}
    </div>
  );
}

export default function PlaylistPanel({ playlist, setPlaylist }) {
  const handleDragEnd = (e) => {
    const { active, over } = e;
    if (!over || active.id === over.id) return;

    setPlaylist((items) => {
      const oldIndex = items.findIndex(
        (i) => i.id === active.id
      );
      const newIndex = items.findIndex(
        (i) => i.id === over.id
      );
      return arrayMove(items, oldIndex, newIndex);
    });
  };

  return (
    <>
      <h2 className="text-lg font-semibold mb-4">
        Live Playlist (Admin)
      </h2>

      <DndContext
        collisionDetection={closestCenter}
        onDragEnd={handleDragEnd}
      >
        <SortableContext
          items={playlist.map((p) => p.id)}
          strategy={verticalListSortingStrategy}
        >
          {playlist.map((song) => (
            <SortableItem key={song.id} song={song} />
          ))}
        </SortableContext>
      </DndContext>
    </>
  );
}
