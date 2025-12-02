"use client";

import EditUserModal from "../../../@modal/(.)users/[id]/edit/page";

export default function EditUserPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return <EditUserModal params={params} />;
}
