"use client";

import EditPostModal from "../../../@modal/(.)posts/[id]/edit/page";

export default function EditPostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return <EditPostModal params={params} />;
}
