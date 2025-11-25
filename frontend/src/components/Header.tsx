"use client";

import { Flower, LogIn, LogOut, LucideIcon, Menu, X } from "lucide-react";
import Link from "next/link";
import { Button } from "./ui/button";
import { DropdownAvatar } from "./ui/dropdown-avatar";
import { UserType } from "@/types/user";
import { ThemeToggle } from "./ThemeToggle";

type HeaderProps = {
  navItems: {
    path: string;
    label: string;
    icon: LucideIcon;
  }[];
  user: UserType | null;
  isAuthenticated: boolean;
  handleLogout: () => void;
  setIsMobileMenuOpen: (isOpen: boolean) => void;
  isMobileMenuOpen: boolean;
  isActive: (path: string) => boolean;
};

const Header = ({
  navItems,
  user,
  isAuthenticated,
  handleLogout,
  setIsMobileMenuOpen,
  isMobileMenuOpen,
  isActive,
}: HeaderProps) => {
  return (
    <header className="sticky top-0 z-50 border-b backdrop-blur-sm bg-white/80 dark:bg-black/80 border-rose-100 dark:border-rose-900 shadow-lg ">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <div className="flex items-center">
            <Link href="/" className="flex items-center space-x-2">
              <div className="w-10 h-10 bg-linear-to-r from-rose-500 to-pink-500 rounded-full flex items-center justify-center">
                <Flower className="text-white" />
              </div>
              <span className="text-xl font-bold bg-linear-to-r from-rose-500 to-pink-500 bg-clip-text text-transparent">
                FlowerShare
              </span>
            </Link>
          </div>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            {navItems.map(({ path, label, icon: Icon }) => (
              <Link
                href={path}
                key={path}
                className={`flex items-center space-x-2 px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200 ${
                  isActive(path)
                    ? "bg-rose-100 dark:bg-rose-900 text-rose-700 dark:text-rose-300 shadow-sm"
                    : "text-gray-600 dark:text-gray-400 hover:bg-rose-50 dark:hover:bg-rose-800/50 hover:text-rose-600 dark:hover:text-rose-400"
                }`}
              >
                <Icon className="size-5" />
                <span>{label}</span>
              </Link>
            ))}
          </nav>

          {/* Desktop avatar and logout button */}
          <div className="hidden md:flex items-center space-x-3">
            <ThemeToggle />
            {isAuthenticated ? (
              <DropdownAvatar
                username={user?.username}
                email={user?.email}
                avatar={user?.avatar}
                onSignOut={handleLogout}
              />
            ) : (
              <div className="flex items-center space-x-3">
                <Button
                  asChild
                  variant="outline"
                  className=" border-rose-200 text-rose-600 hover:bg-rose-50 hover:text-rose-700 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
                >
                  <Link href="/login" className="flex items-center space-x-2">
                    <LogIn className="size-5" />
                    <span className="text-sm">Login</span>
                  </Link>
                </Button>
              </div>
            )}
          </div>

          {/* Mobile menu */}
          <div className="md:hidden">
            <Button
              asChild
              variant="ghost"
              size="sm"
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
            >
              {isMobileMenuOpen ? (
                <X className="w-6 h-6" />
              ) : (
                <Menu className="w-6 h-6" />
              )}
            </Button>
          </div>
        </div>

        {/* Mobile Navigation */}
        {isMobileMenuOpen && (
          <div className="md:hidden border-t border-rose-100 dark:border-rose-900 py-4">
            <nav className="flex flex-col space-y-4">
              {navItems.map(({ path, label, icon: Icon }) => (
                <Link
                  href={path}
                  key={path}
                  onClick={() => setIsMobileMenuOpen(false)}
                  className={`flex items-center space-x-3 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 ${
                    isActive(path)
                      ? "bg-rose-100 dark:bg-rose-900 text-rose-700 dark:text-rose-300 shadow-sm"
                      : "text-gray-600 dark:text-gray-400 hover:bg-rose-50 dark:hover:bg-rose-800/50 hover:text-rose-600 dark:hover:text-rose-400"
                  }`}
                >
                  <Icon className="size-5" />
                  <span>{label}</span>
                </Link>
              ))}
              <div className="pt-4 border-t border-rose-100 dark:border-rose-900">
                {isAuthenticated ? (
                  <Button
                    variant="outline"
                    onClick={handleLogout}
                    className="w-full border-rose-200 text-rose-600 hover:bg-rose-50 hover:text-rose-700 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
                  >
                    <LogOut className="size-5" />
                    <span className="sr-only">Sign out</span>
                  </Button>
                ) : (
                  <Button
                    asChild
                    variant="outline"
                    className="w-full border-rose-200 text-rose-600 hover:bg-rose-50 hover:text-rose-700 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
                  >
                    <Link
                      href="/auth/login"
                      onClick={() => setIsMobileMenuOpen(false)}
                      className="flex items-center space-x-2 w-full justify-center"
                    >
                      <LogIn className="size-5" />
                      <span className="sr-only">Login</span>
                    </Link>
                  </Button>
                )}
              </div>
            </nav>
          </div>
        )}
      </div>
    </header>
  );
};

export { Header };
