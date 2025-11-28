export default function FlowerForm() {
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Create New Flower</h1>
      <form className="max-w-2xl space-y-6">
        <div>
          <label htmlFor="name" className="block text-sm font-medium mb-2">
            Flower Name
          </label>
          <input
            id="name"
            name="name"
            type="text"
            required
            className="w-full rounded-md border border-zinc-300 px-3 py-2 dark:border-zinc-700 dark:bg-zinc-800"
          />
        </div>
        <div>
          <label
            htmlFor="description"
            className="block text-sm font-medium mb-2"
          >
            Description
          </label>
          <textarea
            id="description"
            name="description"
            rows={4}
            required
            className="w-full rounded-md border border-zinc-300 px-3 py-2 dark:border-zinc-700 dark:bg-zinc-800"
          />
        </div>
        <div>
          <label htmlFor="image" className="block text-sm font-medium mb-2">
            Image
          </label>
          <input
            id="image"
            name="image"
            type="file"
            accept="image/*"
            className="w-full"
          />
        </div>
        <div className="flex gap-4">
          <button
            type="submit"
            className="rounded-md bg-foreground px-6 py-2 text-background hover:bg-[#383838] dark:hover:bg-[#ccc]"
          >
            Create Flower
          </button>
          <button
            type="button"
            className="rounded-md border border-zinc-300 px-6 py-2 hover:bg-zinc-100 dark:border-zinc-700 dark:hover:bg-zinc-800"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}
