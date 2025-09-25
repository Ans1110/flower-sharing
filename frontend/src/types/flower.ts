type FlowerType = {
  id: number;
  title: string;
  content: string;
  image_url: string;
  author_id: number;
  author_username: string;
  created_at: string;
  updated_at: string;
  likes: number;
};

type FlowerPayloadType = {
  title: string;
  content: string;
  image_url: string;
};

type PaginationType = {
  page: number;
  limit: number;
  total: number;
  total_pages: number;
};

type PaginatedFlowersResponse = {
  data: FlowerType[];
  pagination: PaginationType;
};

export type {
  FlowerType,
  FlowerPayloadType,
  PaginationType,
  PaginatedFlowersResponse,
};
