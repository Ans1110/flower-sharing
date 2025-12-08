type UserType = {
  id: number;
  username: string;
  email: string;
  role: "user" | "admin";
  avatar: string | null;
};

type UserPublicResponseType = {
  id: number;
  username: string;
  avatar: string;
};

type UserAdminResponseType = {
  id: number;
  username: string;
  email: string;
  role: "user" | "admin";
  avatar: string;
  posts: number;
  likes: number;
  followers: number;
  following: number;
  created_at: string;
};

type UserPayloadType = {
  username: string;
  email: string;
  avatar: File | string | null;
};

export type {
  UserType,
  UserPublicResponseType,
  UserAdminResponseType,
  UserPayloadType,
};
