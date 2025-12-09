import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function toStringSafe(
  value: number | string | null | undefined | unknown
): string {
  return value === null ? "" : String(value);
}

export function toNumberSafe(
  value: number | string | null | undefined | unknown
): number {
  if (value == null) return 0;
  if (typeof value === "number") return value;
  const parsed = Number(value);
  return Number.isNaN(parsed) ? 0 : parsed;
}

export function formatDate(dateString: string) {
  if (!dateString) return "N/A";

  const date = new Date(dateString);

  // Check if date is valid
  if (Number.isNaN(date.getTime())) {
    return "N/A";
  }

  return date.toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
}

export function getUserInitials(username: string) {
  return username.slice(0, 1).toUpperCase();
}

export function fallBackCopyToClipboard(text: string): string {
  const textArea = document.createElement("textarea");
  textArea.value = text;

  textArea.style.position = "fixed";
  textArea.style.left = "-9999px";
  textArea.style.top = "0";

  document.body.appendChild(textArea);
  textArea.focus();
  textArea.select();

  const success = document.execCommand("copy");
  document.body.removeChild(textArea);

  return success ? "Link copied to clipboard!" : "Failed to copy link";
}
