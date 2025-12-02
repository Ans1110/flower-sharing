"use client";

import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Button } from "./ui/button";
import { ControlledInput } from "./ui/controlled-input";
import { Label } from "./ui/label";
import { UserAdminResponseType, UserPayloadType } from "@/types/user";
import { Loader2 } from "lucide-react";
import Image from "next/image";
import { useEffect, useRef, useState } from "react";

const userSchema = z.object({
  username: z.string().min(3, "Username must be at least 3 characters"),
  email: z.string().email("Invalid email address"),
  avatar: z.union([z.instanceof(File), z.string(), z.null()]).optional(),
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
      avatar: null,
    },
  });

  const {
    handleSubmit,
    formState: { isDirty },
    watch,
  } = methods;

  // Store the actual File object separately since ControlledInput converts it to data URL
  const avatarFileRef = useRef<File | null>(null);
  const [avatarPreview, setAvatarPreview] = useState<string | null>(null);

  // Watch avatar field to detect file changes (this will be a data URL string)
  const avatarFile = watch("avatar");

  // Update preview when avatarFile changes (data URL from ControlledInput)
  useEffect(() => {
    if (typeof avatarFile === "string" && avatarFile.startsWith("data:")) {
      setAvatarPreview(avatarFile);
    } else if (!avatarFile) {
      setAvatarPreview(null);
      avatarFileRef.current = null;
    }
  }, [avatarFile]);

  const onSubmitForm = (data: UserFormData) => {
    const payload: UserPayloadType = {
      username: data.username,
      email: data.email,
      avatar: avatarFileRef.current || null,
    };
    onSubmit(payload);
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

          <div className="space-y-2">
            {/* Show preview of selected file or current avatar */}
            {(avatarPreview || user.avatar) && (
              <div className="mb-2">
                <Label className="text-xs text-muted-foreground mb-1 block">
                  {avatarPreview ? "New avatar preview:" : "Current avatar:"}
                </Label>
                <div className="relative h-16 w-16 rounded-full border overflow-hidden">
                  {avatarPreview ? (
                    <Image
                      src={avatarPreview}
                      alt="Avatar preview"
                      fill
                      className="object-cover"
                    />
                  ) : user.avatar ? (
                    <Image
                      src={user.avatar}
                      alt="Current avatar"
                      fill
                      className="object-cover"
                    />
                  ) : null}
                </div>
              </div>
            )}
            <ControlledInput<UserFormData>
              name="avatar"
              label="Upload new avatar (optional)"
              disabled={isLoading}
              type="file"
              accept="image/*"
              onChange={(e) => {
                const file = e.target.files?.[0];
                if (file) {
                  avatarFileRef.current = file;
                } else {
                  avatarFileRef.current = null;
                }
              }}
            />
          </div>
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
          <Button
            type="submit"
            disabled={(!isDirty && !avatarFileRef.current) || isLoading}
          >
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
