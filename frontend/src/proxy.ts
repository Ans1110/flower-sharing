import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const adminRoutes = ["/admin"];

export function proxy(request: NextRequest) {
  const { pathname } = request.nextUrl;

  const isAdminRoute = adminRoutes.some((route) => pathname.startsWith(route));

  // Skip check for non-admin routes
  if (!isAdminRoute) {
    return NextResponse.next();
  }

  // Check role cookie (set by frontend auth store)
  const userRole = request.cookies.get("role")?.value;

  if (userRole === "admin") {
    return NextResponse.next();
  }

  // Redirect non-admin users to home
  return NextResponse.redirect(new URL("/", request.url));
}

// Configure which routes use this proxy
export const config = {
  matcher: ["/admin/:path*"],
};
