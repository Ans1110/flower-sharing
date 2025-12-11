"use client";

import { Button } from "@/components/ui/button";
import CustomCard from "@/components/ui/CustomCard";
import { useAuthStore } from "@/store/auth";
import { Camera, Flower, Heart, Sparkles, Users } from "lucide-react";
import Link from "next/link";

export default function Home() {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);

  return (
    <div className="min-h-screen bg-linear-to-br from-rose-50 via-pink-50 to-violet-50 dark:from-neutral-950 dark:via-neutral-900 dark:to-neutral-950">
      {/* Hero Section */}
      <section className="relative overflow-hidden py-20 px-4">
        <div className="absolute inset-0 bg-linear-to-r from-rose-100/20 to-pink-100/20 dark:from-neutral-900/20 dark:to-neutral-900/20 blur"></div>
        <div className="relative max-w-6xl mx-auto text-center">
          <div className="mb-8">
            <div className="inline-flex items-center gap-2 bg-rose-100 dark:bg-rose-900 text-rose-600 dark:text-rose-200 px-4 py-2 rounded-full text-sm font-medium mb-6">
              <Sparkles className="size-4" />
              Share the beauty of flowers
            </div>
            <h1 className="text-5xl md:text-7xl font-bold bg-linear-to-r from-rose-600 via-pink-600 dark:to-violet-300 bg-clip-text text-transparent mb-6 leading-tight pb-2">
              Flower Sharing
            </h1>
            <p className="text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto mb-8 leading-relaxed">
              Discover, share, and celebrate the beauty of flowers. Connect with
              fellow flower enthusiasts and create a blooming community.
            </p>
          </div>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button
              asChild
              size="lg"
              className="bg-linear-to-r from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600 text-white px-8 py-4 text-lg"
            >
              <Link href="/flowers">
                <Flower className="size-4" />
                Explore Flowers
              </Link>
            </Button>
            {!isAuthenticated && (
              <Button
                asChild
                size="lg"
                variant="outline"
                className="border-rose-300 text-rose-600 hover:bg-rose-50 hover:text-rose-700 dark:border-rose-800 dark:text-rose-400 dark:hover:bg-rose-900/50 dark:hover:text-rose-300 px-8 py-4 text-lg"
              >
                <Link href="/register">Join Community</Link>
              </Button>
            )}
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="px-4 pt-15 pb-25">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 dark:text-gray-100 mb-4">
              Why Choose Flower Sharing?
            </h2>
            <p className="text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
              Experience the joy of sharing and discovering beautiful flowers
              with our community
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            <CustomCard
              title="Share Photos"
              description="Upload and share beautiful flower photos with detailed descriptions"
              icon={Camera}
              iconClassName="text-rose-600 dark:text-rose-400"
            />
            <CustomCard
              title="Connect & Love"
              description="Like and comment on beautiful flowers shared by the community"
              icon={Heart}
              iconClassName="text-pink-600 dark:text-pink-400"
            />
            <CustomCard
              title="Join Community"
              description="Be part of a passionate community of flower lovers and enthusiasts"
              icon={Users}
              iconClassName="text-purple-600 dark:text-purple-400"
            />
          </div>
        </div>
      </section>

      {/* CTA Section */}
      {!isAuthenticated && (
        <section className="py-20 px-4 bg-linear-to-r from-rose-500 to-pink-500 dark:from-neutral-950 dark:to-neutral-900">
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
                className="bg-white dark:bg-white/10 text-rose-600 dark:text-white hover:bg-gray-100 dark:hover:bg-white/20 hover:text-rose-700 dark:hover:text-rose-300 px-8 py-4 text-lg"
              >
                <Link href="/register">Create Account</Link>
              </Button>
              <Button
                asChild
                size="lg"
                variant="outline"
                className="border-white text-rose-600 hover:bg-white/10 px-8 py-4 text-lg"
              >
                <Link href="/login">Sign In</Link>
              </Button>
            </div>
          </div>
        </section>
      )}
    </div>
  );
}
