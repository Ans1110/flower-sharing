import api from "@/services/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import type {
  FlowerType,
  FlowerPayloadType,
  PaginatedFlowersResponse,
} from "@/types/flower";

const useFlowers = (page: number = 1, limit: number = 6) => {
  return useQuery<PaginatedFlowersResponse>({
    queryKey: ["flowers", "paginated", page, limit],
    queryFn: async () => {
      const res = await api.get(`/api/flowers?page=${page}&limit=${limit}`);
      return res.data;
    },
  });
};

const useFlower = (id: string) => {
  return useQuery<FlowerType>({
    queryKey: ["flower", id],
    queryFn: async () => {
      const res = await api.get(`/api/flowers/${id}`);
      return res.data;
    },
    enabled: !!id,
  });
};

const useCreateFlower = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (payload: FlowerPayloadType) =>
      api.post("/api/flowers", payload).then((res) => res.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["flowers"] });
    },
  });
};

const useUpdateFlower = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (payload: FlowerPayloadType) =>
      api.put(`/api/flowers/${id}`, payload).then((res) => res.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["flowers"] });
      queryClient.invalidateQueries({ queryKey: ["flower", id] });
    },
  });
};

const useDeleteFlower = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      api.delete(`/api/flowers/${id}`).then((res) => res.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["flowers"] });
    },
  });
};

const useLikeFlower = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      api.post(`/api/flowers/${id}/like`).then((res) => res.data),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: ["flowers"] });
      queryClient.invalidateQueries({ queryKey: ["flower", id] });
    },
  });
};

const useUnlikeFlower = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      api.delete(`/api/flowers/${id}/unlike`).then((res) => res.data),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: ["flowers"] });
      queryClient.invalidateQueries({ queryKey: ["flower", id] });
    },
  });
};

const useSearchFlowers = (query: string) => {
  return useQuery<FlowerType[]>({
    queryKey: ["flowers", "search", query],
    queryFn: async () => {
      const res = await api.get(
        `/api/search?query=${encodeURIComponent(query)}`
      );
      return res.data;
    },
    enabled: query.length > 0,
  });
};

export {
  useFlowers,
  useFlower,
  useCreateFlower,
  useUpdateFlower,
  useDeleteFlower,
  useLikeFlower,
  useUnlikeFlower,
  useSearchFlowers,
};
