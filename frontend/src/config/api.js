import { API } from "./axios";

// export const APILogin = (form) => API.post("/login", form);
// export const APIPosname = (form) => API.patch("reauth", form);


// admin
export const APIGetSong = () => API.get("/song");
export const ApiDeletesong = () => API.delete("/song");
export const APICreateSong = () => API.get("/song");
export const ApiUpdateesong = () => API.delete("/song");
// export const ApiListSongs = ( { params: { page, limit } }) => API.delete("/song", { params: { page, limit } });

export const ApiListSongs = ({ page = 1, limit = 5 }) => 
    API.get("/song-requests", { params: { page, limit } });

export const APIConfigLocation = (form) => API.patch("update-config-location", form);

export const APIGetConfig = () => API.get("/config-location");
