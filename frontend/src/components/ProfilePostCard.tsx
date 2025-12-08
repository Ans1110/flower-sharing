import { FlowerType } from "@/types/flower";
import { Card, CardContent } from "./ui/card";
import Link from "next/link";
import Image from "next/image";
import { Calendar, Heart, Pencil, Trash2 } from "lucide-react";
import { Button } from "./ui/button";
import { DeletePostDialog } from "./DeletePostDialog";

type ProfilePostCardProps = {
  post: FlowerType;
  isAuthenticated: boolean;
  isAuthor: boolean;
  isLiked: boolean;
  onDelete: () => void;
  onLike: () => void;
};

const ProfilePostCard = ({
  post,
  isAuthenticated,
  isAuthor,
  isLiked,
  onDelete,
  onLike,
}: ProfilePostCardProps) => {
  return (
    <Card className="group overflow-hidden border-0 shadow-lg hover:shadow-xl transition-all duration-300 bg-white dark:bg-zinc-900">
      {/* Image */}
      {post.image_url && (
        <Link href={`/flowers/${post.id}`} className="block relative">
          <div className="aspect-4/3 relative overflow-hidden">
            <Image
              src={post.image_url}
              alt={post.title}
              fill
              className="object-cover group-hover:scale-105 transition-transform duration-500"
            />
            <div className="absolute inset-0 bg-linear-to-t from-black/60 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
          </div>
        </Link>
      )}

      <CardContent className="p-4">
        {/* Title */}
        <Link href={`/flowers/${post.id}`}>
          <h3 className="font-semibold text-gray-900 dark:text-white line-clamp-1 group-hover:text-rose-600 dark:group-hover:text-rose-400 transition-colors">
            {post.title}
          </h3>
        </Link>

        {/* Meta */}
        <div className="flex items-center justify-between mt-3 pt-3 border-t border-gray-100 dark:border-zinc-800">
          <div className="flex items-center gap-2 text-xs text-muted-foreground">
            <Calendar className="h-3.5 w-3.5" />
            <span>
              {new Date(post.created_at).toLocaleDateString("en-US", {
                month: "short",
                day: "numeric",
              })}
            </span>
          </div>

          <div className="flex items-center gap-1">
            <Button
              variant="ghost"
              size="sm"
              onClick={onLike}
              disabled={!isAuthenticated}
              className={`h-7 px-2 ${
                isLiked
                  ? "text-rose-500"
                  : "text-muted-foreground hover:text-rose-500"
              }`}
            >
              <Heart
                className={`h-3.5 w-3.5 mr-1 ${isLiked ? "fill-current" : ""}`}
              />
              <span className="text-xs">{post.likes_count}</span>
            </Button>

            {isAuthor && (
              <>
                <Button
                  asChild
                  variant="ghost"
                  size="sm"
                  className="h-7 w-7 p-0 text-muted-foreground hover:text-blue-500"
                >
                  <Link href={`/flowers/${post.id}/edit`}>
                    <Pencil className="h-3.5 w-3.5" />
                  </Link>
                </Button>

                <DeletePostDialog
                  postTitle={post.title}
                  onDelete={onDelete}
                  trigger={
                    <Button
                      variant="ghost"
                      size="sm"
                      className="h-7 w-7 p-0 text-muted-foreground hover:text-red-500"
                    >
                      <Trash2 className="h-3.5 w-3.5" />
                    </Button>
                  }
                />
              </>
            )}
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default ProfilePostCard;
export { ProfilePostCard };
