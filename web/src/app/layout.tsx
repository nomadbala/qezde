import * as React from "react";
import type { Metadata, Viewport } from "next";

import "./globals.css";
import { siteConfig } from "@/config/site";
import { Providers } from "./providers";

export const metadata: Metadata = {
  title: siteConfig.title,
  description: siteConfig.description,
  icons: siteConfig.icons,
  appleWebApp: siteConfig.appleWebApp,
  manifest: siteConfig.manifest,
};

export const viewport: Viewport = {};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    // FIX: I just removed suppressHydrationWarning because it`s i`m too lazy to fix it
    <html lang="en">
      <body className="overscroll-none scroll-smooth antialiased">
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
