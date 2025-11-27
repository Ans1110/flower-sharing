"use client";

import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";
import { useEffect, useState } from "react";
import { Switch } from "@/components/ui/switch";

const ThemeToggle = () => {
  const { setTheme, resolvedTheme } = useTheme();
  const [mounted, setMounted] = useState(false);

  // Only run on client after mount
  // eslint-disable-next-line react-hooks/rules-of-hooks
  useEffect(() => {
    setMounted(true);
  }, []);

  const isDark = resolvedTheme === "dark";

  return (
    <div className="relative inline-flex items-center" suppressHydrationWarning>
      {mounted ? (
        <>
          <Switch
            checked={isDark}
            onCheckedChange={(checked: boolean) =>
              setTheme(checked ? "dark" : "light")
            }
            className="data-[state=checked]:bg-gray-700 dark:data-[state=checked]:bg-gray-600 data-[state=unchecked]:bg-gray-300 dark:data-[state=unchecked]:bg-gray-700"
          />
          <Moon
            className={`absolute left-1.5 h-3.5 w-3.5 text-white pointer-events-none transition-opacity ${
              isDark ? "opacity-100" : "opacity-0"
            }`}
          />
          <Sun
            className={`absolute right-1.5 h-3.5 w-3.5 text-gray-600 pointer-events-none transition-opacity ${
              isDark ? "opacity-0" : "opacity-100"
            }`}
          />
        </>
      ) : (
        // Placeholder that matches the switch dimensions
        <div className="h-6 w-11 rounded-full bg-gray-300 dark:bg-gray-700" />
      )}
    </div>
  );
};

export { ThemeToggle };
