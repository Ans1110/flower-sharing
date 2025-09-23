import { Button } from "@/components/ui/button";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Link } from "react-router";
import { useAuthStore } from "@/store/auth";
import { Flower, Heart, Users, Camera, Sparkles } from "lucide-react";

const Home = () => {
  const token = useAuthStore((state) => state.token);

  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50">
      {/* Hero Section */}
      <section className="relative overflow-hidden py-20 px-4">
        <div className="absolute inset-0 bg-gradient-to-r from-rose-100/20 to-pink-100/20"></div>
        <div className="relative max-w-6xl mx-auto text-center">
          <div className="mb-8">
            <div className="inline-flex items-center gap-2 bg-rose-100 text-rose-600 px-4 py-2 rounded-full text-sm font-medium mb-6">
              <Sparkles className="w-4 h-4" />
              Share the beauty of flowers
            </div>
            <h1 className="text-5xl md:text-7xl font-bold bg-gradient-to-r from-rose-600 via-pink-600 to-purple-600 bg-clip-text text-transparent mb-6">
              Flower Sharing
            </h1>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto mb-8 leading-relaxed">
              Discover, share, and celebrate the beauty of flowers. Connect with
              fellow flower enthusiasts and create a blooming community.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button
              asChild
              size="lg"
              className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg"
            >
              <Link to="/flowers">
                <Flower className="w-5 h-5 mr-2" />
                Explore Flowers
              </Link>
            </Button>
            {!token && (
              <Button
                asChild
                variant="outline"
                size="lg"
                className="border-rose-300 text-rose-600 hover:bg-rose-50 px-8 py-4 text-lg"
              >
                <Link to="/register">Join Community</Link>
              </Button>
            )}
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 px-4">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Why Choose Flower Sharing?
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Experience the joy of sharing and discovering beautiful flowers
              with our community
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <Card className="border-0 shadow-lg hover:shadow-xl transition-all duration-300 bg-white/80 backdrop-blur-sm">
              <CardHeader className="text-center">
                <div className="w-16 h-16 bg-gradient-to-br from-rose-100 to-pink-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Camera className="w-8 h-8 text-rose-600" />
                </div>
                <CardTitle className="text-2xl text-gray-900">
                  Share Photos
                </CardTitle>
                <CardDescription className="text-gray-600 text-lg">
                  Upload and share beautiful flower photos with detailed
                  descriptions
                </CardDescription>
              </CardHeader>
            </Card>

            <Card className="border-0 shadow-lg hover:shadow-xl transition-all duration-300 bg-white/80 backdrop-blur-sm">
              <CardHeader className="text-center">
                <div className="w-16 h-16 bg-gradient-to-br from-pink-100 to-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Heart className="w-8 h-8 text-pink-600" />
                </div>
                <CardTitle className="text-2xl text-gray-900">
                  Connect & Love
                </CardTitle>
                <CardDescription className="text-gray-600 text-lg">
                  Like and comment on beautiful flowers shared by the community
                </CardDescription>
              </CardHeader>
            </Card>

            <Card className="border-0 shadow-lg hover:shadow-xl transition-all duration-300 bg-white/80 backdrop-blur-sm">
              <CardHeader className="text-center">
                <div className="w-16 h-16 bg-gradient-to-br from-purple-100 to-indigo-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Users className="w-8 h-8 text-purple-600" />
                </div>
                <CardTitle className="text-2xl text-gray-900">
                  Join Community
                </CardTitle>
                <CardDescription className="text-gray-600 text-lg">
                  Be part of a passionate community of flower lovers and
                  enthusiasts
                </CardDescription>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      {!token && (
        <section className="py-20 px-4 bg-gradient-to-r from-rose-500 to-pink-500">
          <div className="max-w-4xl mx-auto text-center text-white">
            <h2 className="text-4xl font-bold mb-4">Ready to Start Sharing?</h2>
            <p className="text-xl mb-8 opacity-90">
              Join thousands of flower enthusiasts and start sharing your
              beautiful moments today
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Button
                asChild
                size="lg"
                variant="secondary"
                className="bg-white text-rose-600 hover:bg-gray-100 px-8 py-4 text-lg"
              >
                <Link to="/register">Create Account</Link>
              </Button>
              <Button
                asChild
                size="lg"
                variant="outline"
                className="border-white text-rose-600 hover:bg-white/10 px-8 py-4 text-lg"
              >
                <Link to="/login">Sign In</Link>
              </Button>
            </div>
          </div>
        </section>
      )}
    </div>
  );
};

export default Home;
