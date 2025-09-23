import { useAuthStore } from "@/store/auth";
import { Navigate, Outlet } from "react-router";

const ProtectedRoute = () => {
  const token = useAuthStore((state) => state.token);

  if (!token) {
    return <Navigate to="/login" replace />;
  }
  return <Outlet />;
};

export default ProtectedRoute;
