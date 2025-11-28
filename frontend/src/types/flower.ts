import { UserType } from "./user";

type FlowerType = {
  id: string;
  title: string;
  content: string;
  imageUrl: string;
  createdAt: string;
  updatedAt: string;
  author: UserType;
  likes: UserType[];
};

type FlowerPayloadType = {
  title: string;
  content: string;
  imageUrl: string;
  authorId: string;
};

type FlowerPaginationType = {
  page: number;
  limit: number;
};

type FlowerPaginationResponseType = {
  posts: FlowerType[];
  totalPages: number;
  page: number;
};

type FlowerResponseType = {
  post: FlowerType;
  message: string;
};

export type {
  FlowerType,
  FlowerPayloadType,
  FlowerPaginationType,
  FlowerPaginationResponseType,
  FlowerResponseType,
};
