import RootLayout from "@/layout/RootLayout";
import Login from "@/pages/auth/Login";
import Flowers from "@/pages/Flowers";
import FlowersDetail from "@/pages/Flowers-detail";
import Home from "@/pages/Home";
import Profile from "@/pages/Profile";
import Register from "@/pages/auth/Register";
import { createBrowserRouter } from "react-router-dom";
import users from "@/pages/admin/Users";
import posts from "@/pages/admin/Posts";
import NotFound from "@/pages/NotFound";
import Error from "@/pages/Error";
import FlowerForm from "@/pages/FlowerForm";
import ProtectedRoute from "@/components/ProtectedRoute";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: RootLayout,
    children: [
      {
        path: "/",
        Component: Home,
      },
      {
        path: "/flowers",
        Component: Flowers,
      },
      {
        path: "/flowers/:id",
        Component: FlowersDetail,
      },
      {
        path: "/login",
        Component: Login,
      },
      {
        path: "/register",
        Component: Register,
      },
      // Protected routes
      {
        path: "/",
        Component: ProtectedRoute,
        children: [
          {
            path: "/flowers/new",
            Component: FlowerForm,
          },
          {
            path: "/flowers/:id/edit",
            Component: FlowerForm,
          },
          {
            path: "/profile",
            Component: Profile,
          },
          {
            path: "/admin",
            children: [
              {
                path: "/admin/posts",
                Component: posts,
              },
              {
                path: "/admin/users",
                Component: users,
              },
            ],
          },
        ],
      },
    ],
  },
  {
    path: "*",
    Component: NotFound,
    ErrorBoundary: Error,
  },
]);
