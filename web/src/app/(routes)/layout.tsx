"use client";

import { Footer, NavigationBar } from "@/components/shared";

import { Noto_Sans } from "next/font/google";

const notoSans = Noto_Sans({
  subsets: ["latin"],
  weight: ["400", "700"],
});

export default function RoutesLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div
      className={`${notoSans.className} flex min-h-screen flex-col items-center bg-swhite pt-4`}
    >
      <NavigationBar />
      <main>{children}</main>
      <Footer />
    </div>
  );
}
