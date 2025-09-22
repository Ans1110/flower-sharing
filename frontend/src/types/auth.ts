type LoginPayloadType = {
  email: string;
  password: string;
};

type RegisterPayloadType = {
  username: string;
  email: string;
  password: string;
};

export type { LoginPayloadType, RegisterPayloadType };
