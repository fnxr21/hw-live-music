import { API } from "./axios";

// export const APILogin = (form) => API.post("/login", form);
// export const APIPosname = (form) => API.patch("reauth", form);

// admin
export const ApiCreateSong = (form) => API.post("/song", form);
export const ApiListSongs = ({ page = 1, limit = 5 }) =>
  API.get("/songs", { params: { page, limit } });
export const ApiGetSong = (id) => API.get(`/song/${id}`);
export const ApiUpdateSong = (form) => API.put(`/song/${form.id}`, form);
export const ApiDeleteSong = (id) => API.delete(`/song/${id}`);

// admin 
export const ApiListSongsRequest = ({ page = 1, limit = 5 }) =>
  API.get("/song-requests", { params: { page, limit } });
export const ApiUpdateSongRequest = (form) =>
  API.put(`/song-request/${form.id}`, form);
