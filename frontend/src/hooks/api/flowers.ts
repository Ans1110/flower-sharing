import { api } from "@/service/api";
import {
  FlowerPaginationResponseType,
  FlowerPaginationType,
  FlowerPayloadType,
  FlowerResponseType,
  FlowerType,
} from "@/types/flower";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { toast } from "sonner";

const useGetAllPosts = () => {
  return useQuery<FlowerType[], AxiosError<{ error: string }>>({
    queryKey: ["post-all"],
    queryFn: async () => {
      const res = await api.get<FlowerType[]>(`/post/all`);
      return res.data;
    },
  });
};

const useGetPostById = (id: string) => {
  return useQuery<FlowerResponseType, AxiosError<{ error: string }>>({
    queryKey: ["post", id],
    queryFn: async () => {
      const res = await api.get<FlowerResponseType>(`/post/${id}`);
      return res.data;
    },
    enabled: !!id,
  });
};

const useGetPostsPagination = ({ page, limit }: FlowerPaginationType) => {
  return useQuery<FlowerPaginationResponseType, AxiosError<{ error: string }>>({
    queryKey: ["post-pagination", page, limit],
    queryFn: async () => {
      const res = await api.get<FlowerPaginationResponseType>(
        `/post/pagination`,
        {
          params: { page, limit },
        }
      );
      return res.data;
    },
    enabled: !!page && !!limit,
  });
};

const useCreatePost = () => {
  const queryClient = useQueryClient();
  return useMutation<
    FlowerResponseType,
    AxiosError<{ error: string }>,
    FlowerPayloadType
  >({
    mutationFn: async (payload) => {
      const res = await api.post<FlowerResponseType>("/post", payload);
      return res.data;
    },
    onSuccess: ({ post }) => {
      toast.success(`${post.title} created successfully`);
      queryClient.invalidateQueries({ queryKey: ["post", post.id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useUpdatePost = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation<
    FlowerResponseType,
    AxiosError<{ error: string }>,
    FlowerPayloadType
  >({
    mutationFn: async (payload) => {
      const res = await api.put<FlowerResponseType>(`/post/${id}`, payload);
      return res.data;
    },
    onSuccess: ({ post }) => {
      toast.success(`${post.title} updated successfully`);
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useDeletePost = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { message: string },
    AxiosError<{ error: string }>,
    string
  >({
    mutationFn: async (id) => {
      const res = await api.delete(`/post/${id}`);
      return res.data;
    },
    onSuccess: ({ message }, id) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useSearchPosts = (query: string) => {
  return useQuery<FlowerType[], AxiosError<{ error: string }>>({
    queryKey: ["post-search", query],
    queryFn: async () => {
      const res = await api.get<FlowerType[]>(`/post/search`, {
        params: { query },
      });
      return res.data;
    },
    enabled: query.length > 0,
  });
};

const useLikePost = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation<{ message: string }, AxiosError<{ error: string }>, void>({
    mutationFn: async () => {
      const res = await api.post<{ message: string }>(`/post/${id}/like`);
      return res.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useDislikePost = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation<{ message: string }, AxiosError<{ error: string }>, void>({
    mutationFn: async () => {
      const res = await api.delete<{ message: string }>(`/post/${id}/dislike`);
      return res.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useGetPostLikes = (id: string) => {
  return useQuery<number, AxiosError<{ error: string }>>({
    queryKey: ["post-likes", id],
    queryFn: async () => {
      const res = await api.get<{ likes: number }>(`/post/${id}/likes`);
      return res.data.likes;
    },
    enabled: !!id,
  });
};

const useGetUserLikedPosts = (userId: string, page: number, limit: number) => {
  return useQuery<FlowerPaginationResponseType, AxiosError<{ error: string }>>({
    queryKey: ["post-user-liked", userId, page, limit],
    queryFn: async () => {
      const res = await api.get<FlowerPaginationResponseType>(
        `/post/user/${userId}/liked`,
        {
          params: { page, limit },
        }
      );
      return res.data;
    },
    enabled: !!userId && !!page && !!limit,
  });
};

export {
  useGetAllPosts,
  useGetPostById,
  useGetPostsPagination,
  useCreatePost,
  useUpdatePost,
  useDeletePost,
  useSearchPosts,
  useLikePost,
  useDislikePost,
  useGetPostLikes,
  useGetUserLikedPosts,
};
