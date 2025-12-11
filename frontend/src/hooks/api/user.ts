import { api } from "@/service/api";
import { useAuthStore } from "@/store/auth";
import { FlowerPaginationResponseType } from "@/types/flower";
import {
  UserAdminResponseType,
  UserPublicResponseType,
  UserType,
} from "@/types/user";
import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseQueryOptions,
} from "@tanstack/react-query";
import { AxiosError } from "axios";
import { toast } from "sonner";

const useGetUserAll = () => {
  return useQuery<UserAdminResponseType[], AxiosError<{ error: string }>>({
    queryKey: ["user-all"],
    queryFn: async () => {
      const res = await api.get<{ users: UserAdminResponseType[] }>(
        "/admin/user/all"
      );
      return res.data.users;
    },
  });
};

const useGetUserById = (id: string) => {
  return useQuery<
    UserAdminResponseType | UserPublicResponseType,
    AxiosError<{ error: string }>
  >({
    queryKey: ["user-id", id],
    queryFn: async () => {
      const isAdmin = useAuthStore.getState().user?.role === "admin";
      if (isAdmin) {
        const res = await api.get<{ user: UserAdminResponseType }>(
          `/admin/user/${id}`
        );
        return res.data.user;
      } else {
        const res = await api.get<{ user: UserPublicResponseType }>(
          `/user/${id}`
        );
        return res.data.user;
      }
    },
    enabled: !!id,
    retry: (failureCount, error) => {
      // Don't retry if user not found (404)
      if (error.response?.status === 404) {
        return false;
      }
      return failureCount < 3;
    },
  });
};

const useGetUserByUsername = (username: string) => {
  return useQuery<
    UserAdminResponseType | UserPublicResponseType,
    AxiosError<{ error: string }>
  >({
    queryKey: ["user-username", username],
    queryFn: async () => {
      const isAdmin = useAuthStore.getState().user?.role === "admin";
      if (isAdmin) {
        const res = await api.get<{ user: UserAdminResponseType }>(
          `/admin/user/username/${username}`
        );
        return res.data.user;
      } else {
        const res = await api.get<{ user: UserPublicResponseType }>(
          `/user/username/${username}`
        );
        return res.data.user;
      }
    },
    enabled: !!username,
    meta: {
      onError: (error: AxiosError<{ error: string }>) => {
        toast.error(
          error.response?.data?.error || "An unexpected error occurred"
        );
      },
    },
  });
};

const useUpdateUserById = () => {
  const queryClient = useQueryClient();
  return useMutation<
    UserAdminResponseType | UserPublicResponseType,
    AxiosError<{ error: string }>,
    { userId: number; formData: FormData; selectFields: string[] }
  >({
    mutationFn: async ({ userId, formData, selectFields }) => {
      const isAdmin = useAuthStore.getState().user?.role === "admin";
      const endpoint = isAdmin
        ? `/admin/user/id/${userId}/select?select=${selectFields.join(",")}`
        : `/user/id/${userId}/select?select=${selectFields.join(",")}`;

      if (isAdmin) {
        const res = await api.put<{ user: UserAdminResponseType }>(
          endpoint,
          formData
        );
        return res.data.user;
      } else {
        const res = await api.put<{ user: UserPublicResponseType }>(
          endpoint,
          formData
        );
        return res.data.user;
      }
    },
    onSuccess: (data, { userId }) => {
      toast.success(`User updated successfully`);
      // Remove old cache and refetch user queries
      queryClient.removeQueries({
        queryKey: ["user-id", userId.toString()],
      });
      queryClient.invalidateQueries({
        queryKey: ["user-username", data.username],
      });
      queryClient.invalidateQueries({ queryKey: ["user-all"] });
    },
    onError: (error: AxiosError<{ error: string }>) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useDeleteUserById = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { message: string },
    AxiosError<{ error: string }>,
    number
  >({
    mutationFn: async (userId) => {
      const res = await api.delete<{ message: string }>(
        `/admin/user/${userId}`
      );
      return res.data;
    },
    onSuccess: ({ message }, userId) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["user-id", userId] });
      queryClient.invalidateQueries({ queryKey: ["user-all"] });
    },
    onError: (error: AxiosError<{ error: string }>) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useFollowUser = (followerId: string, followingId: string) => {
  const queryClient = useQueryClient();
  return useMutation<{ message: string }, AxiosError<{ error: string }>, void>({
    mutationFn: async () => {
      const res = await api.post<{ message: string }>(
        `/user/follow/${followerId}/${followingId}`
      );
      return res.data;
    },
    onSuccess: ({ message }) => {
      toast.success(message);
      // Invalidate user queries
      queryClient.invalidateQueries({ queryKey: ["user-id", followerId] });
      queryClient.invalidateQueries({ queryKey: ["user-id", followingId] });
      // Invalidate followers/following lists
      queryClient.invalidateQueries({
        queryKey: ["user-following", followerId],
      });
      queryClient.invalidateQueries({
        queryKey: ["user-followers", followingId],
      });
    },
    onError: (error: AxiosError<{ error: string }>) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useUnfollowUser = (followerId: string, followingId: string) => {
  const queryClient = useQueryClient();
  return useMutation<{ message: string }, AxiosError<{ error: string }>, void>({
    mutationFn: async () => {
      const res = await api.post<{ message: string }>(
        `/user/unfollow/${followerId}/${followingId}`
      );
      return res.data;
    },
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["user-id", followerId] });
      queryClient.invalidateQueries({ queryKey: ["user-id", followingId] });
      // Invalidate followers/following lists
      queryClient.invalidateQueries({
        queryKey: ["user-following", followerId],
      });
      queryClient.invalidateQueries({
        queryKey: ["user-followers", followingId],
      });
    },
    onError: (error: AxiosError<{ error: string }>) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

// Dynamic follow hook - accepts followingId as mutation parameter
const useToggleFollow = () => {
  const queryClient = useQueryClient();
  return useMutation<
    { message: string },
    AxiosError<{ error: string }>,
    { followerId: string; followingId: string; isFollowing: boolean }
  >({
    mutationFn: async ({ followerId, followingId, isFollowing }) => {
      const endpoint = isFollowing
        ? `/user/unfollow/${followerId}/${followingId}`
        : `/user/follow/${followerId}/${followingId}`;
      const res = await api.post<{ message: string }>(endpoint);
      return res.data;
    },
    onSuccess: ({ message }, { followerId, followingId }) => {
      toast.success(message);
      queryClient.invalidateQueries({
        queryKey: ["user-following", followerId],
      });
      queryClient.invalidateQueries({
        queryKey: ["user-followers", followingId],
      });
    },
    onError: (error: AxiosError<{ error: string }>) => {
      toast.error(
        error.response?.data?.error || "An unexpected error occurred"
      );
    },
  });
};

const useGetUserFollowers = (userId: string) => {
  return useQuery<UserPublicResponseType[], AxiosError<{ error: string }>>({
    queryKey: ["user-followers", userId],
    queryFn: async () => {
      const res = await api.get<{ followers: UserPublicResponseType[] }>(
        `/user/followers/${userId}`
      );
      return res.data.followers;
    },
    enabled: !!userId,
    retry: (failureCount, error) => {
      // Don't retry if user not found (404)
      if (error.response?.status === 404) {
        return false;
      }
      return failureCount < 3;
    },
  });
};

const useGetUserFollowing = (userId: string) => {
  return useQuery<UserPublicResponseType[], AxiosError<{ error: string }>>({
    queryKey: ["user-following", userId],
    queryFn: async () => {
      const res = await api.get<{ following: UserPublicResponseType[] }>(
        `/user/following/${userId}`
      );
      return res.data.following;
    },
    enabled: !!userId,
    retry: (failureCount, error) => {
      // Don't retry if user not found (404)
      if (error.response?.status === 404) {
        return false;
      }
      return failureCount < 3;
    },
  });
};

const useGetUserFollowersCount = (userId: string) => {
  return useQuery<number, AxiosError<{ error: string }>>({
    queryKey: ["user-followers-count", userId],
    queryFn: async () => {
      const res = await api.get<{ followers_count: number }>(
        `/user/followers-count/${userId}`
      );
      return res.data.followers_count;
    },
    enabled: !!userId,
  });
};

const useGetUserFollowingCount = (userId: string) => {
  return useQuery<number, AxiosError<{ error: string }>>({
    queryKey: ["user-following-count", userId],
    queryFn: async () => {
      const res = await api.get<{ following_count: number }>(
        `/user/following-count/${userId}`
      );
      return res.data.following_count;
    },
    enabled: !!userId,
  });
};

const useGetUserFollowingPosts = (
  userId: string,
  page: number,
  limit: number
) => {
  return useQuery<FlowerPaginationResponseType, AxiosError<{ error: string }>>({
    queryKey: ["user-following-posts", userId, page, limit],
    queryFn: async () => {
      const res = await api.get<FlowerPaginationResponseType>(
        `/user/following-posts/${userId}`,
        {
          params: { page, limit },
        }
      );
      return res.data;
    },
    enabled: !!userId && !!page && !!limit,
    meta: {
      onError: (error: AxiosError<{ error: string }>) => {
        toast.error(
          error.response?.data?.error || "An unexpected error occurred"
        );
      },
    },
  });
};

const useGetMe = (
  options?: Pick<
    UseQueryOptions<UserType, AxiosError<{ error: string }>, UserType, ["me"]>,
    "enabled" | "retry" | "meta"
  >
) => {
  return useQuery<UserType, AxiosError<{ error: string }>, UserType, ["me"]>({
    queryKey: ["me"],
    queryFn: async () => {
      const res = await api.get<{ user: UserType }>("/auth/me");
      return res.data.user;
    },
    retry: 1,
    meta: {
      onError: (error: AxiosError<{ error: string }>) => {
        toast.error(
          error.response?.data?.error || "Failed to fetch user information"
        );
      },
    },
    ...options,
  });
};

export {
  useGetUserAll,
  useGetUserById,
  useGetUserByUsername,
  useUpdateUserById,
  useDeleteUserById,
  useFollowUser,
  useUnfollowUser,
  useToggleFollow,
  useGetUserFollowers,
  useGetUserFollowing,
  useGetUserFollowersCount,
  useGetUserFollowingCount,
  useGetUserFollowingPosts,
  useGetMe,
};
