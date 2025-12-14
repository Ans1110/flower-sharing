"use client";

import FlowerCard from "@/components/FlowerCard";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  useDeletePost,
  useDislikePost,
  useGetPostsPagination,
  useGetUserLikedPostIds,
  useLikePost,
  useSearchPosts,
} from "@/hooks/api/flowers";
import { useGetUserFollowing, useToggleFollow } from "@/hooks/api/user";
import { useDebounce } from "@/hooks/useDeboundce";
import { cn } from "@/lib/utils";
import { useAuthStore } from "@/store/auth";
import { FlowerType } from "@/types/flower";
import {
  ChevronFirst,
  ChevronLast,
  Heart,
  Loader2,
  Plus,
  Search,
  X,
} from "lucide-react";
import Link from "next/link";
import { useMemo, useState } from "react";
import { toast } from "sonner";

export default function Flowers() {
  const [searchQuery, setSearchQuery] = useState("");
  const debouncedSearchQuery = useDebounce(searchQuery, 300);
  const { data: searchResults, isLoading: isSearching } =
    useSearchPosts(debouncedSearchQuery);
  const [currentPage, setCurrentPage] = useState(1);
  const { data: posts, isLoading } = useGetPostsPagination({
    page: currentPage,
    limit: 10,
  });
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const user = useAuthStore((state) => state.user);
  // Include debounce pending state in loading check
  const isSearchPending =
    searchQuery !== debouncedSearchQuery && searchQuery.length > 0;
  const isSearchLoading = isSearching || isSearchPending;
  const displayData = searchQuery ? searchResults : posts?.posts;
  const pagination =
    posts && "totalPages" in posts
      ? {
          totalPages: posts.totalPages,
          currentPage: currentPage,
          onPageChange: (page: number) => setCurrentPage(page),
        }
      : null;
  const deleteFlower = useDeletePost();
  const likeFlower = useLikePost();
  const dislikeFlower = useDislikePost();
  const { data: likedPostIds } = useGetUserLikedPostIds(user?.id);
  const { data: followingList } = useGetUserFollowing(
    user?.id?.toString() ?? ""
  );

  // Create a Set of following IDs for O(1) lookup
  const followingIds = useMemo(() => {
    const ids = new Set<number>();
    if (Array.isArray(followingList)) {
      followingList.forEach((u) => ids.add(u.id));
    }
    return ids;
  }, [followingList]);

  const handleDeleteFlower = (id: string) => {
    deleteFlower.mutate(id);
  };

  const handleLikeFlower = (id: string) => {
    if (!isAuthenticated) {
      toast.error("Please login to like a flower");
      return;
    }

    // Try to like the post
    // If it's already liked (400 error), automatically toggle to dislike
    likeFlower.mutate(id, {
      onError: (error) => {
        // If post is already liked, toggle to dislike
        if (
          error.response?.status === 400 &&
          error.response?.data?.error === "post already liked"
        ) {
          dislikeFlower.mutate(id, {
            onError: () => {
              toast.error("Failed to unlike flower");
            },
          });
        }
      },
    });
  };

  const toggleFollow = useToggleFollow();

  const handleFollowAuthor = (
    authorId: number,
    isCurrentlyFollowing: boolean
  ) => {
    if (!isAuthenticated || !user?.id) {
      toast.error("Please login to follow users");
      return;
    }

    toggleFollow.mutate({
      followerId: user.id.toString(),
      followingId: authorId.toString(),
      isFollowing: isCurrentlyFollowing,
    });
  };

  if (isLoading && !searchQuery) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="flex flex-col items-center gap-3">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
          <p className="text-sm text-muted-foreground">Loading flowers...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 py-12">
      <div className="max-w-7xl mx-auto px-4">
        {/* Header */}
        <div className="text-center mb-12">
          <h1 className="text-5xl font-bold bg-linear-to-r from-rose-600 via-pink-600 to-pink-400 bg-clip-text text-transparent mb-4">
            Beautiful Flowers
            <span className="hidden md:inline">, Beautiful Moments</span>
          </h1>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto">
            Discover and share the beauty of flowers with our community.
          </p>
        </div>

        {/* Search Bar */}
        <div className="max-w-2xl mx-auto mb-8">
          <div className="relative">
            <div className="absolute left-4 top-1/2 transform -translate-y-1/2 z-10 pointer-events-none flex items-center justify-center size-5">
              <Search className="text-rose-600 dark:text-rose-400 size-5" />
            </div>
            <Input
              type="text"
              placeholder="Search flowers..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="relative pl-12 pr-12 text-lg border-2 border-rose-200 focus:border-rose-400 dark:border-rose-800 dark:focus:border-rose-600 rounded-xl shadow-sm focus:ring-1 focus:ring-rose-300 dark:focus:ring-rose-600 transition-all duration-200 backdrop-blur-sm z-0 "
            />
            {searchQuery && (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setSearchQuery("")}
                className="absolute right-2 top-1/2 transform -translate-y-1/2 z-10 text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 p-1 h-8 w-8"
              >
                <X className="size-4" />
              </Button>
            )}
            {isSearchLoading && searchQuery && (
              <div className="absolute right-12 top-1/2 transform -translate-y-1/2 z-10">
                <div className="w-4 h-4 border-2 border-rose-200 dark:border-rose-800 rounded-full animate-spin"></div>
              </div>
            )}
          </div>
          {isSearchLoading && searchQuery && (
            <p className="text-center text-sm text-rose-600 dark:text-rose-400">
              Searching flowers...
            </p>
          )}
        </div>

        {/* Add Flower Button */}
        {isAuthenticated && (
          <div className="flex justify-center mb-8">
            <Button
              asChild
              size="lg"
              className="bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg"
            >
              <Link href={"/flowers/new"}>
                <Plus className="size-4" />
                Share a New Flower
              </Link>
            </Button>
          </div>
        )}

        {/* Flower Grid */}
        {isSearchLoading && searchQuery ? (
          // Search Loading State
          <div className="flex justify-center items-center py-20">
            <div className="flex flex-col items-center space-y-4">
              <div className="w-12 h-12 border-4 border-rose-200 dark:border-rose-800 border-t-rose-500 dark:border-t-rose-400 rounded-full animate-spin"></div>
              <p className="text-lg text-gray-600 dark:text-gray-400">
                Searching for &quot;{searchQuery}&quot;...
              </p>
            </div>
          </div>
        ) : displayData && displayData.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {displayData.map((flower: FlowerType) => {
              const isAuthor = flower.author?.id === user?.id;
              const isFollowingAuthor = flower.author?.id
                ? followingIds.has(flower.author.id)
                : false;
              return (
                <FlowerCard
                  key={flower.id}
                  flower={flower}
                  isAuthenticated={isAuthenticated}
                  isAuthor={isAuthor}
                  isLiked={likedPostIds?.has(flower.id) ?? false}
                  isFollowingAuthor={isFollowingAuthor}
                  onDelete={() => handleDeleteFlower(flower.id.toString())}
                  onLike={() => handleLikeFlower(flower.id.toString())}
                  onFollow={
                    flower.author?.id
                      ? () =>
                          handleFollowAuthor(
                            flower.author.id,
                            isFollowingAuthor
                          )
                      : undefined
                  }
                />
              );
            })}
          </div>
        ) : (
          // Empty State
          <div className="text-center py-20">
            <div className="size-32 bg-linear-to-br from-rose-100 to-pink-100 dark:from-rose-900 dark:to-pink-900 rounded-full flex items-center justify-center mx-auto mb-6">
              <Heart className="size-16 text-rose-400 dark:text-rose-600" />
            </div>
            <h3 className="text-2xl font-bold text-gray-900 dark:text-gray-100 mb-4">
              {searchQuery ? "No flowers found" : "No flowers yet"}
            </h3>
            <p className="max-w-md mx-auto text-gray-600 dark:text-gray-400 mb-8">
              {searchQuery
                ? `No flowers match "${searchQuery}". Try a different search term.`
                : "Be the first to share a beautiful flower with our community!"}
            </p>
            {searchQuery ? (
              <Button
                onClick={() => setSearchQuery("")}
                size="lg"
                variant="outline"
                className="border-rose-200 text-rose-600 hover:bg-rose-50 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
              >
                <X className="size-4 mr-2" />
                Clear Search
              </Button>
            ) : (
              isAuthenticated && (
                <Button
                  asChild
                  size="lg"
                  className="bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg"
                >
                  <Link href={"/flowers/new"}>
                    <Plus className="size-4" />
                    Share a New Flower
                  </Link>
                </Button>
              )
            )}
          </div>
        )}

        {/* Pagination */}
        {!searchQuery && pagination && pagination.totalPages > 1 && (
          <div className="flex justify-center items-center mt-12 space-x-2">
            <Button
              variant="outline"
              size="sm"
              disabled={currentPage === 1}
              onClick={() => setCurrentPage(1)}
              className="border-rose-200 text-rose-600 hover:bg-rose-50 disabled:opacity-50 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
            >
              <ChevronFirst className="size-4" />
            </Button>

            {/* page numbers */}
            <div className="flex space-x-1">
              {Array.from(
                { length: Math.min(5, pagination.totalPages) },
                (_, i) => {
                  const page =
                    Math.max(
                      1,
                      Math.min(currentPage - 2, pagination.totalPages - 4)
                    ) + i;

                  if (page > pagination.totalPages) return null;

                  return (
                    <Button
                      key={page}
                      variant={page === currentPage ? "default" : "outline"}
                      size="sm"
                      onClick={() => setCurrentPage(page)}
                      className={cn(
                        "border-rose-200 text-rose-600 hover:bg-rose-50 disabled:opacity-50 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300",
                        page === currentPage && "bg-rose-50 dark:bg-rose-900/50"
                      )}
                    >
                      {page}
                    </Button>
                  );
                }
              )}
            </div>

            <Button
              variant="outline"
              size="sm"
              disabled={currentPage === pagination.totalPages}
              onClick={() => setCurrentPage(pagination.totalPages)}
              className="border-rose-200 text-rose-600 hover:bg-rose-50 disabled:opacity-50 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
            >
              <ChevronLast className="size-4" />
            </Button>
          </div>
        )}
      </div>
    </div>
  );
}
