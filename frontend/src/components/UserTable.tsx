import { UserAdminResponseType } from "@/types/user";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { Button } from "./ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { Badge } from "./ui/badge";
import { Edit, Trash2, FileText, UserCircle } from "lucide-react";
import Link from "next/link";

type UserTableProps = {
  users: UserAdminResponseType[];
  onDeleteUser: (userId: number) => void;
};

const UserTable = ({ users, onDeleteUser }: UserTableProps) => {
  if (users.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-12">
        <UserCircle className="mb-4 h-12 w-12 text-muted-foreground/50" />
        <p className="text-sm font-medium text-muted-foreground">
          No users found
        </p>
        <p className="text-xs text-muted-foreground">
          Users will appear here once they register
        </p>
      </div>
    );
  }

  const formatDate = (dateString: string) => {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (Number.isNaN(date.getTime())) {
      return "N/A";
    }

    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric",
    });
  };

  const getInitials = (username: string) => {
    return username.slice(0, 2).toUpperCase();
  };

  return (
    <div className="overflow-hidden">
      <Table>
        <TableHeader>
          <TableRow className="bg-muted/50">
            <TableHead className="w-[80px]">ID</TableHead>
            <TableHead>User</TableHead>
            <TableHead>Email</TableHead>
            <TableHead className="w-[100px]">Role</TableHead>
            <TableHead className="w-[100px] text-center">Posts</TableHead>
            <TableHead className="w-[140px]">Joined</TableHead>
            <TableHead className="w-[180px] text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {users.map((user) => (
            <TableRow key={user.id} className="hover:bg-muted/50">
              <TableCell className="font-medium text-muted-foreground">
                #{user.id}
              </TableCell>
              <TableCell>
                <div className="flex items-center gap-3">
                  <Avatar className="h-9 w-9">
                    <Link href={`/profile/${user.id}`}>
                      <AvatarImage
                        src={user.avatar}
                        alt={user.username}
                        className="object-cover"
                      />
                    </Link>
                    <AvatarFallback className="bg-primary/10 text-xs font-semibold">
                      {getInitials(user.username)}
                    </AvatarFallback>
                  </Avatar>
                  <div>
                    <p className="font-medium">{user.username}</p>
                  </div>
                </div>
              </TableCell>
              <TableCell className="text-sm text-muted-foreground">
                {user.email}
              </TableCell>
              <TableCell>
                <Badge
                  variant={user.role === "admin" ? "default" : "secondary"}
                  className="capitalize"
                >
                  {user.role}
                </Badge>
              </TableCell>
              <TableCell className="text-center">
                <div className="flex items-center justify-center gap-1.5">
                  <FileText className="h-3.5 w-3.5 text-muted-foreground" />
                  <span className="font-medium">{user.posts}</span>
                </div>
              </TableCell>
              <TableCell className="text-sm text-muted-foreground">
                {formatDate(user.createdAt)}
              </TableCell>
              <TableCell>
                <div className="flex items-center justify-end gap-2">
                  <Link href={`/admin/users/${user.id}/edit`}>
                    <Button variant="outline" size="sm" className="h-8 gap-1.5">
                      <Edit className="h-3.5 w-3.5" />
                      Edit
                    </Button>
                  </Link>
                  <Button
                    variant="destructive"
                    size="sm"
                    className="h-8 gap-1.5"
                    onClick={() => onDeleteUser(user.id)}
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

export default UserTable;
