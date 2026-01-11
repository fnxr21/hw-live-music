// import { useState, useEffect } from "react";
// // import { listSongRequests, updateSongRequest } from "@/config/api/playlist";
// import { ApiListPlaylist, ApiUpdatePlaylist } from "@/config/api";

// export const useLivePlaylist = (perPage = 5) => {
  
//   const [playlist, setPlaylist] = useState([]);


//   const [page, setPage] = useState(1);
//   const [total, setTotal] = useState(0);
//   const [loading, setLoading] = useState(false);

//   const fetchRequests = async () => {
//     setLoading(true);
//     try {
//       const { data } = await ApiListPlaylist({ page, limit: perPage });
//       setPlaylist(data.data || []);
//       setTotal(data.total || 0);
//     } catch (err) {
//       console.error("Failed to fetch playlist:", err);
//     } finally {
//       setLoading(false);
//     }
//   };

//   const updatePlaylist = async (req, statusUuid) => {
//     await ApiUpdatePlaylist({ id: req.song_request_id, status: statusUuid });
//     fetchRequests();
//   };

//   useEffect(() => {
//     fetchRequests();
//   }, [page]);

//   return { playlist, page, setPage, total, loading, updatePlaylist };
// };



"use client";

import { useState, useEffect } from "react";
import { ApiListPlaylist, ApiUpdatePlaylist } from "@/config/api";

export const useLivePlaylist = (perPage = 5) => {
  const [playlist, setPlaylist] = useState([]);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  // Fetch playlist from backend
  const fetchRequests = async () => {
    setLoading(true);
    try {
      const { data } = await ApiListPlaylist({ page, limit: perPage });
      setPlaylist(data.data || []);
      setTotal(data.total || 0);
    } catch (err) {
      console.error("Failed to fetch playlist:", err);
    } finally {
      setLoading(false);
    }
  };

  // Update playlist order or status
  const updatePlaylist = async (song) => {
    try {
      await ApiUpdatePlaylist({
        id: song.song_request_id || song.id,
        order_number: song.order_number,
      });
      fetchRequests();
    } catch (err) {
      console.error("Failed to update playlist:", err);
    }
  };

  useEffect(() => {
    fetchRequests();
  }, [page]);

  return {
    playlist,
    setPlaylist,
    page,
    setPage,
    total,
    loading,
    updatePlaylist,
    fetchRequests,
  };
};
