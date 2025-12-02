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
