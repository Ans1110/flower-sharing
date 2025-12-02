import { UserType } from "./user";

type FlowerType = {
  id: number;
  title: string;
  content: string;
  image_url: string;
  created_at: string;
  updated_at: string;
  author: UserType;
  likes_count: number;
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
