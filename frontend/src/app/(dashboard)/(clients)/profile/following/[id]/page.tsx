"use client";

import FollowingModalPage from "../../@modal/(.)following/[id]/page";

export default function FollowingPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return <FollowingModalPage params={params} />;
}
