import { useAuthStore } from "@/store/auth";
import { Navigate, Outlet } from "react-router";

const ProtectedRoute = () => {
  const token = useAuthStore((state) => state.token);
  const isInitialized = useAuthStore((state) => state.isInitialized);

  if (!isInitialized) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-rose-500"></div>
      </div>
    );
  }

  if (!token) {
    return <Navigate to="/login" replace />;
  }
  return <Outlet />;
};

export default ProtectedRoute;
