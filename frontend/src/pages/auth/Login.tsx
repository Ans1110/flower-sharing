import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import api from "@/services/api";
import { useAuthStore } from "@/store/auth";
import type { LoginPayloadType } from "@/types/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { NavLink, useNavigate } from "react-router-dom";
import { toast } from "sonner";
import z from "zod";

const schema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
});

const Login = () => {
  const setToken = useAuthStore((state) => state.setToken);
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginPayloadType>({
    resolver: zodResolver(schema),
  });

  const mutation = useMutation({
    mutationFn: (data: LoginPayloadType) =>
      api.post("/login", data).then((res) => res.data),
    onSuccess: (data) => {
      setToken(data.token);
      toast.success("Login successful");
      navigate("/");
    },
    onError: () => {
      toast.error("Login failed");
    },
  });

  return (
    <div className="max-w-md mx-auto mt-32 p-6 border rounded-lg shadow">
      <h1 className="text-2xl font-bold mb-4 text-rose-400">Login</h1>
      <form
        onSubmit={handleSubmit((data) => mutation.mutate(data))}
        className="space-y-4"
      >
        <div>
          <Input placeholder="Email" {...register("email")} />
          {errors.email && (
            <p className="text-red-500 text-sm">{errors.email.message}</p>
          )}
        </div>
        <div>
          <Input placeholder="Password" {...register("password")} />
          {errors.password && (
            <p className="text-red-500 text-sm">{errors.password.message}</p>
          )}
        </div>
        <div className="flex justify-between">
          <Button type="submit">Login</Button>
          <div className="flex items-center gap-2">
            <p className="text-sm">
              Don't have an account?{" "}
              <NavLink
                to="/register"
                className="text-rose-400 underline hover:text-rose-500 transition"
              >
                Register
              </NavLink>
            </p>
          </div>
        </div>
      </form>
    </div>
  );
};

export default Login;
