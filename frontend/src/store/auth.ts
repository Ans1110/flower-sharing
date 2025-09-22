import { create } from "zustand";

type AuthStateType = {
  token: string | null;
  role: string | null;
  setToken: (token: string | null) => void;
  logout: () => void;
};

const useAuthStore = create<AuthStateType>((set) => ({
  token: null,
  role: null,
  setToken: (token) => {
    if (token) {
      localStorage.setItem("token", token);
    } else {
      localStorage.removeItem("token");
    }
    set({ token });
  },
  logout: () => {
    localStorage.removeItem("token");
    set({ token: null, role: null });
  },
}));

export { useAuthStore };
