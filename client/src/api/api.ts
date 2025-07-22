import axios from "axios";

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL as string;
console.log(BACKEND_URL);
export const API = axios.create({ baseURL: BACKEND_URL });
