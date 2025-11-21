import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const protectedRoutes = ["/flowers/new", "/profile", "/admin"];

const adminRoutes = ["/admin"];

export function proxy(request: NextRequest) {
  const { pathname } = request.nextUrl;

  const isProtectedRoute = protectedRoutes.some((route) =>
    pathname.startsWith(route)
  );

  const isAdminRoute = adminRoutes.some((route) => pathname.startsWith(route));

  const token = request.cookies.get("accessToken")?.value;
  const userRole = request.cookies.get("role")?.value;

  if (isProtectedRoute && !token) {
    const url = new URL("/login", request.url);
    url.searchParams.set("redirect", pathname);
    return NextResponse.redirect(url);
  }

  if (isAdminRoute && userRole !== "admin") {
    return NextResponse.redirect(new URL("/", request.url));
  }

  return NextResponse.next();
}

// Configure which routes use this middleware
export const config = {
  matcher: ["/flowers/new", "/flowers/:id/edit", "/profile", "/admin/:path*"],
};
