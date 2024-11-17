"use client";

import { cn } from "@/lib/utils";
import { Readex_Pro } from "next/font/google";
import Image from "next/image";

const readexPro = Readex_Pro({
  subsets: ["latin"],
  weight: ["200", "300", "400", "500", "600", "700"],
});

export default function AboutPage() {
  return (
    <section
      className={cn(
        "flex min-h-screen flex-col items-center py-8",
        readexPro.className,
      )}
    >
      <div className="absolute right-0 top-0 -z-10 h-40 w-32 animate-pulse rounded-full bg-blue-300 opacity-50 sm:bottom-0 sm:h-60 sm:w-60"></div>
      <div className="absolute bottom-0 left-0 -z-10 h-40 w-32 animate-pulse rounded-full bg-blue-300 opacity-50 sm:bottom-0 sm:h-60 sm:w-60"></div>

      <article className="flex w-11/12 justify-between">
        <section className="flex w-1/2 flex-col justify-around px-8 py-8">
          <h2 className="text-4xl font-bold">Чё за Qezde?</h2>
          <h3 className="text-2xl">
            Qezde — это приложение, где ты найдешь лучшие места для отдыха,
            работы и тусовок в Астане. Мы про удобство, стиль и немного
            дерзости.
          </h3>
        </section>
        <figure className="flex w-1/2 justify-center">
          <Image
            src={"/kama-tulkibayeva-BpugFbwxOfw-unsplash.jpg"}
            alt="about"
            width={324}
            height={324}
            className="rounded-lg"
          />
        </figure>
      </article>
    </section>
  );
}
