"use client";

import { useCreatePost } from "@/hooks/api/flowers";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, useForm } from "react-hook-form";
import z from "zod";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import {
  ArrowLeft,
  Camera,
  FileText,
  Flower,
  Loader2,
  Share,
} from "lucide-react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ControlledInput } from "@/components/ui/controlled-input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import Image from "next/image";

const flowerSchema = z.object({
  title: z.string().min(1, "Title is required"),
  content: z.string().min(10, "Content must be at least 10 characters"),
  image: z.string().min(1, "Image is required"),
});

type FlowerFormData = z.infer<typeof flowerSchema>;

export default function NewFlower() {
  const router = useRouter();
  const createFlower = useCreatePost();
  const [imagePreview, setImagePreview] = useState<string>("");
  const form = useForm<FlowerFormData>({
    resolver: zodResolver(flowerSchema),
    defaultValues: {
      title: "",
      content: "",
      image: "",
    },
  });

  const {
    handleSubmit,
    register,
    formState: { isSubmitting },
  } = form;

  const onSubmit = async (data: FlowerFormData) => {
    // Convert base64 to File
    const base64Response = await fetch(data.image);
    const blob = await base64Response.blob();
    const file = new File([blob], "flower-image.jpg", { type: blob.type });

    // Create FormData
    const formData = new FormData();
    formData.append("title", data.title);
    formData.append("content", data.content);
    formData.append("image", file);

    createFlower.mutate(formData, {
      onSuccess: () => {
        toast.success("Flower created successfully");
        router.push("/flowers");
      },
      onError: (error) => {
        toast.error(
          error.response?.data?.error || "An unexpected error occurred"
        );
      },
    });
  };

  if (isSubmitting) {
    return (
      <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 flex items-center justify-center py-12 px-4">
        <div className="flex flex-col items-center space-y-4">
          <div className="w-12 h-12 border-4 border-rose-200 border-t-rose-500 rounded-full animate-spin"></div>
          <p className="text-rose-600 font-medium">Loading flower details...</p>
        </div>
      </div>
    );
  }
  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950 py-12">
      <div className="max-w-2xl mx-auto px-4">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center gap-2 bg-rose-100 text-rose-600 px-4 py-2 rounded-full text-sm font-medium mb-4">
            <Flower className="size-4" />
            Share a New Flower
          </div>
          <h1 className="text-4xl font-bold bg-linear-to-r from-rose-600 via-pink-600 to-violet-600 bg-clip-text text-transparent mb-2">
            Share a New Flower
          </h1>
          <p className="text-gray-600 dark:text-gray-400 text-lg">
            Share your beautiful moments with our community
          </p>
        </div>

        {/* Flower Form */}
        <Card className="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-sm border-0 shadow-xl">
          <CardHeader>
            <CardTitle className="text-2xl text-center text-gray-900 dark:text-gray-100">
              Flower Details
            </CardTitle>
            <CardDescription className="text-center text-gray-600 dark:text-gray-400">
              Fill in the details about your flower
            </CardDescription>
          </CardHeader>

          <CardContent>
            <FormProvider {...form}>
              <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                {/* Title Field */}
                <div className="space-y-2">
                  <label
                    htmlFor="title"
                    className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                  >
                    <FileText className="size-4" />
                    Flower Title
                  </label>
                  <ControlledInput
                    name="title"
                    placeholder="Enter the title of your flower"
                    className="h-12 text-lg"
                    required
                  />
                </div>

                {/* Content Field */}
                <div className="space-y-2">
                  <label
                    htmlFor="content"
                    className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                  >
                    <FileText className="size-4" />
                    Description
                  </label>
                  <Textarea
                    id="content"
                    placeholder="Describe your flower in detail... What makes it special? Where did you find it? What colors and features does it have?"
                    className="min-h-32 text-lg resize-none"
                    {...register("content")}
                    required
                  />
                </div>

                {/* Image Field */}
                <div className="space-y-2">
                  <label
                    htmlFor="image"
                    className="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2"
                  >
                    <Camera className="size-4" />
                    Upload Image
                  </label>
                  <ControlledInput
                    name="image"
                    placeholder="Upload an image of your flower"
                    type="file"
                    className="h-12 text-lg"
                    required
                    lang="en"
                    accept="image/*"
                    onChange={(e) => {
                      const file = e.target.files?.[0];
                      if (file) {
                        const reader = new FileReader();
                        reader.onloadend = () => {
                          setImagePreview(reader.result as string);
                        };
                        reader.readAsDataURL(file);
                      }
                    }}
                  />
                  {imagePreview && (
                    <div className="mt-4 relative w-full h-64 rounded-lg overflow-hidden border-2 border-rose-200 dark:border-rose-800">
                      <Image
                        src={imagePreview}
                        alt="Preview"
                        sizes="(max-width: 768px) 100vw, 500px"
                        fill
                        className="object-cover"
                      />
                    </div>
                  )}
                </div>

                {/* Action Buttons */}
                <div className="flex flex-col sm:flex-row gap-4 pt-4">
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => router.push("/flowers")}
                    className="flex-1 h-12 text-lg border-rose-200 text-rose-600 hover:bg-rose-50 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300"
                  >
                    <ArrowLeft className="size-4 mr-2" />
                    Cancel
                  </Button>
                  <Button
                    type="submit"
                    disabled={isSubmitting}
                    className="flex-1 h-12 bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg"
                  >
                    {isSubmitting ? (
                      <>
                        <Loader2 className="size-4 mr-2 animate-spin" />
                        Saving...
                      </>
                    ) : (
                      <>
                        <Share className="size-4 mr-2" />
                        Share Flower
                      </>
                    )}
                  </Button>
                </div>
              </form>
            </FormProvider>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
