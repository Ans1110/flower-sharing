import { useAuthStore } from "@/store/auth";
import { useEffect } from "react";

const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const initialize = useAuthStore((state) => state.initialize);

  useEffect(() => {
    initialize();
  }, [initialize]);

  return <>{children}</>;
};

export default AuthProvider;
