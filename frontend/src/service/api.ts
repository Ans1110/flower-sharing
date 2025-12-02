import axios from "axios";

/**
 * api/v1
 */
const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  timeout: 10000,
  withCredentials: true,
});

// Add request interceptor to set Content-Type appropriately
api.interceptors.request.use((config) => {
  // Don't set Content-Type for FormData - let browser set it with boundary
  if (config.data instanceof FormData) {
    // Remove any existing Content-Type to let browser set multipart/form-data with boundary
    delete config.headers["Content-Type"];
  } else if (!config.headers["Content-Type"]) {
    // Set default Content-Type for non-FormData requests
    config.headers["Content-Type"] = "application/json";
  }
  return config;
});

export { api };
