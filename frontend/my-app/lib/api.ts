import axios from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true, // Necessary if backend sets cookie on its domain
});

// We can add interceptors here later for global error handling
api.interceptors.response.use(
  (response) => response,
  (error) => {
    // e.g., redirect to login if 401
    return Promise.reject(error);
  }
);
