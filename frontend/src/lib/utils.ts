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
