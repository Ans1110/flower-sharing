"use client";

import { Suspense, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { useAuthStore } from "@/store/auth";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";

function OAuthCallbackContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const login = useAuthStore((state) => state.login);

  useEffect(() => {
    const handleOAuthCallback = async () => {
      const accessToken = searchParams.get("access_token");
      const refreshToken = searchParams.get("refresh_token");
      const error = searchParams.get("error");

      if (error) {
        const errorMessages: Record<string, string> = {
          invalid_state: "Invalid state token. Please try again.",
          no_code: "No authorization code received. Please try again.",
          token_exchange_failed: "Failed to exchange token. Please try again.",
          user_info_failed: "Failed to get user information. Please try again.",
          parse_failed: "Failed to parse user information. Please try again.",
          user_creation_failed:
            "Failed to create user account. Please try again.",
          token_generation_failed:
            "Failed to generate authentication token. Please try again.",
          token_save_failed:
            "Failed to save authentication token. Please try again.",
          no_email:
            "No email found in your account. Please make sure your email is public or try another method.",
        };

        toast.error(
          errorMessages[error] || "Authentication failed. Please try again."
        );
        router.push("/login");
        return;
      }

      if (!accessToken || !refreshToken) {
        toast.error("Authentication failed. Missing tokens.");
        router.push("/login");
        return;
      }

      try {
        // Fetch user details from the /auth/me endpoint
        const response = await fetch(
          `${
            process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1"
          }/auth/me`,
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
          }
        );

        if (!response.ok) {
          throw new Error("Failed to fetch user details");
        }

        const data = await response.json();

        // Store refresh token
        localStorage.setItem("refreshToken", refreshToken);

        // Store tokens and user data (API returns { user: {...} })
        login(data.user, accessToken);

        toast.success("Successfully logged in!");
        router.push("/");
      } catch (error) {
        console.error("OAuth callback error:", error);
        toast.error("Failed to complete authentication. Please try again.");
        router.push("/login");
      }
    };

    handleOAuthCallback();
  }, [searchParams, login, router]);

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
