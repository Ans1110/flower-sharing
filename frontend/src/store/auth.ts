import { redirect } from "react-router";
import { toast } from "sonner";
import { create } from "zustand";
import type { UserType } from "@/types/auth";

type AuthStateType = {
  token: string | null;
  user: UserType | null;
  setToken: (token: string | null) => void;
  setUser: (user: UserType | null) => void;
  logout: () => void;
};

const useAuthStore = create<AuthStateType>((set) => ({
  token: null,
  user: null,
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
    set({ token: null, user: null });
    toast.success("Logout successful");
    localStorage.removeItem("role");
    redirect("/login");
  },
}));

export { useAuthStore };
