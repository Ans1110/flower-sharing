import { UserType } from "@/types/user";
import createStore from "@/lib/createStore";
import { scheduleTokenRefresh } from "@/service/refreshToken";

type AuthStateType = {
  user: UserType | null;
  isAuthenticated: boolean;
};

type AuthActionType = {
  login: (user: UserType, accessToken: string) => void;
  register: (user: UserType, accessToken: string) => void;
  logout: () => void;
  validateAuth: () => void;
  updateUser: (updates: Partial<UserType>) => void;
};

type AuthStoreType = AuthStateType & AuthActionType;

const useAuthStore = createStore<AuthStoreType>(
  (set) => ({
    user: null,
    isAuthenticated: false,
    login: (user: UserType, accessToken: string) =>
      set((state) => {
        // Store token in localStorage
        if (globalThis.window !== undefined) {
          localStorage.setItem("accessToken", accessToken);
          // Schedule proactive token refresh
          scheduleTokenRefresh(accessToken);
        }
        state.user = user;
        state.isAuthenticated = true;
      }),
    register: (user: UserType, accessToken: string) =>
      set((state) => {
        // Store token in localStorage
        if (globalThis.window !== undefined) {
          localStorage.setItem("accessToken", accessToken);
          // Schedule proactive token refresh
          scheduleTokenRefresh(accessToken);
        }
        state.user = user;
        state.isAuthenticated = true;
      }),
    logout: () =>
      set((state) => {
        if (globalThis.window !== undefined) {
          localStorage.removeItem("accessToken");
        }
        state.user = null;
        state.isAuthenticated = false;
      }),
    validateAuth: () =>
      set((state) => {
        // Check if accessToken exists in localStorage
        if (globalThis.window !== undefined) {
          const accessToken = localStorage.getItem("accessToken");
          // If no token but user is authenticated, clear the auth state
          if (!accessToken && state.isAuthenticated) {
            state.user = null;
            state.isAuthenticated = false;
          }
        }
      }),
    updateUser: (updates: Partial<UserType>) =>
      set((state) => {
        if (state.user) {
          state.user = { ...state.user, ...updates };
        }
      }),
  }),
  {
    name: "auth-store",
    storage:
      globalThis.window === undefined
        ? undefined
        : globalThis.window.localStorage,
  }
);

export { useAuthStore };
