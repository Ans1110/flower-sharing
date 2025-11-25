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

  const userRole = request.cookies.get("role")?.value;

  if (isProtectedRoute && userRole !== "admin") {
    const url = new URL("/error", request.url);
    url.searchParams.set("redirect", pathname);
    return NextResponse.redirect(url);
  }

  if (isAdminRoute && userRole !== "admin") {
    return NextResponse.redirect(new URL("/error", request.url));
  }

  return NextResponse.next();
}

// Configure which routes use this middleware
export const config = {
  matcher: ["/flowers/new", "/flowers/:id/edit", "/profile", "/admin/:path*"],
};
