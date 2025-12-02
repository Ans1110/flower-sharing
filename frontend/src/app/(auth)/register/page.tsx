"use client";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ControlledInput } from "@/components/ui/controlled-input";
import { Label } from "@/components/ui/label";
import { api } from "@/service/api";
import { useAuthStore } from "@/store/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { Flower, Loader2, Lock, Mail, User, UserPlus } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { FormProvider, useForm } from "react-hook-form";
import { toast } from "sonner";
import z from "zod";

const registerSchema = z.object({
  username: z
    .string()
    .min(2, "Username must be at least 2 characters")
    .max(15, "Username must be less than 15 characters"),
  email: z.string().regex(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/, {
    message: "Please enter a valid email address",
  }),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters")
    .max(255, "Password must be less than 255 characters"),
});

type RegisterSchemaType = z.infer<typeof registerSchema>;

export default function Register() {
  const router = useRouter();
  const register = useAuthStore((state) => state.register);
  const form = useForm<RegisterSchemaType>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      username: "",
      email: "",
      password: "",
    },
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = form;

  const mutation = useMutation({
    mutationFn: (data: RegisterSchemaType) => api.post("/auth/register", data),
    onSuccess: ({ data }) => {
      register(data.user, data.accessToken);
      toast.success("Welcome! Registration successful");
      router.push("/");
    },
    onError: (error: AxiosError<{ error: string }>) => {
      const errorMessage =
        error.response?.data?.error || "An unexpected error occurred";
      toast.error(errorMessage);
    },
  });

  const onSubmit = (data: RegisterSchemaType) => {
    mutation.mutate(data);
  };

  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 flex items-center justify-center py-12 px-4">
      <div className="w-full max-w-md">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full text-sm font-medium mb-4 bg-rose-100 text-rose-600 dark:bg-rose-900 dark:text-rose-200">
            <Flower className="size-4" />
            Welcome to our community
          </div>
          <h1 className="text-4xl font-bold bg-linear-to-r from-rose-600 via-pink-600 to-red-600 bg-clip-text text-transparent mb-2">
            Create Account
          </h1>
          <p className="text-gray-600 dark:text-gray-400 text-lg">
            Start sharing your beautiful moments with our community
          </p>
        </div>

        {/* Register Form */}
        <Card className="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-sm shadow-xl border-0">
          <CardHeader>
            <CardTitle className="text-2xl text-center text-gray-900 dark:text-gray-100">
              Join Flower Sharing
            </CardTitle>
            <CardDescription className="text-center text-gray-600 dark:text-gray-400">
              Create an account to start sharing beautiful flowers
            </CardDescription>
            <CardContent>
              <FormProvider {...form}>
                <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                  {/* Username Field */}
                  <div className="space-y-2">
                    <Label
                      htmlFor="username"
                      className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                    >
                      <User className="size-4" />
                      Username
                    </Label>
                    <ControlledInput
                      name="username"
                      type="text"
                      placeholder="Enter your username"
                      disabled={isSubmitting || mutation.isPending}
                      className="h-12 text-lg"
                    />
                  </div>
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
                  {/* Register Button */}
                  <Button
                    type="submit"
                    disabled={isSubmitting || mutation.isPending}
                    className="w-full h-12 text-lg bg-linear-to-r from-rose-600 to-pink-600 text-white hover:from-rose-700 hover:to-pink-700 dark:from-rose-400 dark:to-pink-400 dark:hover:from-rose-500 dark:hover:to-pink-500"
                  >
                    {isSubmitting || mutation.isPending ? (
                      <>
                        <Loader2 className="size-4 mr-2 animate-spin" />
                        <span>Registering...</span>
                      </>
                    ) : (
                      <>
                        <UserPlus className="size-4 mr-2" />
                        <span>Create Account</span>
                      </>
                    )}
                  </Button>

                  {/* Login Link */}
                  <div className="text-center pt-4">
                    <p className="text-gray-600">
                      Already have an account?
                      <Link
                        href="/login"
                        className="text-rose-600 hover:text-rose-700 dark:text-rose-400 dark:hover:text-rose-300 font-medium transition-colors duration-200"
                      >
                        &nbsp;Sign in
                      </Link>
                    </p>
                  </div>
                </form>
              </FormProvider>
            </CardContent>
          </CardHeader>
        </Card>
      </div>
    </div>
  );
}
