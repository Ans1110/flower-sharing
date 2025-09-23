import api from "@/services/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

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

const useFlowers = () => {
  return useQuery<FlowerType[]>({
    queryKey: ["flowers"],
    queryFn: async () => {
      const res = await api.get("/api/flowers");
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

export {
  useFlowers,
  useFlower,
  useCreateFlower,
  useUpdateFlower,
  useDeleteFlower,
  useLikeFlower,
  useUnlikeFlower,
  type FlowerType,
  type FlowerPayloadType,
};
