import { api } from "./api";
import { useAuthStore } from "@/store/auth";
import { isTokenExpired, getTimeUntilExpiration } from "@/lib/jwt";

interface QueueItem {
  resolve: (value: string) => void;
  reject: (reason?: Error) => void;
}

// flag to prevent multiple simultaneous refresh requests
let isRefreshing = false;
let failedQueue: QueueItem[] = [];
let refreshTimer: NodeJS.Timeout | null = null;

const processQueue = (error: Error | null, token: string | null) => {
  for (const prom of failedQueue) {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token as string);
    }
  }
  failedQueue = [];
};

/**
 * Refresh the access token proactively
 */
async function refreshAccessToken(): Promise<string | null> {
  if (isRefreshing) {
    return null;
  }

  isRefreshing = true;

  try {
    const response = await api.post("/auth/refresh-token");
    const { accessToken } = response.data;

    if (globalThis.window !== undefined) {
      localStorage.setItem("accessToken", accessToken);
    }

    api.defaults.headers.common.Authorization = `Bearer ${accessToken}`;

    // Schedule next refresh
    scheduleTokenRefresh(accessToken);

    return accessToken;
  } catch (error) {
    if (globalThis.window !== undefined) {
      useAuthStore.getState().logout();
      globalThis.window.location.href = "/login";
    }

    return Promise.reject(error);
  } finally {
    isRefreshing = false;
  }
}

/**
 * Schedule token refresh before expiration
 * Refreshes 30 seconds before token expires (or halfway through if token lifetime < 60s)
 */
function scheduleTokenRefresh(token: string) {
  // Clear existing timer
  if (refreshTimer) {
    clearTimeout(refreshTimer);
    refreshTimer = null;
  }

  const timeUntilExpiration = getTimeUntilExpiration(token);

  if (timeUntilExpiration <= 0) {
    return;
  }

  // Refresh 30 seconds before expiration, or halfway through if token lifetime < 60s
  const refreshBuffer = Math.min(30, Math.floor(timeUntilExpiration / 2));
  const refreshIn = (timeUntilExpiration - refreshBuffer) * 1000; // Convert to milliseconds

  refreshTimer = setTimeout(() => {
    refreshAccessToken();
  }, refreshIn);
}

/**
 * Initialize proactive token refresh
 */
function initializeTokenRefresh() {
  if (globalThis.window === undefined) {
    return;
  }

  const accessToken = localStorage.getItem("accessToken");
  if (!accessToken) {
    return;
  }

  // Check if token is already expired
  if (isTokenExpired(accessToken)) {
    refreshAccessToken();
    return;
  }

  // Schedule refresh
  scheduleTokenRefresh(accessToken);
}

// Initialize on module load
if (globalThis.window !== undefined) {
  initializeTokenRefresh();

  // Re-initialize when page becomes visible (handles tab switching)
  document.addEventListener("visibilitychange", () => {
    if (!document.hidden) {
      const accessToken = localStorage.getItem("accessToken");
      if (accessToken && isTokenExpired(accessToken)) {
        refreshAccessToken();
      }
    }
  });
}

// Request interceptor - add access token to requests and check expiration
api.interceptors.request.use(async (config) => {
  if (globalThis.window !== undefined) {
    const accessToken = localStorage.getItem("accessToken");
    if (accessToken) {
      // Check if token is expired or will expire in next 5 seconds
      if (isTokenExpired(accessToken, 5)) {
        const newToken = await refreshAccessToken();
        if (newToken) {
          config.headers.Authorization = `Bearer ${newToken}`;
        }
      } else {
        config.headers.Authorization = `Bearer ${accessToken}`;
      }
    }
  }
  return config;
});

// Response interceptor - handle token refresh on 401
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // if error is 401 and we haven't tried to refresh yet
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      if (isRefreshing) {
        return new Promise<string>((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then((token) => {
            originalRequest.headers.Authorization = `Bearer ${token}`;
            return api(originalRequest);
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        // call refresh token endpoint
        const response = await api.post("/auth/refresh-token");
        const { accessToken } = response.data;

        if (globalThis.window !== undefined) {
          localStorage.setItem("accessToken", accessToken);
        }

        api.defaults.headers.common.Authorization = `Bearer ${accessToken}`;
        originalRequest.headers.Authorization = `Bearer ${accessToken}`;

        // Schedule next refresh
        scheduleTokenRefresh(accessToken);

        processQueue(null, accessToken);
        return api(originalRequest);
      } catch (refreshError) {
        // if refresh token fails, clear auth state and redirect to login
        processQueue(refreshError as Error, null);

        if (globalThis.window !== undefined) {
          // Clear auth store state
          useAuthStore.getState().logout();
          // redirect to login page
          globalThis.window.location.href = "/login";
        }

        return Promise.reject(refreshError);
      } finally {
        isRefreshing = false;
      }
    }
    return Promise.reject(error);
  }
);

// Export for use in auth store
export { scheduleTokenRefresh };
