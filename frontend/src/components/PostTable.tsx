import { FlowerAdminResponseType } from "@/types/flower";
import { Edit, FileText, Heart, Trash2 } from "lucide-react";
import Link from "next/link";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { formatDate, getUserInitials } from "@/lib/utils";
import { Button } from "./ui/button";

type PostTableProps = {
  posts: FlowerAdminResponseType[];
  onDeletePost: (postId: number) => void;
};

const PostTable = ({ posts, onDeletePost }: PostTableProps) => {
  if (posts.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-12">
        <FileText className="mb-4 h-12 w-12 text-muted-foreground/50" />
        <p className="text-sm font-medium text-muted-foreground">
          No posts found
        </p>
        <p className="text-xs text-muted-foreground">
          Posts will appear here once they are created
        </p>
      </div>
    );
  }

  return (
    <div className="overflow-hidden">
      <Table>
        <TableHeader>
          <TableRow className="bg-muted/50">
            <TableHead className="w-[80px]">ID</TableHead>
            <TableHead className="w-[120px]">Title</TableHead>
            <TableHead className="w-[100px] text-center">Author</TableHead>
            <TableHead className="w-[100px] text-center">Likes</TableHead>
            <TableHead className="w-[100px]">Created</TableHead>
            <TableHead className="w-[120px] text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {posts.map((post) => (
            <TableRow key={post.id} className="hover:bg-muted/50">
              <TableCell className="font-medium text-muted-foreground">
                #{post.id}
              </TableCell>
              <TableCell>
                <Link
                  href={`/flowers/${post.id}`}
                  className="flex items-center gap-2 hover:underline"
                >
                  <FileText className="h-4 w-4 text-muted-foreground shrink-0" />
                  <p className="line-clamp-2">{post.title}</p>
                </Link>
              </TableCell>
              <TableCell>
                <div className="flex items-center gap-3">
                  <Avatar className="h-9 w-9">
                    {post.author.avatar && (
                      <Link href={`/profile/${post.author.id}`}>
                        <AvatarImage
                          src={post.author.avatar}
                          alt={post.author.username}
                          className="object-cover"
                        />
                      </Link>
                    )}
                    <AvatarFallback className="bg-primary/10 text-xs font-semibold">
                      {getUserInitials(post.author.username)}
                    </AvatarFallback>
                  </Avatar>
                  <Link
                    href={`/profile/${post.author.id}`}
                    className="hover:underline"
                  >
                    <p className="line-clamp-1 hover:underline">
                      {post.author.username}
                    </p>
                  </Link>
                </div>
              </TableCell>
              <TableCell className="text-center">
                <div className="flex items-center justify-center gap-1.5">
                  <Heart className="h-3.5 w-3.5 text-muted-foreground" />
                  <span className="font-medium">{post.likes_count}</span>
                </div>
              </TableCell>
              <TableCell className="text-sm text-muted-foreground">
                {formatDate(post.created_at)}
              </TableCell>
              <TableCell>
                <div className="flex items-center justify-end gap-2">
                  <Link href={`/admin/posts/${post.id}/edit`}>
                    <Button variant="outline" size="sm" className="h-8 gap-1.5">
                      <Edit className="h-3.5 w-3.5" />
                      Edit
                    </Button>
                  </Link>
                  <Button
                    variant="destructive"
                    size="sm"
                    className="h-8 gap-1.5"
                    onClick={() => onDeletePost(post.id)}
                  >
                    <Trash2 className="h-3.5 w-3.5" />
                    Delete
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default PostTable;
export { PostTable };
