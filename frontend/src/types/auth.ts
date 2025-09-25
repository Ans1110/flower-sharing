type LoginPayloadType = {
  email: string;
  password: string;
};

type RegisterPayloadType = {
  username: string;
  email: string;
  password: string;
};

type UserType = {
  id: number;
  username: string;
  email: string;
  role: "user" | "admin";
};

export type { LoginPayloadType, RegisterPayloadType, UserType };
