"use client";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ControlledInput } from "@/components/ui/controlled-input";
import { useGetUserById, useUpdateUserById } from "@/hooks/api/user";
import { getUserInitials } from "@/lib/utils";
import { useAuthStore } from "@/store/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { Camera, Loader2, Save, X } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

// Form validation schema
const profileFormSchema = z.object({
  username: z
    .string()
    .min(3, "Username must be at least 3 characters")
    .max(50, "Username must not exceed 50 characters")
    .trim(),
  email: z
    .string()
    .email("Please enter a valid email address")
    .trim()
    .optional()
    .or(z.literal("")),
});

type ProfileFormValues = z.infer<typeof profileFormSchema>;

export default function ProfileEditPage() {
  const params = useParams();
  const router = useRouter();
  const userId = params.id as string;
  const currentUser = useAuthStore((state) => state.user);
  const updateAuthUser = useAuthStore((state) => state.updateUser);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const isOwnProfile = currentUser?.id === Number(userId);

  const {
    data: user,
    isLoading: isUserLoading,
    refetch,
  } = useGetUserById(userId);
  const { mutate: updateUser, isPending } = useUpdateUserById();

  const [avatarPreview, setAvatarPreview] = useState<string | null>(null);
  const [avatarFile, setAvatarFile] = useState<File | null>(null);

  const methods = useForm<ProfileFormValues>({
    resolver: zodResolver(profileFormSchema),
    defaultValues: {
      username: "",
      email: "",
    },
  });

  const {
    handleSubmit,
    formState: { isDirty },
    reset,
    watch,
  } = methods;

  // Watch form values for avatar preview
  const username = watch("username");

  // Initialize form with user data
  useEffect(() => {
    if (currentUser && isOwnProfile) {
      reset({
        username: currentUser.username,
        email: currentUser.email || "",
      });
      setAvatarPreview(currentUser.avatar);
    }
  }, [currentUser, isOwnProfile, reset]);

  // Redirect if not own profile
  useEffect(() => {
    if (!isUserLoading && !isOwnProfile) {
      toast.error("You can only edit your own profile");
      router.push(`/profile/${userId}`);
    }
  }, [isOwnProfile, isUserLoading, userId, router]);

  const handleAvatarClick = () => {
    fileInputRef.current?.click();
  };

  const handleAvatarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    // Validate file type
    if (!file.type.startsWith("image/")) {
      toast.error("Please select an image file");
      return;
    }

    // Validate file size (max 5MB)
    if (file.size > 5 * 1024 * 1024) {
      toast.error("Image must be less than 5MB");
      return;
    }

    setAvatarFile(file);
    const reader = new FileReader();
    reader.onloadend = () => {
      setAvatarPreview(reader.result as string);
    };
    reader.readAsDataURL(file);
  };

  const handleRemoveAvatar = () => {
    setAvatarFile(null);
    setAvatarPreview(currentUser?.avatar || null);
  };

  const onSubmit = (data: ProfileFormValues) => {
    const formData = new FormData();
    formData.append("username", data.username);
    formData.append("email", data.email || "");

    if (avatarFile) {
      formData.append("image", avatarFile);
    }

    updateUser(
      {
        userId: Number(userId),
        formData,
        selectFields: ["id", "username", "email", "avatar"],
      },
      {
        onSuccess: async (responseData) => {
          updateAuthUser({
            username: responseData.username,
            email:
              "email" in responseData ? responseData.email : currentUser?.email,
            avatar: responseData.avatar,
          });

          await refetch();
          router.push(`/profile/${userId}`);
        },
      }
    );
  };

  const handleCancel = () => {
    router.push(`/profile/${userId}`);
  };

  if (isUserLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <Loader2 className="h-10 w-10 animate-spin text-rose-500" />
          <p className="text-muted-foreground">Loading profile...</p>
        </div>
      </div>
    );
  }

  if (!user || !isOwnProfile) {
    return null;
  }

  return (
    <div className="min-h-screen bg-linear-to-br from-slate-50 via-white to-rose-50 dark:from-zinc-950 dark:via-zinc-900 dark:to-zinc-950 py-8">
      <div className="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8">
        <Card className="border-0 shadow-xl bg-white/80 dark:bg-zinc-900/80 backdrop-blur-xl">
          <CardHeader className="border-b border-gray-100 dark:border-zinc-800">
            <CardTitle className="text-2xl font-bold text-gray-900 dark:text-white">
              Edit Profile
            </CardTitle>
            <p className="text-sm text-muted-foreground mt-1">
              Update your profile information and avatar
            </p>
          </CardHeader>
          <CardContent className="p-6">
            <FormProvider {...methods}>
              <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                {/* Avatar Section */}
                <div className="flex flex-col items-center gap-4">
                  <div className="relative">
                    <div
                      className="relative group cursor-pointer"
                      onClick={handleAvatarClick}
                    >
                      <Avatar
                        key={avatarPreview}
                        className="h-32 w-32 ring-4 ring-white dark:ring-zinc-800 shadow-2xl"
                      >
                        <AvatarImage
                          src={avatarPreview || undefined}
                          alt={username || "User"}
                          className="object-cover"
                        />
                        <AvatarFallback className="text-4xl font-bold bg-linear-to-br from-rose-400 to-violet-500 text-white">
                          {getUserInitials(username || "User")}
                        </AvatarFallback>
                      </Avatar>
                      {/* Upload overlay */}
                      <div className="absolute inset-0 flex items-center justify-center bg-black/50 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none">
                        <Camera className="h-8 w-8 text-white" />
                      </div>
                    </div>
                    {/* Hidden file input */}
                    <input
                      ref={fileInputRef}
                      type="file"
                      accept="image/*"
                      className="hidden"
                      onChange={handleAvatarChange}
                    />
                  </div>
                  <div className="flex gap-2">
                    <Button
                      type="button"
                      variant="outline"
                      size="sm"
                      onClick={handleAvatarClick}
                      className="border-rose-200 dark:border-rose-800 hover:bg-rose-50 dark:hover:bg-rose-900/30"
                    >
                      <Camera className="h-4 w-4 mr-2" />
                      Change Avatar
                    </Button>
                    {avatarFile && (
                      <Button
                        type="button"
                        variant="outline"
                        size="sm"
                        onClick={handleRemoveAvatar}
                        className="border-gray-200 dark:border-zinc-700"
                      >
                        <X className="h-4 w-4 mr-2" />
                        Remove
                      </Button>
                    )}
                  </div>
                  <p className="text-xs text-muted-foreground text-center">
                    Recommended: Square image, at least 400x400px
                    <br />
                    Max file size: 5MB
                  </p>
                </div>

                {/* Username Field */}
                <div className="space-y-2">
                  <ControlledInput<ProfileFormValues>
                    name="username"
                    type="text"
                    label={
                      <>
                        Username <span className="text-rose-500">*</span>
                      </>
                    }
                    placeholder="Enter your username"
                    className="w-full"
                    containerClassName="space-y-0"
                  />
                  <p className="text-xs text-muted-foreground">
                    Minimum 3 characters, maximum 50 characters
                  </p>
                </div>

                {/* Email Field */}
                {currentUser?.email && (
                  <div className="space-y-2">
                    <ControlledInput<ProfileFormValues>
                      name="email"
                      type="email"
                      label={
                        <>
                          Email <span className="text-rose-500">*</span>
                        </>
                      }
                      placeholder="Enter your email"
                      className="w-full"
                      containerClassName="space-y-0"
                    />
                    <p className="text-xs text-muted-foreground">
                      Your email address for account notifications
                    </p>
                  </div>
                )}

                {/* Provider Info */}
                {currentUser?.provider && (
                  <div className="rounded-lg bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 p-4">
                    <p className="text-sm text-blue-800 dark:text-blue-200">
                      <strong>Account Type:</strong>{" "}
                      {currentUser.provider.charAt(0).toUpperCase() +
                        currentUser.provider.slice(1)}
                    </p>
                    <p className="text-xs text-blue-600 dark:text-blue-300 mt-1">
                      This account is linked with {currentUser.provider}
                    </p>
                  </div>
                )}

                {/* Action Buttons */}
                <div className="flex gap-3 pt-4 border-t border-gray-100 dark:border-zinc-800">
                  <Button
                    type="submit"
                    disabled={(!isDirty && !avatarFile) || isPending}
                    className="flex-1 bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    {isPending ? (
                      <>
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                        Saving...
                      </>
                    ) : (
                      <>
                        <Save className="h-4 w-4 mr-2" />
                        Save Changes
                      </>
                    )}
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={handleCancel}
                    disabled={isPending}
                    className="flex-1 border-gray-200 dark:border-zinc-700"
                  >
                    <X className="h-4 w-4 mr-2" />
                    Cancel
                  </Button>
                </div>
              </form>
            </FormProvider>
          </CardContent>
        </Card>

        {/* Additional Info Card */}
        <Card className="mt-6 border-0 shadow-lg bg-white/60 dark:bg-zinc-900/60 backdrop-blur-xl">
          <CardContent className="p-6">
            <h3 className="font-semibold text-gray-900 dark:text-white mb-2">
              Profile Tips
            </h3>
            <ul className="space-y-2 text-sm text-muted-foreground">
              <li className="flex items-start gap-2">
                <span className="text-rose-500 mt-0.5">•</span>
                <span>Choose a unique username that represents you</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-rose-500 mt-0.5">•</span>
                <span>Use a clear profile picture for better recognition</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-rose-500 mt-0.5">•</span>
                <span>
                  Keep your email up to date for important notifications
                </span>
              </li>
            </ul>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
