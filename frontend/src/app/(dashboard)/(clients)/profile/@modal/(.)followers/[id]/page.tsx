"use client";

import { FollowersModal } from "@/components/FollowersModal";
import { useGetUserFollowers } from "@/hooks/api/user";
import { useRouter } from "next/navigation";
import { use, useState } from "react";

export default function FollowersModalPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const router = useRouter();
  const [open, setOpen] = useState(true);
  const [isNavigating, setIsNavigating] = useState(false);
  const { id } = use(params);
  const { data: followers, isLoading } = useGetUserFollowers(id);

  //avoid double navigation, race condition when clicking the same link multiple times
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
      users={followers || []}
      title="Followers"
      description={`${followers?.length || 0} ${
        followers?.length === 1 ? "person" : "people"
      }`}
    />
  );
}
