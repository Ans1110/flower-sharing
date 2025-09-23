import { Button } from "@/components/ui/button";
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
  type FlowerType,
} from "@/hooks/api/flowers";
import { Link } from "react-router";
import { toast } from "sonner";
import { Plus, Edit, Trash2, Calendar, User, Heart } from "lucide-react";
import { useAuthStore } from "@/store/auth";

const Flowers = () => {
  const { data, isLoading } = useFlowers();
  const deleteFlower = useDeleteFlower();
  const token = useAuthStore((state) => state.token);

  if (isLoading) {
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
        {data && data.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {data.map((flower: FlowerType) => (
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
                      <span>User {flower.author_id}</span>
                    </div>
                    <div className="flex items-center gap-1">
                      <Heart className="w-4 h-4" />
                      <span>{flower.likes || 0} likes</span>
                    </div>
                  </div>
                </CardContent>

                {/* Actions */}
                {token && (
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
              No flowers yet
            </h3>
            <p className="text-gray-600 mb-8 max-w-md mx-auto">
              Be the first to share a beautiful flower with our community!
            </p>
            {token && (
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
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default Flowers;
