"use client";

import UserTable from "@/components/UserTable";
import { useDeleteUserById, useGetUserAll } from "@/hooks/api/user";
import { Users, Loader2 } from "lucide-react";

export default function AdminUsers() {
  const { data: users, isLoading, error } = useGetUserAll();
  const { mutate: deleteUser } = useDeleteUserById();

  const handleDeleteUser = (userId: number) => {
    deleteUser(userId);
  };

  if (isLoading) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="flex flex-col items-center gap-3">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
          <p className="text-sm text-muted-foreground">Loading users...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-[400px] items-center justify-center">
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-6 text-center">
          <p className="text-sm font-medium text-destructive">
            Error: {error.response?.data?.error || "Failed to load users"}
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto space-y-6 px-4 py-6 md:px-6 lg:px-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Manage Users</h1>
          <p className="mt-2 text-sm text-muted-foreground">
            View and manage all registered users
          </p>
        </div>
        <div className="flex items-center gap-2 rounded-lg bg-primary/10 px-4 py-2">
          <Users className="h-5 w-5 text-primary" />
          <span className="text-sm font-medium">
            {users?.length || 0} Total Users
          </span>
        </div>
      </div>
      <div className="rounded-lg border bg-card shadow-sm">
        <UserTable users={users ?? []} onDeleteUser={handleDeleteUser} />
      </div>
    </div>
  );
}
