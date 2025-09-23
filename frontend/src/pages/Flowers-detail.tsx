import { useParams, Link, useNavigate } from "react-router";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  useFlower,
  useDeleteFlower,
  useLikeFlower,
  useUnlikeFlower,
} from "@/hooks/api/flowers";
import {
  ArrowLeft,
  Calendar,
  User,
  Heart,
  Share2,
  Edit,
  Trash2,
} from "lucide-react";
import { useAuthStore } from "@/store/auth";
import { toast } from "sonner";
import { useState } from "react";
import { useUser } from "@/hooks/api/user";

const FlowerDetail = () => {
  const { id } = useParams<{ id: string }>();
  const { data, isLoading } = useFlower(id!);
  const token = useAuthStore((state) => state.token);
  const deleteFlower = useDeleteFlower();
  const likeFlower = useLikeFlower();
  const unlikeFlower = useUnlikeFlower();
  const navigate = useNavigate();
  const [isLiked, setIsLiked] = useState(false);
  const { data: user } = useUser();

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 flex items-center justify-center">
        <div className="flex flex-col items-center space-y-4">
          <div className="w-12 h-12 border-4 border-rose-200 border-t-rose-500 rounded-full animate-spin"></div>
          <p className="text-rose-600 font-medium">Loading flower details...</p>
        </div>
      </div>
    );
  }

  if (!data) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 flex items-center justify-center">
        <div className="text-center">
          <div className="w-32 h-32 bg-gradient-to-br from-rose-100 to-pink-100 rounded-full flex items-center justify-center mx-auto mb-6">
            <Heart className="w-16 h-16 text-rose-400" />
          </div>
          <h2 className="text-3xl font-bold text-gray-900 mb-4">
            Flower Not Found
          </h2>
          <p className="text-gray-600 mb-8">
            The flower you're looking for doesn't exist or has been removed.
          </p>
          <Button
            asChild
            className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
          >
            <Link to="/flowers">
              <ArrowLeft className="w-4 h-4 mr-2" />
              Back to Flowers
            </Link>
          </Button>
        </div>
      </div>
    );
  }

  const handleDelete = () => {
    if (window.confirm("Are you sure you want to delete this flower?")) {
      deleteFlower.mutate(id!, {
        onSuccess: () => {
          toast.success("Flower deleted successfully");
          navigate("/flowers");
        },
        onError: () => {
          toast.error("Failed to delete flower");
        },
      });
    }
  };

  const handleShare = () => {
    navigator.clipboard.writeText(`http://localhost:5173/flowers/${id}`);
    toast.success("Link copied to clipboard");
  };

  const handleLike = () => {
    if (!token) {
      toast.error("Please log in to like flowers");
      return;
    }

    if (isLiked) {
      unlikeFlower.mutate(id!, {
        onSuccess: () => {
          setIsLiked(false);
          toast.success("Removed like");
        },
        onError: (error) => {
          console.error("Unlike error:", error);
          toast.error("Failed to remove like");
        },
      });
    } else {
      likeFlower.mutate(id!, {
        onSuccess: () => {
          setIsLiked(true);
          toast.success("Liked flower!");
        },
        onError: () => {
          toast.error("Failed to like flower");
        },
      });
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 py-12">
      <div className="max-w-4xl mx-auto px-4">
        {/* Back Button */}
        <div className="mb-8">
          <Button
            asChild
            variant="outline"
            className="border-rose-200 text-rose-600 hover:bg-rose-50"
          >
            <Link to="/flowers">
              <ArrowLeft className="w-4 h-4 mr-2" />
              Back to Flowers
            </Link>
          </Button>
        </div>

        <div className="grid lg:grid-cols-2 gap-8">
          {/* Image Section */}
          <div className="space-y-4">
            {data.image_url ? (
              <div className="aspect-square overflow-hidden rounded-2xl shadow-2xl">
                <img
                  src={data.image_url}
                  alt={data.title}
                  className="w-full h-full object-cover hover:scale-105 transition-transform duration-500"
                />
              </div>
            ) : (
              <div className="aspect-square bg-gradient-to-br from-rose-100 to-pink-100 rounded-2xl shadow-2xl flex items-center justify-center">
                <Heart className="w-24 h-24 text-rose-400" />
              </div>
            )}

            {/* Action Buttons */}
            {token &&
              (user?.id === data.author_id || user?.role === "admin") && (
                <div className="flex gap-3">
                  <Button
                    asChild
                    className="flex-1 bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
                  >
                    <Link to={`/flowers/${id}/edit`}>
                      <Edit className="w-4 h-4 mr-2" />
                      Edit Flower
                    </Link>
                  </Button>
                  <Button
                    variant="destructive"
                    className="flex-1 border-rose-200 hover:bg-rose-500 text-white"
                    onClick={handleDelete}
                  >
                    <Trash2 className="w-4 h-4 mr-2" />
                    Delete
                  </Button>
                </div>
              )}
          </div>

          {/* Content Section */}
          <div className="space-y-6">
            <Card className="bg-white/80 backdrop-blur-sm border-0 shadow-xl">
              <CardHeader>
                <CardTitle className="text-3xl font-bold text-gray-900 mb-2">
                  {data.title}
                </CardTitle>
                <div className="flex items-center gap-4 text-sm text-gray-500">
                  <div className="flex items-center gap-1">
                    <Calendar className="w-4 h-4" />
                    <span>
                      {new Date(data.created_at).toLocaleDateString("en-US", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                      })}
                    </span>
                  </div>
                  <div className="flex items-center gap-1">
                    <User className="w-4 h-4" />
                    <span>{data.author_username}</span>
                  </div>
                </div>
              </CardHeader>
              <CardContent>
                <p className="text-gray-700 text-lg leading-relaxed whitespace-pre-wrap">
                  {data.content}
                </p>
              </CardContent>
            </Card>

            {/* Engagement Section */}
            <Card className="bg-white/80 backdrop-blur-sm border-0 shadow-xl">
              <CardContent className="pt-6">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-6">
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={handleLike}
                      className={`border-rose-200 hover:bg-rose-50 transition-all duration-200 ${
                        isLiked
                          ? "bg-rose-100 text-rose-700 border-rose-300"
                          : "text-rose-600"
                      }`}
                    >
                      <Heart
                        className={`w-4 h-4 mr-2 ${
                          isLiked ? "fill-current" : ""
                        }`}
                      />
                      {isLiked ? "Liked" : "Like"}{" "}
                      {data?.likes ? `(${data.likes})` : ""}
                    </Button>
                    <Button
                      variant="outline"
                      size="sm"
                      className="border-rose-200 text-rose-600 hover:bg-rose-50"
                      onClick={handleShare}
                    >
                      <Share2 className="w-4 h-4 mr-2" />
                      Share
                    </Button>
                  </div>
                  <div className="text-sm text-gray-500 ml-10">
                    Last updated:{" "}
                    {new Date(data.updated_at).toLocaleDateString()}
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

export default FlowerDetail;
