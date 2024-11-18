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
      className={`${notoSans.className} relative flex min-h-screen flex-col items-center overflow-x-hidden overflow-y-hidden bg-swhite pt-4`}
    >
      <div className="absolute bottom-0 right-0 h-40 w-40 translate-x-1/2 animate-pulse overflow-hidden rounded-full bg-mutedAccent bg-opacity-35 sm:bottom-0 sm:h-60 sm:w-60"></div>
      <div className="absolute left-0 top-0 h-40 w-40 -translate-x-1/2 -translate-y-1/2 animate-pulse overflow-hidden rounded-full bg-mutedAccent bg-opacity-35 sm:bottom-1/2 sm:h-60 sm:w-60"></div>
      <NavigationBar />
      <main>{children}</main>
      <Footer />
    </div>
  );
}
