"use client";

import { useAuthStore } from "@/store/auth";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import z from "zod";
import { useMutation } from "@tanstack/react-query";
import { api } from "@/service/api";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import { Flower, Loader2, Lock, LogIn, Mail } from "lucide-react";
import { AxiosError } from "axios";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { ControlledInput } from "@/components/ui/controlled-input";
import { Button } from "@/components/ui/button";
import Link from "next/link";

const loginSchema = z.object({
  email: z
    .string()
    .email("Please enter a valid email address")
    .regex(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/, {
      message: "Invalid email address",
    }),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters")
    .max(255, "Password must be less than 255 characters"),
});

type LoginSchemaType = z.infer<typeof loginSchema>;

export default function Login() {
  const router = useRouter();
  const login = useAuthStore((state) => state.login);
  const form = useForm<LoginSchemaType>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = form;

  const mutation = useMutation({
    mutationFn: (data: LoginSchemaType) => api.post("/auth/login", data),
    onSuccess: ({ data }) => {
      login(data.user, data.accessToken);
      toast.success("Welcome back! Login successful");
      router.push("/");
    },
    onError: (error: AxiosError<{ error: string }>) => {
      const errorMessage =
        error.response?.data?.error || "Invalid email or password";
      toast.error(errorMessage);
    },
  });

  const onSubmit = (data: LoginSchemaType) => {
    mutation.mutate(data);
  };

  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 flex items-center justify-center py-12 px-4">
      <div className="w-full max-w-md">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full text-sm font-medium mb-4 bg-rose-100 text-rose-600 dark:bg-rose-900 dark:text-rose-200">
            <Flower className="size-4" />
            Welcome back to our community
          </div>
          <h1 className="text-4xl font-bold bg-linear-to-r from-rose-600 via-pink-600 to-violet-600 bg-clip-text text-transparent mb-2">
            Sign In
          </h1>
          <p className="text-gray-600 dark:text-gray-400 text-lg">
            Access your account to continue sharing beautiful moments with us
          </p>
        </div>

        {/* Login Form */}
        <Card className="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-sm shadow-xl border-0">
          <CardHeader>
            <CardTitle className="text-2xl text-center text-gray-900 dark:text-gray-100">
              Login to Your Account
            </CardTitle>
            <CardDescription className="text-center text-gray-600 dark:text-gray-400">
              Enter your credentials to access your account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <FormProvider {...form}>
              <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                {/* Email Field */}
                <div className="space-y-2">
                  <Label
                    htmlFor="email"
                    className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                  >
                    <Mail className="size-4" />
                    Email Address
                  </Label>
                  <ControlledInput
                    name="email"
                    type="email"
                    placeholder="Enter your email address"
                    disabled={isSubmitting || mutation.isPending}
                    className="h-12 text-lg"
                  />
                </div>
                {/* Password Field */}
                <div className="space-y-2">
                  <Label
                    htmlFor="password"
                    className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                  >
                    <Lock className="size-4" />
                    Password
                  </Label>
                  <ControlledInput
                    name="password"
                    type="password"
                    placeholder="Enter your password"
                    disabled={isSubmitting || mutation.isPending}
                    className="h-12 text-lg"
                  />
                </div>
                {/* Login Button */}
                <Button
                  type="submit"
                  disabled={isSubmitting || mutation.isPending}
                  className="w-full h-12 text-lg bg-linear-to-r from-rose-600 to-pink-600 text-white hover:from-rose-700 hover:to-pink-700 dark:from-rose-400 dark:to-pink-400 dark:hover:from-rose-500 dark:hover:to-pink-500"
                >
                  {isSubmitting || mutation.isPending ? (
                    <>
                      <Loader2 className="size-4 mr-2 animate-spin" />
                      <span>Signing in...</span>
                    </>
                  ) : (
                    <>
                      <LogIn className="size-4 mr-2" />
                      <span>Sign In</span>
                    </>
                  )}
                </Button>

                {/* Divider */}
                <div className="relative">
                  <div className="absolute inset-0 flex items-center">
                    <span className="w-full border-t border-gray-300 dark:border-gray-600" />
                  </div>
                  <div className="relative flex justify-center text-sm">
                    <span className="px-2 bg-white dark:bg-zinc-800 text-gray-500 dark:text-gray-400">
                      Or continue with
                    </span>
                  </div>
                </div>

                {/* OAuth Buttons */}
                <div className="grid grid-cols-2 gap-4">
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => {
                      window.location.href = `${
                        process.env.NEXT_PUBLIC_API_URL ||
                        "http://localhost:8080/api/v1"
                      }/auth/google`;
                    }}
                    className="h-12"
                  >
                    <svg className="size-5 mr-2" viewBox="0 0 24 24">
                      <path
                        fill="currentColor"
                        d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                      />
                      <path
                        fill="currentColor"
                        d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                      />
                      <path
                        fill="currentColor"
                        d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                      />
                      <path
                        fill="currentColor"
                        d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                      />
                    </svg>
                    Google
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => {
                      window.location.href = `${
                        process.env.NEXT_PUBLIC_API_URL ||
                        "http://localhost:8080/api/v1"
                      }/auth/github`;
                    }}
                    className="h-12"
                  >
                    <svg
                      className="size-5 mr-2"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
                    </svg>
                    GitHub
                  </Button>
                </div>

                {/* Register Link */}
                <div className="text-center pt-4">
                  <p className="text-gray-600">
                    Don&apos;t have an account?
                    <Link
                      href="/register"
                      className="text-rose-600 hover:text-rose-700 dark:text-rose-400 dark:hover:text-rose-300 font-medium transition-colors duration-200"
                    >
                      &nbsp;Create Account
                    </Link>
                  </p>
                </div>
              </form>
            </FormProvider>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
