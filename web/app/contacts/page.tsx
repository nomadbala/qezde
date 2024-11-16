"use client";

import { Autocomplete, AutocompleteItem } from "@nextui-org/autocomplete";
import { Input, Textarea } from "@nextui-org/input";
import { contactOptions } from "./data";
import { WorkingGuyIcon } from "@/components/icons";

export default function Contacts() {
  return (
    <section className="sm:flex sm:items-center sm:justify-between gap-4 py-8 md:py-5 px-2">
      <div className="absolute top-0 left-16 w-24 h-24 bg-pink-300 rounded-full opacity-50 animate-pulse"></div>
      <div className="absolute bottom-5 sm:bottom-0 right-0 w-32 h-40 sm:w-60 sm:h-60 bg-blue-300 rounded-full opacity-50 animate-pulse"></div>
      <div className="absolute bottom-0 left-0 w-24 h-48 sm:w-48 bg-yellow-300 rounded-full opacity-50 animate-pulse"></div>
      <div className="absolute top-24 right-10 w-16 h-32 sm:w-32 bg-green-300 rounded-full opacity-50 animate-pulse"></div>

      <article className="w-full sm:w-1/2 py-12 flex flex-col items-center gap-8">
        <Autocomplete
          isRequired
          label="Вариант обращения"
          defaultItems={contactOptions}
          placeholder="Выберите вариант"
          defaultSelectedKey="found_bug"
          className="w-full sm:w-2/3"
        >
          {(item) => (
            <AutocompleteItem key={item.value}>{item.label}</AutocompleteItem>
          )}
        </Autocomplete>
        <Input
          type="name"
          label="Имя"
          placeholder="Введите ваше имя"
          className="w-full sm:w-2/3"
        />
        <Input
          type="email"
          label="Почта"
          placeholder="Введите вашу почту"
          className="w-full sm:w-2/3"
        />
        <Textarea
          label="Сообщение"
          placeholder="Введите ваше сообщение"
          className="w-full sm:w-2/3"
        />
      </article>
      <WorkingGuyIcon size={512} className="w-1/2 hidden sm:block" />
    </section>
  );
}
