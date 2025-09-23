import { useQuery } from "@tanstack/react-query";
import api from "@/services/api";
import type { UserType } from "@/types/auth";

const useUser = () => {
  return useQuery<UserType>({
    queryKey: ["user"],
    queryFn: () => api.get("/api/flowers/user").then((res) => res.data),
  });
};

export { useUser };
