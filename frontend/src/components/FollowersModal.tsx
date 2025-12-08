import { UserPublicResponseType } from "@/types/user";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "./ui/dialog";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { getUserInitials } from "@/lib/utils";
import { Button } from "./ui/button";
import { Users } from "lucide-react";

type FollowersModalProps = {
  isOpen: boolean;
  onClose: () => void;
  onNavigate?: (path: string) => void;
  users: UserPublicResponseType[];
  title: string;
  description: string;
};

export const FollowersModal = ({
  isOpen,
  onClose,
  onNavigate,
  users,
  title,
  description,
}: FollowersModalProps) => {
  const handleNavigate = (userId: number) => {
    const path = `/profile/${userId}`;
    if (onNavigate) {
      onNavigate(path);
    } else {
      // Fallback to hard navigation
      window.location.assign(path);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[500px] max-h-[600px]">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2">
            <Users className="h-5 w-5 text-rose-500" />
            {title}
          </DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>

        <div className="mt-4 max-h-[400px] overflow-y-auto">
          {users.length === 0 ? (
            <div className="flex flex-col items-center justify-center py-12 text-center">
              <div className="w-16 h-16 bg-rose-100 dark:bg-rose-900/30 rounded-full flex items-center justify-center mb-4">
                <Users className="h-8 w-8 text-rose-500" />
              </div>
              <p className="text-muted-foreground">No users to show</p>
            </div>
          ) : (
            <div className="space-y-3">
              {users.map((user) => (
                <div
                  key={user.id}
                  onClick={() => handleNavigate(user.id)}
                  className="flex items-center gap-3 p-3 rounded-lg hover:bg-rose-50 dark:hover:bg-rose-900/20 transition-colors group cursor-pointer"
                >
                  <Avatar className="h-12 w-12 ring-2 ring-rose-100 dark:ring-rose-900/30">
                    <AvatarImage
                      src={user.avatar}
                      alt={user.username}
                      className="object-cover"
                    />
                    <AvatarFallback className="bg-linear-to-br from-rose-400 to-violet-500 text-white font-semibold">
                      {getUserInitials(user.username)}
                    </AvatarFallback>
                  </Avatar>
                  <div className="flex-1 min-w-0">
                    <p className="font-semibold text-gray-900 dark:text-white group-hover:text-rose-600 dark:group-hover:text-rose-400 transition-colors truncate">
                      {user.username}
                    </p>
                    <p className="text-sm text-muted-foreground">
                      View profile
                    </p>
                  </div>
                  <Button
                    variant="ghost"
                    size="sm"
                    className="opacity-0 group-hover:opacity-100 transition-opacity"
                    asChild
                  >
                    <span className="text-xs text-rose-600 dark:text-rose-400">
                      →
                    </span>
                  </Button>
                </div>
              ))}
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default FollowersModal;
