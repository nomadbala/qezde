import type { Metadata } from "next";
import { Montserrat } from "next/font/google"
import "./globals.css";
import { cn } from "@/lib/utils";
import InitializeSuperTokens from "@/lib/supertoken";
import { Providers } from './providers';

InitializeSuperTokens();

const montserrat = Montserrat({subsets: ['latin']})

export const metadata: Metadata = {
  title: "Qezde",
  description: "Qezde - alternative to traditional maps of Qazaqstan",
  keywords: "Qezde, Qazaqstan, Kazakhstan, Astana, travel, guide, maps",
  authors: {
    "name": "Qezde Team"
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="scroll-smooth">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </head>
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
