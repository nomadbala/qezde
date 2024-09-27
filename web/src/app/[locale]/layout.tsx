import type { Metadata } from "next";
import { Montserrat } from "next/font/google";
import { NextIntlClientProvider } from "next-intl";
import { getMessages } from "next-intl/server";

import "./globals.css";

const montSerrat = Montserrat({
  subsets: ["latin"],
  display: "swap",
  weight: ["500"],
});

export const metadata: Metadata = {
  title: "Qezde",
  description: "Qezde is modern guide to Astana",
  keywords:
    "Astana guide, best places in Astana, cafes in Astana, restaurants, coworking spaces, Астана, кафе, кафе Астана, Астана кафе, Астана рестораны, рестораны, коворкинги, коворкинги Астана, бесплатные коворкинги",
  authors: [
    {
      name: "qezde",
      url: "https://qezde.com",
    },
  ],
  creator: "QezdeTeam",
  twitter: {
    card: "summary_large_image",
    title: "Qezde — Guide to Astana",
    description: "Discover the best places in Astana, Kazakhstan.",
    images: ["https://qezde.kz/twitter-image.jpg"],
  },
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const messages = await getMessages();

  return (
    <html lang="en" className="scroll-smooth">
      <head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </head>
      <body className={montSerrat.className}>
        <NextIntlClientProvider messages={messages}>
          {children}
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
