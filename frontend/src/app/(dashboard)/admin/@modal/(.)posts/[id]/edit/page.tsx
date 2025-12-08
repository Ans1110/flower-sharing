"use client";

import EditPostForm from "@/components/EditPostForm";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useGetPostById, useUpdatePost } from "@/hooks/api/flowers";
import { FlowerPayloadType } from "@/types/flower";
import { Loader2 } from "lucide-react";
import { useRouter } from "next/navigation";
import { use, useEffect, useState } from "react";

export default function EditPostModal({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const router = useRouter();
  const [open, setOpen] = useState(true);
  const { id } = use(params);
  const { data: post, isLoading, error, refetch } = useGetPostById(id);
  const { mutate: updatePost, isPending } = useUpdatePost();

  const handleClose = () => {
    setOpen(false);
  };

  const handleSubmit = (payload: FlowerPayloadType) => {
    const formData = new FormData();
    formData.append("title", payload.title);
    formData.append("content", payload.content);
    if (payload.imageUrl instanceof File) {
      formData.append("image", payload.imageUrl);
    }

    updatePost(
      {
        postId: id,
        formData,
        selectFields: ["id", "title", "content", "image_url"],
      },
      {
        onSuccess: async () => {
          await refetch();
          handleClose();
        },
      }
    );
  };

  useEffect(() => {
    if (!open) {
      const timer = setTimeout(() => {
        router.back();
      }, 150);
      return () => clearTimeout(timer);
    }
  }, [open, router]);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Edit Post</DialogTitle>
          <DialogDescription>
            Update your post. Click save when you&apos;re done.
          </DialogDescription>
        </DialogHeader>

        {isLoading && (
          <div className="flex items-center justify-center py-8">
            <Loader2 className="h-8 w-8 animate-spin text-primary" />
          </div>
        )}

        {error && (
          <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
            <p className="text-sm text-destructive">
              {error.response?.data?.error || "Failed to load post"}
            </p>
          </div>
        )}

        {post && (
          <EditPostForm
            post={post.post}
            onSubmit={handleSubmit}
            onCancel={handleClose}
            isLoading={isPending}
          />
        )}
      </DialogContent>
    </Dialog>
  );
}
