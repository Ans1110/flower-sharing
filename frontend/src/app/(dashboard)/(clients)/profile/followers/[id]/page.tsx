"use client";

import FollowersModalPage from "../../@modal/(.)followers/[id]/page";

export default function FollowersPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return <FollowersModalPage params={params} />;
}
