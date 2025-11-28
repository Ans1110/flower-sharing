"use client";

import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Button } from "./ui/button";
import { ControlledInput } from "./ui/controlled-input";
import { UserAdminResponseType, UserPayloadType } from "@/types/user";
import { Loader2 } from "lucide-react";

const userSchema = z.object({
  username: z.string().min(3, "Username must be at least 3 characters"),
  email: z.string().email("Invalid email address"),
  avatar: z.string().url("Invalid avatar URL").or(z.literal("")),
});

type UserFormData = z.infer<typeof userSchema>;

type EditUserFormProps = Readonly<{
  user: UserAdminResponseType;
  onSubmit: (data: UserPayloadType) => void;
  onCancel: () => void;
  isLoading?: boolean;
}>;

export default function EditUserForm({
  user,
  onSubmit,
  onCancel,
  isLoading = false,
}: EditUserFormProps) {
  const methods = useForm<UserFormData>({
    resolver: zodResolver(userSchema),
    defaultValues: {
      username: user.username,
      email: user.email,
      avatar: user.avatar,
    },
  });

  const {
    handleSubmit,
    formState: { isDirty },
  } = methods;

  const onSubmitForm = (data: UserFormData) => {
    onSubmit(data);
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmitForm)} className="space-y-6">
        <div className="space-y-4">
          <ControlledInput<UserFormData>
            name="username"
            label="Username"
            placeholder="Enter username"
            disabled={isLoading}
          />

          <ControlledInput<UserFormData>
            name="email"
            type="email"
            label="Email"
            placeholder="Enter email"
            disabled={isLoading}
          />

          <ControlledInput<UserFormData>
            name="avatar"
            label="Avatar URL"
            placeholder="Enter avatar URL"
            disabled={isLoading}
          />
        </div>

        <div className="flex justify-end gap-3">
          <Button
            type="button"
            variant="outline"
            onClick={onCancel}
            disabled={isLoading}
          >
            Cancel
          </Button>
          <Button type="submit" disabled={!isDirty || isLoading}>
            {isLoading ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Saving...
              </>
            ) : (
              "Save Changes"
            )}
          </Button>
        </div>
      </form>
    </FormProvider>
  );
}
