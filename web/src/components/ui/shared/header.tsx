import Image from "next/image";
import Link from "next/link";
import { useTranslations } from "next-intl";

export default function Header() {
  const t = useTranslations("header");

  return (
    <header className="w-full h-20 flex items-center px-32 gap-32">
      <Link href="/">
        <Image src="/qezde.png" alt="qezde" width={128} height={128} />
      </Link>
      <menu className="flex gap-8">
        <li>
          <Link href="/">{t("item1.title")}</Link>
        </li>
        <li>
          <Link href="#faq">{t("item2.title")}</Link>
        </li>
      </menu>
    </header>
  );
}
