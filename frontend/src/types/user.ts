type UserType = {
  id: number;
  username: string;
  email: string;
  role: "user" | "admin";
  avatar: string;
};

export type { UserType };
