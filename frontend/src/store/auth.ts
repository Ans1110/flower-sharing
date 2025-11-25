import { UserType } from "@/types/user";
import createStore from "@/lib/createStore";

type AuthStateType = {
  user: UserType | null;
  isAuthenticated: boolean;
  login: (user: UserType, accessToken: string) => void;
  register: (user: UserType, accessToken: string) => void;
  logout: () => void;
};

const useAuthStore = createStore<AuthStateType>(
  (set) => ({
    user: null,
    isAuthenticated: false,
    login: (user: UserType, accessToken: string) =>
      set((state) => {
        // Store token in localStorage
        if (globalThis.window !== undefined) {
          localStorage.setItem("accessToken", accessToken);
        }
        state.user = user;
        state.isAuthenticated = true;
      }),
    register: (user: UserType, accessToken: string) =>
      set((state) => {
        // Store token in localStorage
        if (globalThis.window !== undefined) {
          localStorage.setItem("accessToken", accessToken);
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
