import { UserType } from "@/types/user";
import createStore from "@/lib/createStore";

type AuthStateType = {
  user: UserType | null;
  isAuthenticated: boolean;
  accessToken: string | null;
  setAccessToken: (accessToken: string) => void;
  setUser: (user: UserType) => void;
  logout: () => void;
};

const useAuthStore = createStore<AuthStateType>(
  (set) => ({
    user: null,
    isAuthenticated: false,
    accessToken: null,
    setAccessToken: (accessToken: string) =>
      set((state) => {
        if (accessToken) {
          localStorage.setItem("accessToken", accessToken);
          state.accessToken = accessToken;
          state.isAuthenticated = true;
        } else {
          localStorage.removeItem("accessToken");
          state.accessToken = null;
          state.isAuthenticated = false;
        }
      }),
    setUser: (user: UserType) =>
      set((state) => {
        state.user = user;
      }),
    logout: () =>
      set((state) => {
        localStorage.removeItem("accessToken");
        state.accessToken = null;
        state.isAuthenticated = false;
        state.user = null;
      }),
  }),
  {
    name: "auth-store",
    excludeFromPersist: ["accessToken"],
  }
);

export { useAuthStore };
