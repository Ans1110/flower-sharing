import { Button } from "@/components/ui/button";
import { Link } from "react-router";
import { Home, Search, Flower } from "lucide-react";

const NotFound = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-50 via-pink-50 to-purple-50 flex items-center justify-center py-12 px-4">
      <div className="max-w-2xl mx-auto text-center">
        {/* 404 Number */}
        <div className="mb-8">
          <h1 className="text-9xl md:text-[12rem] font-bold bg-gradient-to-r from-rose-600 via-pink-600 to-purple-600 bg-clip-text text-transparent leading-none">
            404
          </h1>
        </div>

        {/* Error Message */}
        <div className="mb-8">
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Oops! Page Not Found
          </h2>
          <p className="text-xl text-gray-600 leading-relaxed">
            The page you're looking for seems to have wilted away. Don't worry,
            there are plenty of beautiful flowers to discover!
          </p>
        </div>

        {/* Decorative Element */}
        <div className="mb-8">
          <div className="w-32 h-32 bg-gradient-to-br from-rose-100 to-pink-100 rounded-full flex items-center justify-center mx-auto mb-6">
            <Flower className="w-16 h-16 text-rose-400" />
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
          <Button
            asChild
            size="lg"
            className="bg-gradient-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg shadow-lg hover:shadow-xl transition-all duration-300"
          >
            <Link to="/">
              <Home className="w-5 h-5 mr-2" />
              Go Home
            </Link>
          </Button>

          <Button
            asChild
            size="lg"
            variant="outline"
            className="border-rose-200 text-rose-600 hover:bg-rose-50 px-8 py-4 text-lg transition-all duration-300"
          >
            <Link to="/flowers">
              <Search className="w-5 h-5 mr-2" />
              Browse Flowers
            </Link>
          </Button>
        </div>
      </div>
    </div>
  );
};

export default NotFound;
