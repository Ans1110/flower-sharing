"use client";

import { FollowersModal } from "@/components/FollowersModal";
import { useGetUserFollowing } from "@/hooks/api/user";
import { useRouter } from "next/navigation";
import { use, useState } from "react";

export default function FollowingModalPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const router = useRouter();
  const [open, setOpen] = useState(true);
  const [isNavigating, setIsNavigating] = useState(false);
  const { id } = use(params);
  const { data: following, isLoading } = useGetUserFollowing(id);

  const handleClose = () => {
    if (isNavigating) return;
    setOpen(false);
    setTimeout(() => {
      router.back();
    }, 150);
  };

  const handleNavigate = (path: string) => {
    setIsNavigating(true);
    setOpen(false);
    // Use hard navigation to ensure page fully updates when navigating from modal
    window.location.href = path;
  };

  if (isLoading) {
    return null;
  }

  return (
    <FollowersModal
      isOpen={open}
      onClose={handleClose}
      onNavigate={handleNavigate}
      users={following || []}
      title="Following"
      description={`${following?.length || 0} ${
        following?.length === 1 ? "person" : "people"
      }`}
    />
  );
}
