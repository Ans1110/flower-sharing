export default function Profile() {
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Profile</h1>
      <div className="max-w-2xl space-y-6">
        <div className="rounded-lg border border-zinc-300 p-6 dark:border-zinc-700">
          <h2 className="text-xl font-semibold mb-4">User Information</h2>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-1">Username</label>
              <p className="text-zinc-600 dark:text-zinc-400">username</p>
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Email</label>
              <p className="text-zinc-600 dark:text-zinc-400">
                user@example.com
              </p>
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Role</label>
              <p className="text-zinc-600 dark:text-zinc-400">User</p>
            </div>
          </div>
        </div>
        <div className="flex gap-4">
          <button className="rounded-md bg-foreground px-6 py-2 text-background hover:bg-[#383838] dark:hover:bg-[#ccc]">
            Edit Profile
          </button>
          <button className="rounded-md border border-zinc-300 px-6 py-2 hover:bg-zinc-100 dark:border-zinc-700 dark:hover:bg-zinc-800">
            Change Password
          </button>
        </div>
      </div>
    </div>
  );
}
