import { redirect } from "react-router";
import { toast } from "sonner";
import { create } from "zustand";

type AuthStateType = {
  token: string | null;
  role: string | null;
  setToken: (token: string | null) => void;
  logout: () => void;
  userId: number | null;
  setUserId: (userId: number | null) => void;
};

const useAuthStore = create<AuthStateType>((set) => ({
  token: null,
  role: null,
  userId: null,
  setToken: (token) => {
    if (token) {
      localStorage.setItem("token", token);
    } else {
      localStorage.removeItem("token");
    }
    set({ token });
  },
  setUserId: (userId) => {
    if (userId) {
      localStorage.setItem("userId", userId.toString());
    } else {
      localStorage.removeItem("userId");
    }
    set({ userId });
  },
  logout: () => {
    localStorage.removeItem("token");
    set({ token: null, role: null, userId: null });
    toast.success("Logout successful");
    localStorage.removeItem("userId");
    redirect("/login");
  },
}));

export { useAuthStore };
