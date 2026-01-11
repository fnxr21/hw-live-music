import { useState, useEffect } from "react";
import { fetchSongRequests, updateSongRequestStatus } from "@/services/songRequests";

export const useSongRequests = (perPage = 5) => {
  const [requests, setRequests] = useState([]);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);

  const loadRequests = async () => {
    setLoading(true);
    try {
      const data = await fetchSongRequests({ page, limit: perPage });
      setRequests(data.data || []);
      setTotal(data.total || 0);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const updateStatus = async (req, statusUuid) => {
    await updateSongRequestStatus(req.song_request_id, statusUuid);
    loadRequests();
  };

  useEffect(() => {
    loadRequests();
  }, [page]);

  return { requests, page, setPage, total, loading, updateStatus };
};
