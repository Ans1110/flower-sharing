"use client";

import { FlowerPayloadType, FlowerType } from "@/types/flower";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import z from "zod";
import { ControlledInput } from "./ui/controlled-input";
import Image from "next/image";
import { Label } from "./ui/label";
import { Button } from "./ui/button";
import { Camera, ImageIcon, Loader2 } from "lucide-react";

const postSchema = z.object({
  title: z.string().min(1, "Title is required"),
  content: z.string().min(10, "Content must be at least 10 characters"),
  image: z.union([z.instanceof(File), z.string(), z.null()]).optional(),
});

type PostFormData = z.infer<typeof postSchema>;

type EditPostFormProps = Readonly<{
  post: FlowerType;
  onSubmit: (data: FlowerPayloadType) => void;
  onCancel: () => void;
  isLoading?: boolean;
}>;

const EditPostForm = ({
  post,
  onSubmit,
  onCancel,
  isLoading = false,
}: EditPostFormProps) => {
  const methods = useForm<PostFormData>({
    resolver: zodResolver(postSchema),
    defaultValues: {
      title: post.title,
      content: post.content,
      image: null,
    },
  });

  const {
    handleSubmit,
    formState: { isDirty },
    watch,
    setValue,
  } = methods;

  const imageFileRef = useRef<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [imagePreview, setImagePreview] = useState<string | null>(null);

  const imageFile = watch("image");

  useEffect(() => {
    if (typeof imageFile === "string" && imageFile.startsWith("data:")) {
      setImagePreview(imageFile);
    } else if (!imageFile) {
      setImagePreview(null);
      imageFileRef.current = null;
    }
  }, [imageFile]);

  const onSubmitForm = (data: PostFormData) => {
    const payload: FlowerPayloadType = {
      title: data.title,
      content: data.content,
      imageUrl: imageFileRef.current || null,
      authorId: post.author.id.toString(),
    };
    onSubmit(payload);
  };

  const handleImageClick = () => {
    fileInputRef.current?.click();
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      imageFileRef.current = file;
      // Create preview URL
      const reader = new FileReader();
      reader.onloadend = () => {
        setValue("image", reader.result as string, { shouldDirty: true });
      };
      reader.readAsDataURL(file);
    }
  };

  const displayImage = imagePreview || post.image_url;

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit(onSubmitForm)} className="space-y-5">
        {/* Image Section */}
        <div className="space-y-2">
          <Label className="text-sm font-medium">
            {imagePreview ? "New Image" : "Current Image"}
          </Label>
          <div
            onClick={handleImageClick}
            className="relative w-full aspect-video rounded-xl overflow-hidden border-2 border-dashed border-rose-200 dark:border-rose-800 hover:border-rose-400 dark:hover:border-rose-600 transition-colors cursor-pointer group bg-linear-to-br from-rose-50 to-pink-50 dark:from-rose-950/30 dark:to-pink-950/30"
          >
            {displayImage ? (
              <>
                <Image
                  src={displayImage}
                  alt={imagePreview ? "New image preview" : "Current image"}
                  fill
                  className="object-cover transition-transform duration-300 group-hover:scale-105"
                  sizes="(max-width: 768px) 100vw, 500px"
                />
                {/* Overlay on hover */}
                <div className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                  <div className="flex flex-col items-center text-white">
                    <Camera className="size-8 mb-2" />
                    <span className="text-sm font-medium">Change Image</span>
                  </div>
                </div>
              </>
            ) : (
              <div className="absolute inset-0 flex flex-col items-center justify-center text-rose-400 dark:text-rose-500">
                <ImageIcon className="size-12 mb-2" />
                <span className="text-sm font-medium">
                  Click to upload image
                </span>
              </div>
            )}
          </div>
          {/* Hidden file input */}
          <ControlledInput<PostFormData>
            name="image"
            ref={fileInputRef}
            type="file"
            accept="image/*"
            disabled={isLoading}
            className="hidden"
            onChange={handleFileChange}
          />
          <p className="text-xs text-muted-foreground">
            Click on the image to upload a new one
          </p>
        </div>

        {/* Title and Content */}
        <div className="space-y-4">
          <ControlledInput<PostFormData>
            name="title"
            label="Title"
            placeholder="Enter a beautiful title for your flower"
            disabled={isLoading}
          />

          <ControlledInput<PostFormData>
            name="content"
            label="Description"
            placeholder="Describe this flower..."
            disabled={isLoading}
          />
        </div>

        {/* Action Buttons */}
        <div className="flex justify-end gap-3 pt-2">
          <Button
            type="button"
            variant="outline"
            onClick={onCancel}
            disabled={isLoading}
            className="border-rose-200 hover:bg-rose-50 dark:border-rose-800 dark:hover:bg-rose-900/30"
          >
            Cancel
          </Button>
          <Button
            type="submit"
            disabled={(!isDirty && !imageFileRef.current) || isLoading}
            className="bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white"
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
};

export default EditPostForm;
export { EditPostForm };
