import { FlowerType } from "@/types/flower";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import Image from "next/image";
import Link from "next/link";
import { Calendar, Heart, Pencil, Trash2, User } from "lucide-react";
import { Button } from "./ui/button";
import { cn } from "@/lib/utils";

type FlowerCardProps = {
  flower: FlowerType;
  isAuthenticated: boolean;
  isAuthor: boolean;
  isLiked?: boolean;
  onDelete: () => void;
  onLike: () => void;
};

const FlowerCard = ({
  flower,
  isAuthenticated,
  isAuthor,
  isLiked = false,
  onDelete,
  onLike,
}: FlowerCardProps) => {
  return (
    <Card
      key={flower.id}
      className="group relative bg-white/90 dark:bg-zinc-900/90 backdrop-blur-md border border-rose-100 dark:border-rose-900/30 shadow-xl hover:shadow-2xl transition-all duration-500 hover:-translate-y-2 overflow-hidden rounded-2xl"
    >
      {/* Image */}
      {flower.image_url && (
        <div className="relative aspect-square overflow-hidden">
          <Image
            src={flower.image_url}
            alt={flower.title}
            sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
            fill
            loading="eager"
            className="object-cover group-hover:scale-110 transition-transform duration-700 ease-out"
          />
          {/* Gradient Overlay */}
          <Link href={`/flowers/${flower.id}`} className="absolute inset-0">
            <div className="absolute inset-0 bg-linear-to-t from-black/50 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />
          </Link>
        </div>
      )}

      {/* Content */}
      <CardHeader className="pb-3">
        <CardTitle className="text-xl font-bold text-gray-900 dark:text-gray-100 line-clamp-2 group-hover:text-rose-600 dark:group-hover:text-rose-400 transition-colors duration-300">
          <Link href={`/flowers/${flower.id}`} className="hover:underline">
            {flower.title}
          </Link>
        </CardTitle>
      </CardHeader>

      <CardContent className="space-y-4 pb-4">
        <p className="text-gray-600 dark:text-gray-400 line-clamp-3 leading-relaxed text-sm">
          {flower.content}
        </p>

        {/* Meta Info */}
        <div className="flex flex-wrap items-center gap-3 pt-2 border-t border-rose-100 dark:border-rose-900/30">
          <div className="flex items-center gap-1.5 text-xs text-gray-500 dark:text-gray-400">
            <Calendar className="size-3.5 text-rose-500 dark:text-rose-400" />
            <span>
              {new Date(flower.created_at).toLocaleDateString("en-US", {
                month: "short",
                day: "numeric",
              })}
            </span>
          </div>
          <div className="flex items-center gap-1.5 text-xs text-gray-500 dark:text-gray-400">
            <User className="size-3.5 text-rose-500 dark:text-rose-400" />
            <span className="font-medium">
              {isAuthor ? "You" : flower.author?.username || "Unknown"}
            </span>
          </div>
          <div className="flex items-center gap-1 ml-auto">
            <Button
              disabled={!isAuthenticated}
              variant="ghost"
              size="sm"
              onClick={onLike}
              className={cn(
                "h-7 px-2 transition-colors",
                isLiked
                  ? "text-rose-600 dark:text-rose-400 hover:bg-rose-50 dark:hover:bg-rose-900/30"
                  : "hover:bg-rose-50 dark:hover:bg-rose-900/30 hover:text-rose-600 dark:hover:text-rose-400"
              )}
            >
              <Heart
                className={cn(
                  "size-4 mr-1",
                  isLiked && "fill-rose-500 text-rose-500"
                )}
              />
              <span className="text-xs font-medium">{flower.likes_count}</span>
            </Button>
          </div>
        </div>
      </CardContent>

      {/* Actions */}
      {isAuthenticated && isAuthor && (
        <CardFooter className="pt-0 pb-4 gap-2">
          <Button
            asChild
            variant="outline"
            size="sm"
            className="flex-1 border-rose-200 text-rose-600 hover:bg-rose-50 hover:border-rose-300 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/30 transition-all"
          >
            <Link href={`/flowers/${flower.id}/edit`}>
              <Pencil className="size-3.5 mr-1.5" />
              Edit
            </Link>
          </Button>
          <Button
            variant="outline"
            size="sm"
            className="flex-1 border-red-200 text-red-600 hover:bg-red-50 hover:border-red-300 dark:border-red-800 dark:text-red-400 dark:hover:bg-red-900/30 transition-all"
            onClick={onDelete}
          >
            <Trash2 className="size-3.5 mr-1.5" />
            Delete
          </Button>
        </CardFooter>
      )}
    </Card>
  );
};

export default FlowerCard;
export { FlowerCard };
