import axios from "axios";

export const API = axios.create({
  baseURL: process.env.END_POINT||  'http://localhost:8080/api/v1'
  // baseURL: process.env.API_URL,
})
