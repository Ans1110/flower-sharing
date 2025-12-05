"use client";

import PostTable from "@/components/PostTable";
import { useDeletePost, useGetAllPosts } from "@/hooks/api/flowers";
import { FileText, Loader2 } from "lucide-react";

export default function AdminPosts() {
  const { data: posts, isLoading, error } = useGetAllPosts();
  const { mutate: deletePost } = useDeletePost();

  const handleDeletePost = (postId: number) => {
    deletePost(postId.toString());
  };

  if (isLoading) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="flex flex-col items-center gap-3">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
          <p className="text-sm text-muted-foreground">Loading posts...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-6 text-center">
          <p className="text-sm font-medium text-destructive">
            Error: {error.response?.data?.error || "Failed to load posts"}
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto space-y-6 px-4 py-6 md:px-6 lg:px-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Manage Posts</h1>
          <p className="mt-2 text-sm text-muted-foreground">
            View and manage all registered posts
          </p>
        </div>
        <div className="flex items-center gap-2 rounded-lg bg-primary/10 px-4 py-2">
          <FileText className="h-5 w-5 text-primary" />
          <span className="text-sm font-medium">
            {posts?.length || 0} Total Posts
          </span>
        </div>
      </div>
      <div className="rounded-lg border bg-card shadow-sm">
        <PostTable posts={posts ?? []} onDeletePost={handleDeletePost} />
      </div>
    </div>
  );
}
