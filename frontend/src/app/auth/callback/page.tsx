"use client";

import { Suspense, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { oauthErrorMessages } from "@/lib/errorMap";
import { useAuthStore } from "@/store/auth";
import { useGetMe } from "@/hooks/api/user";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";

function OAuthCallbackContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const login = useAuthStore((state) => state.login);

  const accessToken = searchParams.get("access_token");
  const refreshToken = searchParams.get("refresh_token");
  const error = searchParams.get("error");

  const { refetch: fetchMe } = useGetMe({
    enabled: false,
    retry: 1,
  });

  useEffect(() => {
    const handleOAuthCallback = async () => {
      if (error) {
        toast.error(
          oauthErrorMessages[error] ||
            "Authentication failed. Please try again."
        );
        router.push("/login");
        return;
      }

      if (!accessToken || !refreshToken) {
        toast.error("Authentication failed. Missing tokens.");
        router.push("/login");
        return;
      }

      localStorage.setItem("accessToken", accessToken);
      localStorage.setItem("refreshToken", refreshToken);

      try {
        const { data: user } = await fetchMe({ throwOnError: true });

        if (!user) {
          throw new Error("Failed to fetch user details");
        }

        login(user, accessToken);
        toast.success("Successfully logged in!");
        router.push("/");
      } catch (callbackError) {
        console.error("OAuth callback error:", callbackError);
        localStorage.removeItem("accessToken");
        localStorage.removeItem("refreshToken");
        toast.error("Failed to complete authentication. Please try again.");
        router.push("/login");
      }
    };

    handleOAuthCallback();
  }, [accessToken, refreshToken, error, router, fetchMe, login]);

  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 flex items-center justify-center">
      <div className="text-center">
        <Loader2 className="size-12 animate-spin text-rose-600 dark:text-rose-400 mx-auto mb-4" />
        <h2 className="text-2xl font-semibold text-gray-900 dark:text-gray-100 mb-2">
          Completing authentication...
        </h2>
        <p className="text-gray-600 dark:text-gray-400">
          Please wait while we log you in
        </p>
      </div>
    </div>
  );
}

function LoadingFallback() {
  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 flex items-center justify-center">
      <div className="text-center">
        <Loader2 className="size-12 animate-spin text-rose-600 dark:text-rose-400 mx-auto mb-4" />
        <h2 className="text-2xl font-semibold text-gray-900 dark:text-gray-100 mb-2">
          Loading...
        </h2>
      </div>
    </div>
  );
}

export default function OAuthCallback() {
  return (
    <Suspense fallback={<LoadingFallback />}>
      <OAuthCallbackContent />
    </Suspense>
  );
}
