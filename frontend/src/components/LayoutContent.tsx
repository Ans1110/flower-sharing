"use client";

import { BookOpenText, Flower, Home, User, Users } from "lucide-react";
import { useAuthStore } from "@/store/auth";
import { useState, useEffect } from "react";
import { toast } from "sonner";
import { redirect, usePathname } from "next/navigation";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";

export default function LayoutContent({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const user = useAuthStore((state) => state.user);
  const logout = useAuthStore((state) => state.logout);
  const validateAuth = useAuthStore((state) => state.validateAuth);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const pathname = usePathname();
  const isActive = (path: string) =>
    pathname === path || (pathname.startsWith(path) && path !== "/");

  // Validate auth state on mount and when localStorage changes
  useEffect(() => {
    validateAuth();

    // Listen for storage changes (e.g., when token is deleted in another tab)
    const handleStorageChange = (e: StorageEvent) => {
      if (e.key === "accessToken" && e.newValue === null) {
        validateAuth();
      }
    };

    globalThis.window.addEventListener("storage", handleStorageChange);
    return () =>
      globalThis.window.removeEventListener("storage", handleStorageChange);
  }, [validateAuth]);

  const handleLogout = () => {
    logout();
    setIsMobileMenuOpen(false);
    toast.success("Logged out successfully");
    redirect("/login");
  };

  const navItems = [
    { path: "/", label: "Home", icon: Home },
    { path: "/flowers", label: "Flowers", icon: Flower },
    ...(isAuthenticated
      ? [{ path: `/profile/${user?.id}`, label: "Profile", icon: User }]
      : []),
  ];

  const isAdmin = user?.role === "admin";
  if (isAdmin) {
    navItems.push(
      { path: "/admin/users", label: "Users", icon: Users },
      { path: "/admin/posts", label: "Posts", icon: BookOpenText }
    );
  }

  return (
    <div className="min-h-screen flex flex-col bg-linear-to-br from-rose-50 via-pink-50 to-purple-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950">
      <Header
        navItems={navItems}
        user={user}
        isAuthenticated={isAuthenticated}
        handleLogout={handleLogout}
        setIsMobileMenuOpen={setIsMobileMenuOpen}
        isMobileMenuOpen={isMobileMenuOpen}
        isActive={isActive}
      />
      <main className="flex-1">{children}</main>
      <Footer />
    </div>
  );
}

export { LayoutContent };
