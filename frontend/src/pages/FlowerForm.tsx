import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  useCreateFlower,
  useFlower,
  useUpdateFlower,
} from "@/hooks/api/flowers";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router";
import { toast } from "sonner";
import z from "zod";
import { Flower, Camera, FileText, ArrowLeft, Loader2 } from "lucide-react";
import type { FlowerPayloadType } from "@/types/flower";

const schema = z.object({
  title: z.string().min(1, "Title must be at least 1 character"),
  content: z.string().min(5, "Content must be at least 5 characters"),
  image_url: z.string().url("Please enter a valid image URL"),
});

const FlowerForm = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const createFlower = useCreateFlower();
  const updateFlower = useUpdateFlower(id || "");
  const { data, isLoading } = useFlower(id!);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FlowerPayloadType>({
    resolver: zodResolver(schema),
    defaultValues: data || { title: "", content: "", image_url: "" },
  });

  const onSubmit = (val: FlowerPayloadType) => {
    if (id) {
      updateFlower.mutate(val, {
        onSuccess: () => {
          toast.success("Flower updated successfully");
          navigate(`/flowers`);
        },
        onError: () => {
          toast.error("Failed to update flower");
        },
      });
    } else {
      createFlower.mutate(val, {
        onSuccess: () => {
          toast.success("Flower created successfully");
          navigate(`/flowers`);
        },
        onError: () => {
          toast.error("Failed to create flower");
        },
      });
    }
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 flex items-center justify-center">
        <div className="flex flex-col items-center space-y-4">
          <div className="w-12 h-12 border-4 border-rose-200 border-t-rose-500 rounded-full animate-spin"></div>
          <p className="text-rose-600 font-medium">Loading flower details...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 py-12">
      <div className="max-w-2xl mx-auto px-4">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center gap-2 bg-rose-100 text-rose-600 px-4 py-2 rounded-full text-sm font-medium mb-4">
            <Flower className="w-4 h-4" />
            {id ? "Edit your beautiful flower" : "Share a new flower"}
          </div>
          <h1 className="text-4xl font-bold bg-gradient-to-r from-rose-600 via-pink-600 to-purple-600 bg-clip-text text-transparent mb-2">
            {id ? "Edit Flower" : "Create New Flower"}
          </h1>
          <p className="text-gray-600 text-lg">
            {id
              ? "Update your flower details"
              : "Share the beauty of your flower with our community"}
          </p>
        </div>

        {/* Form Card */}
        <Card className="bg-white/80 backdrop-blur-sm border-0 shadow-xl">
          <CardHeader>
            <CardTitle className="text-2xl text-center text-gray-900">
              Flower Details
            </CardTitle>
            <CardDescription className="text-center text-gray-600">
              Fill in the details about your beautiful flower
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
              {/* Title Field */}
              <div className="space-y-2">
                <Label
                  htmlFor="title"
                  className="text-sm font-medium text-gray-700 flex items-center gap-2"
                >
                  <FileText className="w-4 h-4" />
                  Flower Title
                </Label>
                <Input
                  id="title"
                  placeholder="Enter a beautiful title for your flower"
                  {...register("title")}
                  className="h-12 text-lg"
                />
                {errors.title && (
                  <p className="text-red-500 text-sm flex items-center gap-1">
                    <span className="w-1 h-1 bg-red-500 rounded-full"></span>
                    {errors.title.message}
                  </p>
                )}
              </div>

              {/* Content Field */}
              <div className="space-y-2">
                <Label
                  htmlFor="content"
                  className="text-sm font-medium text-gray-700 flex items-center gap-2"
                >
                  <FileText className="w-4 h-4" />
                  Description
                </Label>
                <Textarea
                  id="content"
                  placeholder="Describe your flower in detail... What makes it special? Where did you find it? What colors and features does it have?"
                  {...register("content")}
                  className="min-h-32 text-lg resize-none"
                />
                {errors.content && (
                  <p className="text-red-500 text-sm flex items-center gap-1">
                    <span className="w-1 h-1 bg-red-500 rounded-full"></span>
                    {errors.content.message}
                  </p>
                )}
              </div>

              {/* Image URL Field */}
              <div className="space-y-2">
                <Label
                  htmlFor="image_url"
                  className="text-sm font-medium text-gray-700 flex items-center gap-2"
                >
                  <Camera className="w-4 h-4" />
                  Image URL
                </Label>
                <Input
                  id="image_url"
                  placeholder="https://example.com/your-flower-image.jpg"
                  {...register("image_url")}
                  className="h-12 text-lg"
                />
                {errors.image_url && (
                  <p className="text-red-500 text-sm flex items-center gap-1">
                    <span className="w-1 h-1 bg-red-500 rounded-full"></span>
                    {errors.image_url.message}
                  </p>
                )}
              </div>

              {/* Action Buttons */}
              <div className="flex flex-col sm:flex-row gap-4 pt-4">
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => navigate("/flowers")}
                  className="flex-1 h-12 text-lg border-rose-200 text-rose-600 hover:bg-rose-50"
                >
                  <ArrowLeft className="w-4 h-4 mr-2" />
                  Cancel
                </Button>
                <Button
                  type="submit"
                  disabled={isSubmitting}
                  className="flex-1 h-12 text-lg bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white shadow-lg hover:shadow-xl transition-all duration-300"
                >
                  {isSubmitting ? (
                    <>
                      <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                      {id ? "Updating..." : "Creating..."}
                    </>
                  ) : (
                    <>
                      <Flower className="w-4 h-4 mr-2" />
                      {id ? "Update Flower" : "Create Flower"}
                    </>
                  )}
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default FlowerForm;
