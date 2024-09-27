import Footer from "@/components/ui/shared/footer";
import Header from "@/components/ui/shared/header";

export default function HomeLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <main className="flex-1 bg-[#f8f9fa]">{children}</main>
      <Footer />
    </div>
  );
}
