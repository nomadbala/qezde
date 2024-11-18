"use client";

import { Logo } from "@/components/ui/icons/Logo";
import { cn } from "@/lib/utils";
import { PT_Sans } from "next/font/google";
import { Image } from "@nextui-org/image";
import NextImage from "next/image";
import { Tooltip, Link } from "@nextui-org/react";

const merriweather = PT_Sans({
  subsets: ["latin"],
  weight: ["400", "700"],
});

export default function AboutPage() {
  return (
    <section
      className={cn(
        "relative flex min-h-screen flex-col items-center gap-16 overflow-x-hidden overflow-y-hidden py-8",
      )}
    >
      <header className="relative w-8/12 rounded-lg">
        <Image
          src="/la-riviere.jpg"
          alt="header image"
          width={0}
          height={0}
          className="w-full"
        />
        <Tooltip
          showArrow={true}
          content="Нажимая на такие подсказки, вы сможете узнать больше об этом месте"
          color="foreground"
          className="opacity-85"
        >
          <Link
            href="/"
            className="absolute bottom-0 right-0 z-10 mb-1 mr-1 h-10 w-fit text-4xl text-mutedAccent"
          >
            ?
          </Link>
        </Tooltip>
      </header>

      <article className="flex w-11/12 items-start justify-between">
        <section className="flex w-1/2 flex-col justify-between px-8 py-8">
          <h2 className="flex items-center gap-2">
            <Logo size={152} />{" "}
          </h2>
          <h3 className={cn(merriweather.className, "ml-8 text-2xl")}>
            - это приложение, где ты найдешь лучшие места для отдыха, работы и
            тусовок в Астане. Мы про удобство и стиль.
          </h3>
        </section>
        <figure className="flex w-1/2 justify-center">
          <Image
            as={NextImage}
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
