import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  useDeleteFlower,
  useFlowers,
  useSearchFlowers,
} from "@/hooks/api/flowers";
import { Link } from "react-router";
import { toast } from "sonner";
import {
  Plus,
  Edit,
  Trash2,
  Calendar,
  User,
  Heart,
  Search,
  X,
} from "lucide-react";
import { useAuthStore } from "@/store/auth";
import { useState } from "react";
import type { FlowerType } from "@/types/flower";

const Flowers = () => {
  const [searchQuery, setSearchQuery] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const { data: flowersData, isLoading } = useFlowers(currentPage, 6);
  const { data: searchResults, isLoading: isSearching } =
    useSearchFlowers(searchQuery);
  const deleteFlower = useDeleteFlower();
  const token = useAuthStore((state) => state.token);
  const user = useAuthStore((state) => state.user);

  // Use search results if searching, otherwise use paginated flowers
  const displayData = searchQuery ? searchResults : flowersData?.data;
  const pagination = flowersData?.pagination;

  // Only show loading screen for initial load, not for search
  if (isLoading && !searchQuery) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 py-12">
        <div className="max-w-7xl mx-auto px-4">
          <div className="flex justify-center items-center h-64">
            <div className="flex flex-col items-center space-y-4">
              <div className="w-12 h-12 border-4 border-rose-200 border-t-rose-500 rounded-full animate-spin"></div>
              <p className="text-rose-600 font-medium">
                Loading beautiful flowers...
              </p>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 py-12">
      <div className="max-w-7xl mx-auto px-4">
        {/* Header */}
        <div className="text-center mb-12">
          <h1 className="text-5xl font-bold bg-gradient-to-r from-rose-600 via-pink-600 to-purple-600 bg-clip-text text-transparent mb-4">
            Beautiful Flowers
          </h1>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto">
            Discover and share the most beautiful flowers from our community
          </p>
        </div>

        {/* Search Bar */}
        <div className="max-w-2xl mx-auto mb-8">
          <div className="relative">
            <div className="absolute left-4 top-1/2 transform -translate-y-1/2 z-10 pointer-events-none flex items-center justify-center w-5 h-5">
              <Search className="text-rose-400 w-5 h-5" />
            </div>
            <Input
              type="text"
              placeholder="Search flowers by title..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pl-12 pr-12 h-12 text-lg border-2 border-rose-200 focus:border-rose-400 rounded-xl bg-white/80 backdrop-blur-sm relative z-0"
            />
            {searchQuery && (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setSearchQuery("")}
                className="absolute right-2 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 p-1 h-8 w-8"
              >
                <X className="w-4 h-4" />
              </Button>
            )}
            {isSearching && searchQuery && (
              <div className="absolute right-12 top-1/2 transform -translate-y-1/2">
                <div className="w-4 h-4 border-2 border-rose-200 border-t-rose-500 rounded-full animate-spin"></div>
              </div>
            )}
          </div>
          {isSearching && searchQuery && (
            <p className="text-center text-rose-600 text-sm mt-2">
              Searching flowers...
            </p>
          )}
        </div>

        {/* Add Flower Button */}
        {token && (
          <div className="flex justify-center mb-8">
            <Button
              asChild
              size="lg"
              className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg shadow-lg hover:shadow-xl transition-all duration-300"
            >
              <Link to="/flowers/new">
                <Plus className="w-5 h-5 mr-2" />
                Share a New Flower
              </Link>
            </Button>
          </div>
        )}

        {/* Flowers Grid */}
        {displayData && displayData.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {displayData.map((flower: FlowerType) => (
              <Card
                key={flower.id}
                className="group bg-white/80 backdrop-blur-sm border-0 shadow-lg hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-2 overflow-hidden"
              >
                {/* Image */}
                {flower.image_url && (
                  <div className="aspect-square overflow-hidden">
                    <img
                      src={flower.image_url}
                      alt={flower.title}
                      className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
                    />
                  </div>
                )}

                {/* Content */}
                <CardHeader className="pb-3">
                  <CardTitle className="text-xl font-bold text-gray-900 line-clamp-2 group-hover:text-rose-600 transition-colors duration-300">
                    <Link to={`/flowers/${flower.id}`}>{flower.title}</Link>
                  </CardTitle>
                </CardHeader>

                <CardContent className="pb-4">
                  <p className="text-gray-600 line-clamp-3 leading-relaxed">
                    {flower.content}
                  </p>

                  {/* Meta Info */}
                  <div className="flex items-center gap-4 mt-4 text-sm text-gray-500">
                    <div className="flex items-center gap-1">
                      <Calendar className="w-4 h-4" />
                      <span>
                        {new Date(flower.created_at).toLocaleDateString()}
                      </span>
                    </div>
                    <div className="flex items-center gap-1">
                      <User className="w-4 h-4" />
                      <span>
                        {user?.id === flower.author_id
                          ? "You"
                          : flower.author_username}
                      </span>
                    </div>
                    <div className="flex items-center gap-1">
                      <Heart className="w-4 h-4" />
                      <span>{flower.likes || 0} likes</span>
                    </div>
                  </div>
                </CardContent>

                {/* Actions */}
                {token &&
                  (user?.id === flower.author_id || user?.role === "admin") && (
                    <CardFooter className="pt-0">
                      <div className="flex gap-2 w-full">
                        <Button
                          asChild
                          size="sm"
                          variant="outline"
                          className="flex-1 border-rose-200 text-rose-600 hover:bg-rose-50"
                        >
                          <Link to={`/flowers/${flower.id}/edit`}>
                            <Edit className="w-4 h-4 mr-1" />
                            Edit
                          </Link>
                        </Button>
                        <Button
                          size="sm"
                          variant="destructive"
                          className="flex-1 border-rose-200 hover:bg-rose-500 text-white"
                          onClick={() =>
                            deleteFlower.mutate(flower.id.toString(), {
                              onSuccess: () =>
                                toast.success("Flower deleted successfully"),
                            })
                          }
                        >
                          <Trash2 className="w-4 h-4 mr-1" />
                          Delete
                        </Button>
                      </div>
                    </CardFooter>
                  )}
              </Card>
            ))}
          </div>
        ) : (
          /* Empty State */
          <div className="text-center py-20">
            <div className="w-32 h-32 bg-gradient-to-br from-rose-100 to-pink-100 rounded-full flex items-center justify-center mx-auto mb-6">
              <Heart className="w-16 h-16 text-rose-400" />
            </div>
            <h3 className="text-2xl font-bold text-gray-900 mb-4">
              {searchQuery ? "No flowers found" : "No flowers yet"}
            </h3>
            <p className="text-gray-600 mb-8 max-w-md mx-auto">
              {searchQuery
                ? `No flowers match "${searchQuery}". Try a different search term.`
                : "Be the first to share a beautiful flower with our community!"}
            </p>
            {searchQuery ? (
              <Button
                onClick={() => setSearchQuery("")}
                size="lg"
                variant="outline"
                className="border-rose-200 text-rose-600 hover:bg-rose-50"
              >
                <X className="w-5 h-5 mr-2" />
                Clear Search
              </Button>
            ) : (
              token && (
                <Button
                  asChild
                  size="lg"
                  className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                >
                  <Link to="/flowers/new">
                    <Plus className="w-5 h-5 mr-2" />
                    Share Your First Flower
                  </Link>
                </Button>
              )
            )}
          </div>
        )}

        {/* Pagination Controls */}
        {!searchQuery && pagination && pagination.total_pages > 1 && (
          <div className="flex justify-center items-center mt-12 space-x-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() => setCurrentPage(1)}
              disabled={currentPage === 1}
              className="border-rose-200 text-rose-600 hover:bg-rose-50 disabled:opacity-50"
            >
              First
            </Button>

            {/* Page Numbers */}
            <div className="flex space-x-1">
              {Array.from(
                { length: Math.min(5, pagination.total_pages) },
                (_, i) => {
                  const pageNum =
                    Math.max(
                      1,
                      Math.min(pagination.total_pages - 4, currentPage - 2)
                    ) + i;

                  if (pageNum > pagination.total_pages) return null;

                  return (
                    <Button
                      key={pageNum}
                      variant={pageNum === currentPage ? "default" : "outline"}
                      size="sm"
                      onClick={() => setCurrentPage(pageNum)}
                      className={
                        pageNum === currentPage
                          ? "bg-rose-500 text-white hover:bg-rose-600"
                          : "border-rose-200 text-rose-600 hover:bg-rose-50"
                      }
                    >
                      {pageNum}
                    </Button>
                  );
                }
              )}
            </div>

            <Button
              variant="outline"
              size="sm"
              onClick={() => setCurrentPage(pagination.total_pages)}
              disabled={currentPage === pagination.total_pages}
              className="border-rose-200 text-rose-600 hover:bg-rose-50 disabled:opacity-50"
            >
              Last
            </Button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Flowers;
