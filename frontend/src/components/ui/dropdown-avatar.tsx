"use client";

import { LogOut } from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "./avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "./dropdown-menu";

type DropdownAvatarProps = {
  username?: string;
  email?: string;
  avatar?: string;
  onSignOut: () => void;
};

export default function DropdownAvatar({
  username = "",
  email,
  avatar,
  onSignOut,
}: DropdownAvatarProps) {
  // Get initials from username for fallback
  const getInitials = (name: string) => {
    return name
      .split(" ")
      .map((n) => n[0])
      .join("")
      .toUpperCase()
      .slice(0, 2);
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button className="relative focus:outline-none focus-visible:ring-2 focus-visible:ring-rose-500 focus-visible:ring-offset-2 rounded-full">
          <Avatar className="size-10 cursor-pointer hover:opacity-80 transition-opacity">
            <AvatarImage src={avatar} alt={username} className="object-cover" />
            <AvatarFallback className="bg-linear-to-br from-rose-500 to-pink-500 text-white font-semibold">
              {getInitials(username)}
            </AvatarFallback>
          </Avatar>
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="end" forceMount>
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col space-y-1">
            <p className="text-sm font-medium leading-none">{username}</p>
            <p className="text-xs leading-none text-muted-foreground">
              {email}
            </p>
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onClick={onSignOut}
          variant="destructive"
          className="cursor-pointer"
        >
          <LogOut />
          Sign out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export { DropdownAvatar };
