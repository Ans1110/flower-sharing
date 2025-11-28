"use client";

import { useRouter } from "next/navigation";
import { useGetUserById, useUpdateUserById } from "@/hooks/api/user";
import EditUserForm from "@/components/EditUserForm";
import { UserPayloadType, UserAdminResponseType } from "@/types/user";
import { Loader2, ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { use } from "react";

export default function EditUserPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const router = useRouter();
  const { id } = use(params);
  const { data: user, isLoading, error } = useGetUserById(id);
  const { mutate: updateUser, isPending } = useUpdateUserById();

  const handleSubmit = (payload: UserPayloadType) => {
    updateUser(
      { userId: Number(id), payload },
      {
        onSuccess: () => {
          router.push("/admin/users");
        },
      }
    );
  };

  const handleCancel = () => {
    router.push("/admin/users");
  };

  if (isLoading) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="flex flex-col items-center gap-3">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
          <p className="text-sm text-muted-foreground">Loading user...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container mx-auto max-w-2xl px-4 py-6">
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-6 text-center">
          <p className="text-sm font-medium text-destructive">
            {error.response?.data?.error || "Failed to load user"}
          </p>
          <Button
            variant="outline"
            className="mt-4"
            onClick={() => router.push("/admin/users")}
          >
            <ArrowLeft className="mr-2 h-4 w-4" />
            Back to Users
          </Button>
        </div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <div className="container mx-auto max-w-2xl space-y-6 px-4 py-6">
      <div className="flex items-center gap-4">
        <Link href="/admin/users">
          <Button variant="ghost" size="icon">
            <ArrowLeft className="h-5 w-5" />
          </Button>
        </Link>
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Edit User</h1>
          <p className="mt-1 text-sm text-muted-foreground">
            Update user information for {user.username}
          </p>
        </div>
      </div>

      <div className="rounded-lg border bg-card p-6 shadow-sm">
        <EditUserForm
          user={user as UserAdminResponseType}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          isLoading={isPending}
        />
      </div>
    </div>
  );
}
