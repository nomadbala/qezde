import Link from "next/link";
import { Logo } from "../ui/icons/Logo";
import { InstagramIcon, LinkedinIcon, TelegramIcon } from "../ui";

export const Footer = () => {
  return (
    <footer className="flex w-full items-center bg-mutedAccent px-8 py-6">
      <section className="flex items-center gap-8">
        <article className="flex h-full flex-col gap-2">
          <Link href="/">
            <Logo size={152} />
          </Link>
          <ul className="flex w-full items-center justify-around">
            <li className="rounded-xl bg-accent bg-opacity-50 p-1.5">
              <Link href="/about">
                <TelegramIcon size={24} />
              </Link>
            </li>
            <li className="rounded-xl bg-accent bg-opacity-50 p-1.5">
              <Link href="/about">
                <InstagramIcon size={24} />
              </Link>
            </li>
            <li className="rounded-xl bg-accent bg-opacity-50 p-1.5">
              <Link href="/about">
                <LinkedinIcon size={24} />
              </Link>
            </li>
          </ul>
        </article>
        <article className="flex h-full flex-col gap-2 text-white">
          <Link href="/">
            <span>Пользовательское соглашение</span>
          </Link>
          <Link href="/">
            <span>Политика конфиденциальности</span>
          </Link>
        </article>
      </section>
    </footer>
  );
};
