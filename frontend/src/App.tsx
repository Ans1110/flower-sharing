import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider } from "react-router";
import { router } from "./routers";
import { Toaster } from "@/components/ui/sonner";
import AuthProvider from "@/components/AuthProvider";

const queryClient = new QueryClient();

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <RouterProvider router={router} />
        <Toaster />
      </AuthProvider>
    </QueryClientProvider>
  );
};

export default App;
