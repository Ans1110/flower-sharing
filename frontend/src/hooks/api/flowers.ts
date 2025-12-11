import { api } from "@/service/api";
import {
  FlowerPaginationResponseType,
  FlowerPaginationType,
  FlowerResponseType,
  FlowerType,
  FlowerAdminResponseType,
} from "@/types/flower";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { toast } from "sonner";

const useGetAllPosts = () => {
  return useQuery<FlowerAdminResponseType[], AxiosError<{ error: string }>>({
    queryKey: ["post-all"],
    queryFn: async () => {
      const res = await api.get<{ posts: FlowerAdminResponseType[] }>(
        `/post/all`
      );
      return res.data.posts;
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
    { post: FlowerType },
    AxiosError<{ error: string }>,
    FormData
  >({
    mutationFn: async (formData) => {
      const res = await api.post<{ post: FlowerType }>("/post", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
      return res.data;
    },
    onSuccess: (data) => {
      if (data?.post) {
        toast.success(`${data.post.title} created successfully`);
        queryClient.invalidateQueries({ queryKey: ["post", data.post.id] });
        queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
        queryClient.invalidateQueries({ queryKey: ["post-search"] });
        queryClient.invalidateQueries({ queryKey: ["post-all"] });
      }
    },
    onError: (error) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useUpdatePost = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { post: FlowerType },
    AxiosError<{ error: string }>,
    { postId: string; formData: FormData; selectFields: string[] }
  >({
    mutationFn: async ({ postId, formData, selectFields }) => {
      const res = await api.put<{ post: FlowerType }>(
        `/post/${postId}?select=${selectFields.join(",")}`,
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      return res.data;
    },
    onSuccess: (data, { postId }) => {
      if (data?.post) {
        toast.success(`${data.post.title} updated successfully`);
        queryClient.invalidateQueries({ queryKey: ["post", postId] });
        queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
        queryClient.invalidateQueries({ queryKey: ["post-search"] });
        queryClient.invalidateQueries({ queryKey: ["post-all"] });
      }
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
    mutationFn: async (postId) => {
      const res = await api.delete(`/post/${postId}`);
      return res.data;
    },
    onSuccess: ({ message }, postId) => {
      toast.success(message);
      // Invalidate all post-related queries
      queryClient.invalidateQueries({ queryKey: ["post", postId] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-likes"] });
      queryClient.invalidateQueries({ queryKey: ["posts-by-user"] });
      queryClient.invalidateQueries({ queryKey: ["post-user-liked"] });
      queryClient.invalidateQueries({ queryKey: ["user-following-posts"] });
      queryClient.invalidateQueries({ queryKey: ["user-liked-post-ids"] });
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
      const res = await api.get<{ posts: FlowerType[] }>(`/post/search`, {
        params: { query },
      });
      return res.data.posts;
    },
    enabled: query.length > 0,
  });
};

const useLikePost = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { message: string },
    AxiosError<{ error: string }>,
    string
  >({
    mutationFn: async (id) => {
      const res = await api.post<{ message: string }>(`/post/${id}/like`);
      return res.data;
    },
    onSuccess: (_, id) => {
      // Invalidate all queries that might contain this post
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-likes", id] });
      queryClient.invalidateQueries({ queryKey: ["posts-by-user"] });
      queryClient.invalidateQueries({ queryKey: ["post-user-liked"] });
      queryClient.invalidateQueries({ queryKey: ["user-following-posts"] });
      queryClient.invalidateQueries({ queryKey: ["user-liked-post-ids"] });
    },
    onError: (error) => {
      const errorMessage =
        error.response?.data?.error || "An unexpected error occurred";

      if (
        error.response?.status === 400 &&
        errorMessage === "post already liked"
      ) {
        return;
      }
      toast.error(errorMessage);
    },
  });
};

const useDislikePost = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { message: string },
    AxiosError<{ error: string }>,
    string
  >({
    mutationFn: async (id) => {
      const res = await api.delete<{ message: string }>(`/post/${id}/dislike`);
      return res.data;
    },
    onSuccess: (_, id) => {
      // Invalidate all queries that might contain this post
      queryClient.invalidateQueries({ queryKey: ["post", id] });
      queryClient.invalidateQueries({ queryKey: ["post-pagination"] });
      queryClient.invalidateQueries({ queryKey: ["post-all"] });
      queryClient.invalidateQueries({ queryKey: ["post-search"] });
      queryClient.invalidateQueries({ queryKey: ["post-likes", id] });
      queryClient.invalidateQueries({ queryKey: ["posts-by-user"] });
      queryClient.invalidateQueries({ queryKey: ["post-user-liked"] });
      queryClient.invalidateQueries({ queryKey: ["user-following-posts"] });
      queryClient.invalidateQueries({ queryKey: ["user-liked-post-ids"] });
    },
    onError: (error) => {
      const errorMessage =
        error.response?.data?.error || "An unexpected error occurred";
      toast.error(errorMessage);
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

// Hook to get all liked post IDs for the current user (for checking if a post is liked)
const useGetUserLikedPostIds = (userId: number | undefined) => {
  return useQuery<Set<number>, AxiosError<{ error: string }>>({
    queryKey: ["user-liked-post-ids", userId],
    queryFn: async () => {
      // Fetch all liked posts (use high limit to get all)
      const res = await api.get<FlowerPaginationResponseType>(
        `/post/user/${userId}/liked`,
        {
          params: { page: 1, limit: 1000 },
        }
      );
      // Return a Set of post IDs for O(1) lookup
      return new Set(res.data.posts.map((post) => post.id));
    },
    enabled: !!userId,
    staleTime: 30000, // Cache for 30 seconds
  });
};

const useGetPostsByUserId = (userId: string) => {
  return useQuery<FlowerType[], AxiosError<{ error: string }>>({
    queryKey: ["posts-by-user", userId],
    queryFn: async () => {
      const res = await api.get<{ posts: FlowerType[] }>(
        `/post/user/${userId}/all`
      );
      return res.data.posts;
    },
    enabled: !!userId,
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
  useGetUserLikedPostIds,
  useGetPostsByUserId,
};
