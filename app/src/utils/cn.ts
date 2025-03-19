import type { ClassValue } from "clsx";
import clsx from "clsx";
import { twMerge } from "tailwind-merge";

// This function is a wrapper around clsx and tailwind-merge
// and is used by shadcn-Solid components

export const cn = (...classLists: ClassValue[]) => twMerge(clsx(classLists));
