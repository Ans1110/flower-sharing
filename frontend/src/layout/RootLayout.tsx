import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/store/auth";
import { Outlet, NavLink, useNavigate } from "react-router-dom";
import {
  Menu,
  X,
  Flower,
  Home,
  User,
  LogOut,
  LogIn,
  UserPlus,
} from "lucide-react";
import { useState } from "react";

const RootLayout = () => {
  const token = useAuthStore((state) => state.token);
  const logout = useAuthStore((state) => state.logout);
  const navigate = useNavigate();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const handleLogout = () => {
    logout();
    navigate("/login");
    setIsMobileMenuOpen(false);
  };

  const navItems = [
    { path: "/", label: "Home", icon: Home },
    { path: "/flowers", label: "Flowers", icon: Flower },
    ...(token ? [{ path: "/profile", label: "Profile", icon: User }] : []),
  ];

  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50">
      {/* Header */}
      <header className="bg-white/80 backdrop-blur-sm border-b border-rose-100 shadow-lg sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            {/* Logo */}
            <div className="flex items-center">
              <NavLink to="/" className="flex items-center space-x-2">
                <div className="w-10 h-10 bg-gradient-to-r from-rose-500 to-pink-500 rounded-full flex items-center justify-center">
                  <Flower className="w-6 h-6 text-white" />
                </div>
                <span className="text-xl font-bold bg-gradient-to-r from-rose-600 to-pink-600 bg-clip-text text-transparent">
                  FlowerShare
                </span>
              </NavLink>
            </div>

            {/* Desktop Navigation */}
            <nav className="hidden md:flex items-center space-x-8">
              {navItems.map(({ path, label, icon: Icon }) => (
                <NavLink
                  key={path}
                  to={path}
                  className={({ isActive }) =>
                    `flex items-center space-x-2 px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200 ${
                      isActive
                        ? "bg-rose-100 text-rose-700 shadow-sm"
                        : "text-gray-600 hover:text-rose-600 hover:bg-rose-50"
                    }`
                  }
                >
                  <Icon className="w-4 h-4" />
                  <span>{label}</span>
                </NavLink>
              ))}
            </nav>

            {/* Desktop Auth */}
            <div className="hidden md:flex items-center space-x-4">
              {token ? (
                <Button
                  variant="outline"
                  onClick={handleLogout}
                  className="border-rose-200 text-rose-600 hover:bg-rose-50 flex items-center space-x-2"
                >
                  <LogOut className="w-4 h-4" />
                  <span>Logout</span>
                </Button>
              ) : (
                <div className="flex items-center space-x-3">
                  <Button
                    asChild
                    variant="outline"
                    className="border-rose-200 text-rose-600 hover:bg-rose-50"
                  >
                    <NavLink
                      to="/login"
                      className="flex items-center space-x-2"
                    >
                      <LogIn className="w-4 h-4" />
                      <span>Sign In</span>
                    </NavLink>
                  </Button>
                  <Button
                    asChild
                    className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                  >
                    <NavLink
                      to="/register"
                      className="flex items-center space-x-2"
                    >
                      <UserPlus className="w-4 h-4" />
                      <span>Sign Up</span>
                    </NavLink>
                  </Button>
                </div>
              )}
            </div>

            {/* Mobile menu button */}
            <div className="md:hidden">
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                className="text-gray-600 hover:text-rose-600"
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
            <div className="md:hidden border-t border-rose-100 py-4">
              <nav className="flex flex-col space-y-2">
                {navItems.map(({ path, label, icon: Icon }) => (
                  <NavLink
                    key={path}
                    to={path}
                    onClick={() => setIsMobileMenuOpen(false)}
                    className={({ isActive }) =>
                      `flex items-center space-x-3 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 ${
                        isActive
                          ? "bg-rose-100 text-rose-700"
                          : "text-gray-600 hover:text-rose-600 hover:bg-rose-50"
                      }`
                    }
                  >
                    <Icon className="w-5 h-5" />
                    <span>{label}</span>
                  </NavLink>
                ))}
                <div className="pt-4 border-t border-rose-100">
                  {token ? (
                    <Button
                      variant="outline"
                      onClick={handleLogout}
                      className="w-full border-rose-200 text-rose-600 hover:bg-rose-50 flex items-center justify-center space-x-2"
                    >
                      <LogOut className="w-4 h-4" />
                      <span>Logout</span>
                    </Button>
                  ) : (
                    <div className="flex flex-col space-y-2">
                      <Button
                        asChild
                        variant="outline"
                        className="w-full border-rose-200 text-rose-600 hover:bg-rose-50"
                      >
                        <NavLink
                          to="/login"
                          onClick={() => setIsMobileMenuOpen(false)}
                        >
                          Sign In
                        </NavLink>
                      </Button>
                      <Button
                        asChild
                        className="w-full bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                      >
                        <NavLink
                          to="/register"
                          onClick={() => setIsMobileMenuOpen(false)}
                        >
                          Sign Up
                        </NavLink>
                      </Button>
                    </div>
                  )}
                </div>
              </nav>
            </div>
          )}
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1">
        <Outlet />
      </main>

      {/* Footer */}
      <footer className="bg-white/80 backdrop-blur-sm border-t border-rose-100 py-8">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <div className="flex items-center justify-center space-x-2 mb-4">
              <div className="w-8 h-8 bg-gradient-to-r from-rose-500 to-pink-500 rounded-full flex items-center justify-center">
                <Flower className="w-5 h-5 text-white" />
              </div>
              <span className="text-lg font-bold bg-gradient-to-r from-rose-600 to-pink-600 bg-clip-text text-transparent">
                FlowerShare
              </span>
            </div>
            <p className="text-gray-600 text-sm">
              &copy; {new Date().getFullYear()} Flower Sharing Platform. All
              rights reserved.
            </p>
            <p className="text-gray-500 text-xs mt-2">
              Share the beauty of flowers with our community
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default RootLayout;
