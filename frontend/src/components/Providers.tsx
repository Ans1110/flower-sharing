"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { toast, Toaster } from "sonner";
import { ThemeProvider } from "next-themes";
import AlertDialogProvider from "./ui/alert-dialog-prodiver";

type ProvidersProps = {
  children: React.ReactNode;
};

export default function Providers({ children }: ProvidersProps) {
  const queryClient = new QueryClient({
    defaultOptions: {
      mutations: {
        onError: (e) => {
          if (e.message === "NEXT_REDIRECT") return;
          toast.error(e.message);
        },
        onSuccess: () => {
          toast.success("Operation successful");
        },
      },
    },
  });

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider
        attribute="class"
        defaultTheme="light"
        enableSystem
        disableTransitionOnChange
      >
        <Toaster />
        <AlertDialogProvider />
        {children}
      </ThemeProvider>
    </QueryClientProvider>
  );
}

export { Providers };
