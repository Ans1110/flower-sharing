export default function AdminPosts() {
  return (
    <div>
      <div className="mb-6 flex items-center justify-between">
        <h1 className="text-3xl font-bold">Manage Posts</h1>
        <button className="rounded-md bg-foreground px-4 py-2 text-background hover:bg-[#383838] dark:hover:bg-[#ccc]">
          Create New Post
        </button>
      </div>
      <div className="rounded-lg border border-zinc-300 bg-white dark:border-zinc-700 dark:bg-zinc-900">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="border-b border-zinc-300 dark:border-zinc-700">
              <tr>
                <th className="px-6 py-3 text-left text-sm font-medium">ID</th>
                <th className="px-6 py-3 text-left text-sm font-medium">
                  Title
                </th>
                <th className="px-6 py-3 text-left text-sm font-medium">
                  Author
                </th>
                <th className="px-6 py-3 text-left text-sm font-medium">
                  Status
                </th>
                <th className="px-6 py-3 text-left text-sm font-medium">
                  Created
                </th>
                <th className="px-6 py-3 text-left text-sm font-medium">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody>
              <tr className="border-b border-zinc-200 dark:border-zinc-800">
                <td className="px-6 py-4 text-sm" colSpan={6}>
                  No posts found
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
