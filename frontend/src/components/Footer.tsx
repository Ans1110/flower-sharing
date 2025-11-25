import { Flower } from "lucide-react";

const Footer = () => {
  return (
    <footer className="bg-white/80 dark:bg-black/80 border-t border-rose-100 dark:border-rose-900 backdrop-blur-sm py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <div className="flex items-center justify-center space-x-2 mb-4">
            <div className="w-8 h-8 bg-linear-to-r from-rose-500 to-pink-500 rounded-full flex items-center justify-center ">
              <Flower className="w-5 h-5 text-white" />
            </div>
            <span className="text-lg font-bold bg-linear-to-r from-rose-600 to-pink-600 bg-clip-text text-transparent">
              FlowerShare
            </span>
          </div>
          <p className="text-gray-600 dark:text-gray-400 text-sm">
            &copy; {new Date().getFullYear()} FlowerShare. All rights reserved.
          </p>
          <p className="text-gray-500 dark:text-gray-400 text-xs mt-2">
            Share the beauty of flowers with our community
          </p>
        </div>
      </div>
    </footer>
  );
};

export { Footer };
