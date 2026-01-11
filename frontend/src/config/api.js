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



export const APIConfigLocation = (form) =>
  API.patch("update-config-location", form);

export const APIGetConfig = () => API.get("/config-location");
