"use client";

import ProfilePostCard from "@/components/ProfilePostCard";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  useDeletePost,
  useDislikePost,
  useGetPostsByUserId,
  useGetUserLikedPostIds,
  useGetUserLikedPosts,
  useLikePost,
} from "@/hooks/api/flowers";
import {
  useFollowUser,
  useGetUserById,
  useGetUserFollowers,
  useGetUserFollowing,
  useUnfollowUser,
  useUpdateUserById,
} from "@/hooks/api/user";
import { getUserInitials } from "@/lib/utils";
import { useAuthStore } from "@/store/auth";
import { FlowerType } from "@/types/flower";
import {
  Calendar,
  Camera,
  FileText,
  Heart,
  Loader2,
  Pencil,
  UserCheck,
  UserMinus,
  UserPlus,
  Users,
} from "lucide-react";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useRef } from "react";
import { toast } from "sonner";

export default function ProfilePage() {
  const params = useParams();
  const userId = params.id as string;
  const currentUser = useAuthStore((state) => state.user);
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const updateAuthUser = useAuthStore((state) => state.updateUser);
  const isOwnProfile = currentUser?.id === Number(userId);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Fetch user data
  const { data: user, isLoading: isUserLoading } = useGetUserById(userId);
  const { data: posts, isLoading: isPostsLoading } =
    useGetPostsByUserId(userId);
  const { data: likedPostsData } = useGetUserLikedPosts(userId, 1, 100);
  const { data: followers } = useGetUserFollowers(userId);
  const { data: following } = useGetUserFollowing(userId);
  const { data: likedPostIds } = useGetUserLikedPostIds(currentUser?.id);

  // Extract liked posts array
  const likedPosts = likedPostsData?.posts ?? [];

  // Check if current user follows this profile
  const isFollowing =
    Array.isArray(followers) && followers.some((f) => f.id === currentUser?.id);

  // Mutations
  const followUser = useFollowUser(currentUser?.id?.toString() ?? "", userId);
  const unfollowUser = useUnfollowUser(
    currentUser?.id?.toString() ?? "",
    userId
  );
  const deletePost = useDeletePost();
  const likePost = useLikePost();
  const dislikePost = useDislikePost();
  const updateUser = useUpdateUserById();

  const handleFollow = () => {
    if (!isAuthenticated) {
      toast.error("Please login to follow users");
      return;
    }
    if (isFollowing) {
      unfollowUser.mutate();
    } else {
      followUser.mutate();
    }
  };

  const handleAvatarClick = () => {
    if (isOwnProfile) {
      fileInputRef.current?.click();
    }
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

    const formData = new FormData();
    formData.append("image", file);

    updateUser.mutate(
      {
        userId: Number(userId),
        formData,
        selectFields: ["avatar"],
      },
      {
        onSuccess: (data) => {
          // Update auth store with new avatar
          if ("avatar" in data) {
            updateAuthUser({ avatar: data.avatar });
          }
        },
      }
    );
  };

  const handleDeletePost = (postId: string) => {
    deletePost.mutate(postId);
  };

  const handleLikePost = (postId: string) => {
    if (!isAuthenticated) {
      toast.error("Please login to like posts");
      return;
    }
    likePost.mutate(postId, {
      onError: (error) => {
        if (
          error.response?.status === 400 &&
          error.response?.data?.error === "post already liked"
        ) {
          dislikePost.mutate(postId);
        }
      },
    });
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

  if (!user) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
            User not found
          </h1>
          <p className="text-muted-foreground mt-2">
            The user you&apos;re looking for doesn&apos;t exist.
          </p>
        </div>
      </div>
    );
  }

  const stats = [
    {
      label: "Posts",
      value: Array.isArray(posts) ? posts.length : 0,
      icon: FileText,
      href: undefined,
    },
    {
      label: "Followers",
      value: Array.isArray(followers) ? followers.length : 0,
      icon: Users,
      href: `/profile/followers/${userId}`,
    },
    {
      label: "Following",
      value: Array.isArray(following) ? following.length : 0,
      icon: UserCheck,
      href: `/profile/following/${userId}`,
    },
  ];

  return (
    <div className="min-h-screen bg-linear-to-br from-slate-50 via-white to-rose-50 dark:from-zinc-950 dark:via-zinc-900 dark:to-zinc-950">
      {/* Header Banner */}
      <div className="h-32 sm:h-48 bg-linear-to-r from-rose-400 via-pink-500 to-rose-500 dark:from-rose-600 dark:via-pink-700 dark:to-rose-600" />

      {/* Profile Content */}
      <div className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Profile Header Card */}
        <div className="relative -mt-16 sm:-mt-20">
          <Card className="border-0 shadow-xl bg-white/80 dark:bg-zinc-900/80 backdrop-blur-xl">
            <CardContent className="p-6 sm:p-8">
              <div className="flex flex-col sm:flex-row gap-6">
                {/* Avatar */}
                <div className="flex justify-center sm:justify-start">
                  <div className="relative">
                    <div
                      className={`relative group ${
                        isOwnProfile ? "cursor-pointer" : ""
                      }`}
                      onClick={handleAvatarClick}
                    >
                      <Avatar
                        key={user.avatar}
                        className="h-28 w-28 sm:h-36 sm:w-36 ring-4 ring-white dark:ring-zinc-800 shadow-2xl"
                      >
                        <AvatarImage
                          src={user.avatar}
                          alt={user.username}
                          className="object-cover"
                        />
                        <AvatarFallback className="text-3xl sm:text-4xl font-bold bg-linear-to-br from-rose-400 to-violet-500 text-white">
                          {getUserInitials(user.username)}
                        </AvatarFallback>
                      </Avatar>
                      {/* Upload overlay - only show for own profile */}
                      {isOwnProfile && (
                        <div className="absolute inset-0 flex items-center justify-center bg-black/50 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none">
                          {updateUser.isPending ? (
                            <Loader2 className="h-8 w-8 text-white animate-spin" />
                          ) : (
                            <Camera className="h-8 w-8 text-white" />
                          )}
                        </div>
                      )}
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
                </div>

                {/* User Info */}
                <div className="flex-1 text-center sm:text-left">
                  <div className="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-4">
                    <div>
                      <h1 className="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-white">
                        {user.username}
                      </h1>
                      {"email" in user && (
                        <p className="text-muted-foreground mt-1">
                          {user.email}
                        </p>
                      )}
                      {"created_at" in user && (
                        <div className="flex items-center justify-center sm:justify-start gap-2 mt-3 text-sm text-muted-foreground">
                          <Calendar className="h-4 w-4" />
                          <span>
                            Joined{" "}
                            {new Date(user.created_at ?? "").toLocaleDateString(
                              "en-US",
                              {
                                month: "long",
                                year: "numeric",
                              }
                            )}
                          </span>
                        </div>
                      )}
                    </div>

                    {/* Action Buttons */}
                    <div className="flex justify-center sm:justify-end gap-3">
                      {isOwnProfile ? (
                        <Button
                          asChild
                          variant="outline"
                          className="border-rose-200 dark:border-rose-800 hover:bg-rose-50 dark:hover:bg-rose-900/30"
                        >
                          <Link href={`/profile/${userId}/edit`}>
                            <Pencil className="h-4 w-4 mr-2" />
                            Edit Profile
                          </Link>
                        </Button>
                      ) : (
                        isAuthenticated && (
                          <Button
                            onClick={handleFollow}
                            disabled={
                              followUser.isPending || unfollowUser.isPending
                            }
                            variant={isFollowing ? "outline" : "default"}
                            className={
                              isFollowing
                                ? "border-rose-200 dark:border-rose-800 hover:bg-rose-50 dark:hover:bg-rose-900/30 hover:text-rose-600"
                                : "bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                            }
                          >
                            {followUser.isPending || unfollowUser.isPending ? (
                              <Loader2 className="h-4 w-4 animate-spin" />
                            ) : isFollowing ? (
                              <>
                                <UserMinus className="h-4 w-4 mr-2" />
                                Unfollow
                              </>
                            ) : (
                              <>
                                <UserPlus className="h-4 w-4 mr-2" />
                                Follow
                              </>
                            )}
                          </Button>
                        )
                      )}
                    </div>
                  </div>

                  {/* Stats */}
                  <div className="flex flex-row items-center justify-center sm:justify-start gap-6 sm:gap-8 mt-6 pt-6 border-t border-gray-100 dark:border-zinc-800">
                    {stats.map((stat) =>
                      stat.href ? (
                        <Link
                          key={stat.label}
                          href={stat.href}
                          className="flex flex-col items-center text-center cursor-pointer hover:bg-rose-50 dark:hover:bg-rose-900/20 rounded-lg p-2 transition-colors min-w-[80px]"
                        >
                          <div className="flex items-center justify-center gap-1.5">
                            <stat.icon className="h-4 w-4 text-rose-500" />
                            <span className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-white">
                              {stat.value}
                            </span>
                          </div>
                          <p className="text-xs sm:text-sm text-muted-foreground mt-1">
                            {stat.label}
                          </p>
                        </Link>
                      ) : (
                        <div
                          key={stat.label}
                          className="flex flex-col items-center text-center min-w-[80px]"
                        >
                          <div className="flex items-center justify-center gap-1.5">
                            <stat.icon className="h-4 w-4 text-rose-500" />
                            <span className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-white">
                              {stat.value}
                            </span>
                          </div>
                          <p className="text-xs sm:text-sm text-muted-foreground mt-1">
                            {stat.label}
                          </p>
                        </div>
                      )
                    )}
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Posts Section */}
        <div className="mt-8 pb-12">
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-2">
              <FileText className="h-5 w-5 sm:h-6 sm:w-6 text-rose-500" />
              {isOwnProfile ? "Your Flowers" : `${user.username}'s Flowers`}
            </h2>
            {isOwnProfile && (
              <Button
                asChild
                className="bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
              >
                <Link href="/flowers/new">Share New Flower</Link>
              </Button>
            )}
          </div>

          {isPostsLoading ? (
            <div className="flex justify-center py-12">
              <Loader2 className="h-8 w-8 animate-spin text-rose-500" />
            </div>
          ) : Array.isArray(posts) && posts.length > 0 ? (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              {posts.map((post: FlowerType) => (
                <ProfilePostCard
                  key={post.id}
                  post={post}
                  isAuthenticated={isAuthenticated}
                  isAuthor={isOwnProfile}
                  isLiked={likedPostIds?.has(post.id) ?? false}
                  onDelete={() => handleDeletePost(post.id.toString())}
                  onLike={() => handleLikePost(post.id.toString())}
                />
              ))}
            </div>
          ) : (
            <Card className="border-dashed border-2 border-gray-200 dark:border-zinc-700 bg-transparent">
              <CardContent className="flex flex-col items-center justify-center py-16">
                <div className="h-16 w-16 rounded-full bg-rose-100 dark:bg-rose-900/30 flex items-center justify-center mb-4">
                  <FileText className="h-8 w-8 text-rose-500" />
                </div>
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white mb-2">
                  No flowers yet
                </h3>
                <p className="text-muted-foreground text-center max-w-sm">
                  {isOwnProfile
                    ? "Share your first flower with the community!"
                    : `${user.username} hasn't shared any flowers yet.`}
                </p>
                {isOwnProfile && (
                  <Button
                    asChild
                    className="mt-4 bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                  >
                    <Link href="/flowers/new">Share Your First Flower</Link>
                  </Button>
                )}
              </CardContent>
            </Card>
          )}
        </div>

        {/* Liked Posts Section - Only show if there are liked posts */}
        {likedPosts.length > 0 && (
          <div className="mt-8 pb-12">
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-2">
                <Heart className="h-5 w-5 sm:h-6 sm:w-6 text-rose-500 fill-rose-500" />
                {isOwnProfile ? "Liked Flowers" : `${user.username}'s Likes`}
              </h2>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              {likedPosts.map((post: FlowerType) => (
                <ProfilePostCard
                  key={post.id}
                  post={post}
                  isAuthenticated={isAuthenticated}
                  isAuthor={post.author?.id === currentUser?.id}
                  isLiked={true}
                  onDelete={() => handleDeletePost(post.id.toString())}
                  onLike={() => handleLikePost(post.id.toString())}
                />
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
