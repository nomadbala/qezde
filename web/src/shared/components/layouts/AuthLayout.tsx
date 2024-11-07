import { ReactNode } from 'react';
import { oswald } from "@/lib/fonts";
import { cn } from "@/lib/utils";

interface AuthLayoutProps {
  children: ReactNode;
  title: string;
  subtitle?: string;
}

export function AuthLayout({ children, title, subtitle }: AuthLayoutProps) {
  return (
    <div className="flex min-h-screen">
      <div className="flex-1 flex flex-col items-center justify-center">
        <h1 className={cn("text-accent text-4xl font-semibold", oswald.className)}>
          {title}
        </h1>
        {subtitle && (
          <p className={cn("text-white text-center text-xl w-3/5 font-light", oswald.className)}>
            {subtitle}
          </p>
        )}
        {children}
      </div>
    </div>
  );
} 