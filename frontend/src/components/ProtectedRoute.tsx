"use client";

import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";

interface ProtectedRouteProps {
  children: React.ReactNode;
  requireAdmin?: boolean;
}

export default function ProtectedRoute({
  children,
  requireAdmin = false,
}: ProtectedRouteProps) {
  const router = useRouter();
  const pathname = usePathname();
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        // Check if user is authenticated
        // This should be replaced with your actual auth check logic
        const token = document.cookie
          .split("; ")
          .find((row) => row.startsWith("accessToken="))
          ?.split("=")[1];

        if (!token) {
          router.push(`/login?redirect=${pathname}`);
          return;
        }

        // If admin is required, check role
        if (requireAdmin) {
          const role = document.cookie
            .split("; ")
            .find((row) => row.startsWith("role="))
            ?.split("=")[1];

          if (role !== "admin") {
            router.push("/");
            return;
          }
        }

        setIsAuthorized(true);
      } catch (error) {
        console.error("Auth check failed:", error);
        router.push("/login");
      } finally {
        setIsLoading(false);
      }
    };

    checkAuth();
  }, [router, pathname, requireAdmin]);

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-lg">Loading...</div>
      </div>
    );
  }

  if (!isAuthorized) {
    return null;
  }

  return <>{children}</>;
}
