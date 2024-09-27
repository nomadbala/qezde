import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/shadcn/accordion";
import { useTranslations } from "next-intl";

export default function FAQSection() {
  const t = useTranslations("faq");

  return (
    <div
      className="mx-auto my-32 w-3/4 scroll-m-40 scroll-smooth rounded-xl md:p-16 md:scroll-m-10"
      id="faq"
    >
      <Accordion type="multiple" className="flex w-full flex-col gap-8">
        <AccordionItem
          value="item-1"
          className="w-full rounded-lg border border-solid border-gray-300 bg-[#f8f9fa] px-8 opacity-100 transition-all duration-150 hover:bg-[#d7dce185]"
        >
          <AccordionTrigger>{t("item1.question")}</AccordionTrigger>
          <AccordionContent>{t("item1.answer")}</AccordionContent>
        </AccordionItem>
        <AccordionItem
          value="item-2"
          className="w-full rounded-lg border border-solid border-gray-300 bg-[#f8f9fa] px-8 text-left opacity-100 transition-all duration-150 hover:bg-[#d7dce185]"
        >
          <AccordionTrigger className="text-left">
            {t("item2.question")}
          </AccordionTrigger>
          <AccordionContent>{t("item2.answer")}</AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}
