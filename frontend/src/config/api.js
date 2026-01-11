import { API } from "./axios";

// export const APILogin = (form) => API.post("/login", form);
// export const APIPosname = (form) => API.patch("reauth", form);

// admin
export const APIGetSong = () => API.get("/song");
export const ApiDeletesong = () => API.delete("/song");
export const APICreateSong = () => API.get("/song");
export const ApiUpdateesong = () => API.delete("/song");
export const ApiListSongs = ({ page = 1, limit = 5, tableId = 0 }) =>
  API.get("/songs", { params: { page, limit, tableId } });


export const APIConfigLocation = (form) =>
  API.patch("update-config-location", form);

export const APIGetConfig = () => API.get("/config-location");
