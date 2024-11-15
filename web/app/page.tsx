import Image from "next/image";

export default function Home() {
  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-5 bg-red-500">
      <Image
        src="/grandalatau.jpg"
        alt="logo"
        width={324}
        height={324}
        className="rounded-xl"
      />
    </section>
  );
}
