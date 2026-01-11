// "use client";


// import {
//   DndContext,
//   closestCenter,
// } from "@dnd-kit/core";
// import {
//   SortableContext,
//   useSortable,
//   verticalListSortingStrategy,
//   arrayMove,
// } from "@dnd-kit/sortable";
// import { CSS } from "@dnd-kit/utilities";
// import { List, Trash, SkipForward } from "phosphor-react";

// function SortableItem({ index, song, next, deleteSong }) {
//   const { attributes, listeners, setNodeRef, transform, transition } =
//     useSortable({ id: song.id });

//   const style = {
//     transform: CSS.Transform.toString(transform),
//     transition,
//   };

//   return (
//     <div
//       ref={setNodeRef}
//       style={style}
//       {...attributes}
//       {...listeners}
//       className="bg-neutral-900 flex justify-between items-center rounded px-4 py-3 mb-2 cursor-grab"
//     >
//       <span>
//         {index + 1}. {song.title}
//       </span>

//       <span className="flex gap-4">
//         {index === 0 ? (
//           <SkipForward
//             onClick={() => next()}
//             className="text-white bg-green-600 hover:bg-green-500 p-2 rounded-full cursor-pointer transition-colors duration-200"
//             size={32}
//           />
//         ) : (
//           <Trash
//             onClick={() => deleteSong(song.id)}
//             className="text-white bg-red-600 hover:bg-red-500 p-2 rounded-full cursor-pointer transition-colors duration-200"
//             size={32}
//           />
//         )}

//         <List
//           size={32}
//           className="text-neutral-400 hover:text-white cursor-pointer transition-colors duration-200"
//         />
//       </span>

//     </div>
//   );
// }

// export default function PlaylistPanel({ playlist, setPlaylist }) {
//   const handleDragEnd = (e) => {
//     const { active, over } = e;
//     if (!over || active.id === over.id) return;

//     setPlaylist((items) => {
//       const oldIndex = items.findIndex((i) => i.id === active.id);
//       const newIndex = items.findIndex((i) => i.id === over.id);
//       return arrayMove(items, oldIndex, newIndex);
//     });
//   };

//   const next = () => {
//     console.log("Next song triggered!");
//     // implement logic to play next song
//   };

//   const deleteSong = (id) => {

//     // setPlaylist((prev) => prev.filter((song) => song.id !== id));
//   };

//   return (
//     <>
//       <h2 className="text-lg font-semibold mb-4">Live Playlist </h2>

//       <DndContext collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
//         <SortableContext
//           items={playlist.map((p) => p.id)}
//           strategy={verticalListSortingStrategy}
//         >
//           {playlist.map((song, index) => (
//             <SortableItem
//               key={song.id}
//               index={index}
//               song={song}
//               next={next}
//               deleteSong={deleteSong}
//             />
//           ))}
//         </SortableContext>
//       </DndContext>
//     </>
//   );
// }



"use client";

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
import { List, Trash, SkipForward } from "phosphor-react";

function SortableItem({ index, song, next, deleteSong }) {
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
      className="bg-neutral-900 flex justify-between items-center rounded px-4 py-3 mb-2 cursor-grab"
    >
      <span>
        {index + 1}. {song.title}
      </span>

      <span className="flex gap-4">
        {index === 0 ? (
          <SkipForward
            onClick={() => next()}
            className="text-white bg-green-600 hover:bg-green-500 p-2 rounded-full cursor-pointer transition-colors duration-200"
            size={32}
          />
        ) : (
          <Trash
            onClick={() => deleteSong(song.id)}
            className="text-white bg-red-600 hover:bg-red-500 p-2 rounded-full cursor-pointer transition-colors duration-200"
            size={32}
          />
        )}

        <List
          size={32}
          className="text-neutral-400 hover:text-white cursor-pointer transition-colors duration-200"
        />
      </span>
    </div>
  );
}

export default function PlaylistPanel({ playlist, setPlaylist, onReorder, next, deleteSong }) {
  const handleDragEnd = (e) => {
    const { active, over } = e;
    if (!over || active.id === over.id) return;

    setPlaylist((items) => {
      const oldIndex = items.findIndex((i) => i.id === active.id);
      const newIndex = items.findIndex((i) => i.id === over.id);
      const newItems = arrayMove(items, oldIndex, newIndex);

      // Update order_number based on new index
      const updatedItems = newItems.map((item, idx) => ({
        ...item,
        order_number: idx + 1,
      }));

      if (onReorder) onReorder(updatedItems); // notify parent
      return updatedItems;
    });
  };

  return (
    <>
      <h2 className="text-lg font-semibold mb-4">Live Playlist</h2>
      <DndContext collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
        <SortableContext
          items={playlist.map((p) => p.id)}
          strategy={verticalListSortingStrategy}
        >
          {playlist.map((song, index) => (
            <SortableItem
              key={index}
              index={index}
              song={song}
              next={next}
              deleteSong={deleteSong}
            />
          ))}
        </SortableContext>
      </DndContext>
    </>
  );
}
