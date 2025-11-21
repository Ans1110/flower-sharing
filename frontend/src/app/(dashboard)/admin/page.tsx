import Link from "next/link";

export default function AdminDashboard() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Admin Dashboard</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Link
          href="/admin/posts"
          className="rounded-lg border border-zinc-300 p-6 hover:border-zinc-400 dark:border-zinc-700 dark:hover:border-zinc-600"
        >
          <h2 className="text-xl font-semibold mb-2">Manage Posts</h2>
          <p className="text-zinc-600 dark:text-zinc-400">
            View, edit, and delete posts
          </p>
        </Link>
        <Link
          href="/admin/users"
          className="rounded-lg border border-zinc-300 p-6 hover:border-zinc-400 dark:border-zinc-700 dark:hover:border-zinc-600"
        >
          <h2 className="text-xl font-semibold mb-2">Manage Users</h2>
          <p className="text-zinc-600 dark:text-zinc-400">
            View, edit, and manage users
          </p>
        </Link>
      </div>
    </div>
  );
}
