"use client";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { DeletePostDialog } from "@/components/DeletePostDialog";
import {
  useDeletePost,
  useDislikePost,
  useGetPostById,
  useGetUserLikedPostIds,
  useLikePost,
} from "@/hooks/api/flowers";
import {
  useFollowUser,
  useGetUserFollowers,
  useUnfollowUser,
} from "@/hooks/api/user";
import { fallBackCopyToClipboard, getUserInitials } from "@/lib/utils";
import { useAuthStore } from "@/store/auth";
import {
  ArrowLeft,
  Calendar,
  Heart,
  Loader2,
  Pencil,
  Share2,
  Trash2,
  UserMinus,
  UserPlus,
} from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { use } from "react";
import { toast } from "sonner";

export default function FlowersDetail({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const router = useRouter();
  const { data: postData, isLoading, error } = useGetPostById(id);
  const post = postData?.post;

  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const user = useAuthStore((state) => state.user);
  const isAuthor = user?.id === post?.author.id;

  const deletePost = useDeletePost();
  const likePost = useLikePost();
  const dislikePost = useDislikePost();
  const { data: likedPostIds } = useGetUserLikedPostIds(user?.id);
  const isLiked = likedPostIds?.has(post?.id ?? 0) ?? false;

  const followUser = useFollowUser(
    user?.id?.toString() ?? "",
    post?.author.id?.toString() ?? ""
  );
  const { data: followers } = useGetUserFollowers(
    post?.author.id?.toString() ?? ""
  );
  const isFollowing =
    Array.isArray(followers) && followers.some((f) => f.id === user?.id);

  const unfollowUser = useUnfollowUser(
    user?.id?.toString() ?? "",
    post?.author.id?.toString() ?? ""
  );

  const handleDeletePost = () => {
    deletePost.mutate(id, {
      onSuccess: () => {
        router.push("/");
      },
    });
  };

  const handleLikePost = () => {
    if (!isAuthenticated) {
      toast.error("Please login to like posts");
      return;
    }
    likePost.mutate(id, {
      onError: (error) => {
        if (
          error.response?.status === 400 &&
          error.response?.data?.error === "post already liked"
        ) {
          dislikePost.mutate(id);
        }
      },
    });
  };

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

  const handleShare = async () => {
    const shareUrl = window.location.href;

    try {
      await navigator.clipboard.writeText(shareUrl);
      toast.success(fallBackCopyToClipboard(shareUrl));
    } catch (err) {
      if ((err as Error).name !== "AbortError") return;
      toast.success(fallBackCopyToClipboard(shareUrl));
    }
  };

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <Loader2 className="h-10 w-10 animate-spin text-rose-500" />
          <p className="text-muted-foreground">Loading flower...</p>
        </div>
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <p className="text-muted-foreground">
            {error ? "Error loading flower" : "Flower not found"}
          </p>
          <Button asChild variant="outline" className="rounded-none">
            <Link href="/">
              <ArrowLeft className="h-4 w-4 mr-2" />
              Go Back Home
            </Link>
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-linear-to-br from-slate-50 via-white to-rose-50 dark:from-zinc-950 dark:via-zinc-900 dark:to-zinc-950">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Back Button */}
        <Button
          variant="ghost"
          className="mb-6 rounded-none text-muted-foreground hover:text-foreground hover:bg-white/50 dark:hover:bg-zinc-800/50"
          onClick={() => router.back()}
        >
          <ArrowLeft className="h-4 w-4 mr-2" />
          Back
        </Button>

        {/* Main Content Card - Glassmorphism */}
        <Card className="md:rounded-none border border-white/20 dark:border-white/10 shadow-2xl bg-white/60 dark:bg-zinc-900/60 backdrop-blur-2xl backdrop-saturate-150 overflow-hidden">
          {/* Hero Image */}
          {post.image_url && (
            <div className="relative aspect-video sm:aspect-16/10 w-full overflow-hidden border-b border-white/20 dark:border-white/10">
              <Image
                src={post.image_url}
                alt={post.title}
                fill
                priority
                className="object-cover"
                sizes="(max-width: 768px) 100vw, 896px"
              />
              <div className="absolute inset-0 bg-linear-to-t from-black/30 via-transparent to-transparent" />
            </div>
          )}

          <CardContent className="p-6 sm:p-8">
            {/* Title */}
            <h1 className="text-2xl sm:text-3xl lg:text-4xl font-bold text-gray-900 dark:text-white mb-4">
              {post.title}
            </h1>

            {/* Author Info & Actions */}
            <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 pb-6 border-b border-white/20 dark:border-white/10">
              {/* Author */}
              <div className="flex items-center gap-3">
                <Link href={`/profile/${post.author.id}`}>
                  <Avatar className="h-12 w-12 ring-2 ring-white/30 dark:ring-white/10">
                    <AvatarImage
                      src={post.author.avatar ?? undefined}
                      alt={post.author.username}
                      className="object-cover "
                    />
                    <AvatarFallback className=" bg-linear-to-br from-rose-400 to-violet-500 text-white font-semibold">
                      {getUserInitials(post.author.username)}
                    </AvatarFallback>
                  </Avatar>
                </Link>
                <div>
                  <Link
                    href={`/profile/${post.author.id}`}
                    className="font-semibold text-gray-900 dark:text-white hover:text-rose-600 dark:hover:text-rose-400 transition-colors"
                  >
                    {isAuthor ? "You" : post.author.username}
                  </Link>
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Calendar className="h-3.5 w-3.5" />
                    <span>
                      {new Date(post.created_at).toLocaleDateString("en-US", {
                        month: "long",
                        day: "numeric",
                        year: "numeric",
                      })}
                    </span>
                  </div>
                </div>

                {/* Follow Button */}
                {isAuthenticated && !isAuthor && (
                  <Button
                    onClick={handleFollow}
                    disabled={followUser.isPending || unfollowUser.isPending}
                    variant={isFollowing ? "outline" : "default"}
                    size="sm"
                    className={
                      isFollowing
                        ? "rounded-none border-rose-200/50 dark:border-rose-800/50 bg-white/30 dark:bg-zinc-800/30 backdrop-blur-sm hover:bg-rose-50/50 dark:hover:bg-rose-900/30 hover:text-rose-600 ml-2"
                        : "rounded-none bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white ml-2"
                    }
                  >
                    {followUser.isPending || unfollowUser.isPending ? (
                      <Loader2 className="h-4 w-4 animate-spin" />
                    ) : isFollowing ? (
                      <>
                        <UserMinus className="h-4 w-4 mr-1" />
                        Unfollow
                      </>
                    ) : (
                      <>
                        <UserPlus className="h-4 w-4 mr-1" />
                        Follow
                      </>
                    )}
                  </Button>
                )}
              </div>

              {/* Action Buttons */}
              <div className="flex items-center gap-2">
                {/* Share Button */}
                <Button
                  onClick={handleShare}
                  variant="outline"
                  className="rounded-none border-violet-200/50 dark:border-violet-800/50 bg-white/30 dark:bg-zinc-800/30 backdrop-blur-sm text-violet-600 dark:text-violet-400 hover:bg-violet-50/50 dark:hover:bg-violet-900/30 hover:text-violet-700"
                >
                  <Share2 className="h-4 w-4" />
                </Button>

                {/* Like Button */}
                <Button
                  onClick={handleLikePost}
                  disabled={!isAuthenticated || likePost.isPending}
                  variant="outline"
                  className={`rounded-none border-rose-200/50 dark:border-rose-800/50 bg-white/30 dark:bg-zinc-800/30 backdrop-blur-sm ${
                    isLiked
                      ? "text-rose-600 dark:text-rose-400 bg-rose-50/50 dark:bg-rose-900/30"
                      : "hover:bg-rose-50/50 dark:hover:bg-rose-900/30 hover:text-rose-600"
                  }`}
                >
                  {likePost.isPending || dislikePost.isPending ? (
                    <Loader2 className="h-4 w-4 animate-spin mr-2" />
                  ) : (
                    <Heart
                      className={`h-4 w-4 mr-2 ${
                        isLiked ? "fill-rose-500 text-rose-500" : ""
                      }`}
                    />
                  )}
                  {post.likes_count} {post.likes_count === 1 ? "Like" : "Likes"}
                </Button>

                {/* Edit/Delete for Author */}
                {isAuthenticated && isAuthor && (
                  <>
                    <Button
                      asChild
                      variant="outline"
                      className="rounded-none border-blue-200/50 dark:border-blue-800/50 bg-white/30 dark:bg-zinc-800/30 backdrop-blur-sm text-blue-600 hover:bg-blue-50/50 hover:border-blue-300/50 dark:text-blue-400 dark:hover:bg-blue-900/30"
                    >
                      <Link href={`/flowers/${post.id}/edit`}>
                        <Pencil className="h-4 w-4 mr-2" />
                        Edit
                      </Link>
                    </Button>

                    <DeletePostDialog
                      postTitle={post.title}
                      onDelete={handleDeletePost}
                      trigger={
                        <Button
                          variant="outline"
                          disabled={deletePost.isPending}
                          className="rounded-none border-red-200/50 dark:border-red-800/50 bg-white/30 dark:bg-zinc-800/30 backdrop-blur-sm text-red-600 hover:bg-red-50/50 hover:border-red-300/50 dark:text-red-400 dark:hover:bg-red-900/30"
                        >
                          {deletePost.isPending ? (
                            <Loader2 className="h-4 w-4 animate-spin mr-2" />
                          ) : (
                            <Trash2 className="h-4 w-4 mr-2" />
                          )}
                          Delete
                        </Button>
                      }
                    />
                  </>
                )}
              </div>
            </div>

            {/* Content */}
            <div className="pt-6">
              <p className="text-gray-700 dark:text-gray-300 text-lg leading-relaxed whitespace-pre-wrap">
                {post.content}
              </p>
            </div>

            {/* Updated At */}
            {post.updated_at !== post.created_at && (
              <div className="mt-6 pt-4 border-t border-white/20 dark:border-white/10">
                <p className="text-sm text-muted-foreground">
                  Last updated:{" "}
                  {new Date(post.updated_at).toLocaleDateString("en-US", {
                    month: "long",
                    day: "numeric",
                    year: "numeric",
                  })}
                </p>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
