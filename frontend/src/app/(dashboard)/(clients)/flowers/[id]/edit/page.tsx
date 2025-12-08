"use client";

import PostEditModal from "../../@modal/(.)flowers/[id]/edit/page";

export default function FlowerEditPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return <PostEditModal params={params} />;
}
