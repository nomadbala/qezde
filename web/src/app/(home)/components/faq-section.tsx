import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/shadcn/accordion";

export default function FAQSection() {
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
          <AccordionTrigger>What is Qezde?</AccordionTrigger>
          <AccordionContent>
            Qezde is an online platform that helps people discover the best
            places in Astana. It’s a guide to cafes, restaurants, coworking
            spaces, and other cool spots in the city, based on user reviews.
          </AccordionContent>
        </AccordionItem>
        <AccordionItem
          value="item-2"
          className="w-full rounded-lg border border-solid border-gray-300 bg-[#f8f9fa] px-8 text-left opacity-100 transition-all duration-150 hover:bg-[#d7dce185]"
        >
          <AccordionTrigger className="text-left">
            How do I contact Qezde support?
          </AccordionTrigger>
          <AccordionContent>
            If you have any questions or issues, contact our support team via
            email at support@qezde.kz or use the contact form on the website.
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}
