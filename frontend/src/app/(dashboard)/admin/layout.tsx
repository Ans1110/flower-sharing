import Link from "next/link";

export default function AdminLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="min-h-screen bg-zinc-50 dark:bg-black">
      <div className="border-b border-zinc-200 bg-white dark:border-zinc-800 dark:bg-zinc-900">
        <div className="container mx-auto px-4">
          <div className="flex h-16 items-center justify-between">
            <h1 className="text-xl font-bold">Admin Dashboard</h1>
            <nav className="flex gap-4">
              <Link
                href="/admin/posts"
                className="rounded-md px-4 py-2 hover:bg-zinc-100 dark:hover:bg-zinc-800"
              >
                Posts
              </Link>
              <Link
                href="/admin/users"
                className="rounded-md px-4 py-2 hover:bg-zinc-100 dark:hover:bg-zinc-800"
              >
                Users
              </Link>
              <Link
                href="/"
                className="rounded-md px-4 py-2 hover:bg-zinc-100 dark:hover:bg-zinc-800"
              >
                Back to Site
              </Link>
            </nav>
          </div>
        </div>
      </div>
      <div className="container mx-auto px-4 py-8">{children}</div>
    </div>
  );
}
