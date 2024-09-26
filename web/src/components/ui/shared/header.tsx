import Image from "next/image";
import Link from "next/link";

export default function Header() {
  return (
    <header className="w-full h-20 flex items-center px-32 gap-32">
      <Link href="/">
        <Image src="/qezde.png" alt="qezde" width={128} height={128} />
      </Link>
      <menu className="flex gap-8">
        <li>
          <Link href="/">About us</Link>
        </li>
        <li>
          <Link href="#faq">FAQ</Link>
        </li>
      </menu>
    </header>
  );
}
