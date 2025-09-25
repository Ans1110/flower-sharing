import { redirect } from "react-router";
import { toast } from "sonner";
import { create } from "zustand";
import type { UserType } from "@/types/auth";

type AuthStateType = {
  token: string | null;
  user: UserType | null;
  isInitialized: boolean;
  setToken: (token: string | null) => void;
  setUser: (user: UserType | null) => void;
  logout: () => void;
  initialize: () => void;
};

const useAuthStore = create<AuthStateType>((set) => ({
  token: null,
  user: null,
  isInitialized: false,
  setToken: (token) => {
    if (token) {
      localStorage.setItem("token", token);
    } else {
      localStorage.removeItem("token");
    }
    set({ token });
  },
  setUser: (user) => {
    if (user) {
      localStorage.setItem("user", JSON.stringify(user));
    } else {
      localStorage.removeItem("user");
    }
    set({ user });
  },
  logout: () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    localStorage.removeItem("role");
    set({ token: null, user: null });
    toast.success("Logout successful");
    redirect("/login");
  },
  initialize: () => {
    const token = localStorage.getItem("token");
    const userStr = localStorage.getItem("user");
    const user = userStr ? JSON.parse(userStr) : null;

    set({
      token,
      user,
      isInitialized: true,
    });
  },
}));

export { useAuthStore };
