import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/store/auth";
import { Outlet, NavLink, useNavigate } from "react-router-dom";
const RootLayout = () => {
  const token = useAuthStore((state) => state.token);
  const logout = useAuthStore((state) => state.logout);
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <div className="min-h-screen flex flex-col">
      <header className="bg-pink-300 text-white flex justify-between p-4">
        <nav className="flex gap-4">
          <NavLink
            to="/"
            className={({ isActive }) =>
              `transition hover:text-rose-200 ${
                isActive ? "font-bold underline" : ""
              }`
            }
          >
            Home
          </NavLink>
          <NavLink
            to="/flowers"
            className={({ isActive }) =>
              `transition hover:text-rose-200 ${
                isActive ? "font-bold underline" : ""
              }`
            }
          >
            Flowers
          </NavLink>
          <NavLink
            to="/profile"
            className={({ isActive }) =>
              `transition hover:text-rose-200 ${
                isActive ? "font-bold underline" : ""
              }`
            }
          >
            Profile
          </NavLink>
        </nav>
        <div className="flex items-center gap-3">
          {token ? (
            <Button
              variant="secondary"
              className="bg-rose-400 hover:bg-rose-500 text-white"
              onClick={handleLogout}
            >
              Logout
            </Button>
          ) : (
            <NavLink
              to="/login"
              className="px-3 py-1 rounded-lg bg-rose-400 text-white hover:bg-rose-500 transition"
            >
              Login
            </NavLink>
          )}
        </div>
      </header>

      <main className="flex-1 p-6">
        <Outlet />
      </main>

      <footer className="bg-pink-200 text-pink-700 p-4 text-center text-sm">
        <p className="text-center text-sm">
          &copy; {new Date().getFullYear()} Flower Sharing Platform. All rights
          reserved.
        </p>
      </footer>
    </div>
  );
};

export default RootLayout;
