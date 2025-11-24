"use client";

import { Flower, Home, User } from "lucide-react";
import { useAuthStore } from "@/store/auth";
import { useState } from "react";
import { toast } from "sonner";
import { redirect, usePathname } from "next/navigation";
import { Header } from "@/components/Header";

export default function LayoutContent({
  children,
}: {
  children: React.ReactNode;
}) {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const logout = useAuthStore((state) => state.logout);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const pathname = usePathname();
  const isActive = (path: string) => pathname === path;

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
      ? [{ path: "/profile", label: "Profile", icon: User }]
      : []),
  ];

  const user = useAuthStore((state) => state.user);

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
    </div>
  );
}

export { LayoutContent };
