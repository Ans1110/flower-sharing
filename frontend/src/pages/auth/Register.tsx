import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { useForm } from "react-hook-form";
import type { RegisterPayloadType } from "@/types/auth";
import { useMutation } from "@tanstack/react-query";
import api from "@/services/api";
import { toast } from "sonner";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { NavLink, useNavigate } from "react-router-dom";

const schema = z.object({
  username: z.string().min(3, "Username must be at least 3 characters"),
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
});

const Register = () => {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterPayloadType>({
    resolver: zodResolver(schema),
  });

  const mutation = useMutation({
    mutationFn: (data: RegisterPayloadType) =>
      api.post("/register", data).then((res) => res.data),
    onSuccess: () => {
      toast.success("Register successful");
      navigate("/login");
    },
    onError: () => {
      toast.error("Register failed");
    },
  });

  return (
    <div className="max-w-md mx-auto mt-32 p-6 border rounded-lg shadow">
      <h1 className="text-2xl font-bold mb-4 text-rose-400">Register</h1>
      <form
        onSubmit={handleSubmit((data) => mutation.mutate(data))}
        className="space-y-4"
      >
        <div>
          <Input placeholder="Username" {...register("username")} />
          {errors.username && (
            <p className="text-red-500 text-sm">{errors.username.message}</p>
          )}
        </div>
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
          <Button type="submit">Register</Button>
          <div className="flex items-center gap-2">
            <p className="text-sm">
              Already have an account?{" "}
              <NavLink
                to="/login"
                className="text-rose-400 underline hover:text-rose-500 transition"
              >
                Login
              </NavLink>
            </p>
          </div>
        </div>
      </form>
    </div>
  );
};

export default Register;
