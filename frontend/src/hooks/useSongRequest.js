import { useState, useEffect } from "react";
// import { listSongRequests, updateSongRequest } from "@/config/api/requests";
import { ApiListSongsRequest, ApiUpdateSongRequest } from "@/config/api";

export const useSongRequests = (perPage = 5) => {
  const [requests, setRequests] = useState([]);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  const fetchRequests = async () => {
    setLoading(true);
    try {
      const { data } = await ApiListSongsRequest({ page, limit: perPage });
      setRequests(data.data || []);
      setTotal(data.total || 0);
    } catch (err) {
      console.error("Failed to fetch requests:", err);
    } finally {
      setLoading(false);
    }
  };

  const updateStatus = async (req, statusUuid) => {
    await ApiUpdateSongRequest({ id: req.song_request_id, status: statusUuid });
    fetchRequests();
  };

  useEffect(() => {
    fetchRequests();
  }, [page]);

  return { requests, page, setPage, total, loading, updateStatus };
};
