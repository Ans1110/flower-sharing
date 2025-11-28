"use client";

import { useRouter } from "next/navigation";
import { useGetUserById, useUpdateUserById } from "@/hooks/api/user";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import EditUserForm from "@/components/EditUserForm";
import { UserPayloadType, UserAdminResponseType } from "@/types/user";
import { Loader2 } from "lucide-react";
import { use, useEffect, useState } from "react";

export default function EditUserModal({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const router = useRouter();
  const [open, setOpen] = useState(true);
  const { id } = use(params);
  const { data: user, isLoading, error } = useGetUserById(id);
  const { mutate: updateUser, isPending } = useUpdateUserById();

  const handleClose = () => {
    setOpen(false);
  };

  const handleSubmit = (payload: UserPayloadType) => {
    updateUser(
      { userId: Number(id), payload },
      {
        onSuccess: () => {
          handleClose();
        },
      }
    );
  };

  // Handle dialog close (ESC key or backdrop click)
  useEffect(() => {
    if (!open) {
      // Small delay to allow the dialog close animation to complete
      const timer = setTimeout(() => {
        router.back();
      }, 150);
      return () => clearTimeout(timer);
    }
  }, [open, router]);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-[500px]">
        <DialogHeader>
          <DialogTitle>Edit User</DialogTitle>
          <DialogDescription>
            Update user information. Click save when you&apos;re done.
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
              {error.response?.data?.error || "Failed to load user"}
            </p>
          </div>
        )}

        {user && (
          <EditUserForm
            user={user as UserAdminResponseType}
            onSubmit={handleSubmit}
            onCancel={handleClose}
            isLoading={isPending}
          />
        )}
      </DialogContent>
    </Dialog>
  );
}
