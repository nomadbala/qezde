"use client";

import { Logo } from "@/components/ui";
import { useState } from "react";
import {
  Navbar,
  NavbarBrand,
  NavbarMenuToggle,
  NavbarMenuItem,
  NavbarMenu,
  NavbarContent,
  NavbarItem,
  Link,
  Button,
} from "@nextui-org/react";
import { layoutConfig } from "@/config";

export const NavigationBar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <Navbar
      isMenuOpen={isMenuOpen}
      onMenuOpenChange={setIsMenuOpen}
      position="static"
      className={`flex w-11/12 items-center rounded-full bg-mutedAccent py-2 shadow-lg`}
    >
      <NavbarContent className="basis-1 pl-4 sm:hidden">
        <NavbarMenuToggle
          aria-label={isMenuOpen ? "Close menu" : "Open menu"}
        />
      </NavbarContent>

      <NavbarContent className="sm:hidden">
        <NavbarBrand>
          <Link href="/">
            <Logo size={152} />
          </Link>
        </NavbarBrand>
      </NavbarContent>

      <NavbarContent className="hidden w-full sm:flex sm:justify-between">
        <NavbarBrand className="flex-1">
          <Link href="/">
            <Logo size={152} />
          </Link>
        </NavbarBrand>

        <div className="flex flex-1 justify-center gap-x-6">
          {layoutConfig.header.navigationTabs.map((item, index) => (
            <NavbarItem key={index}>
              <Link
                color="foreground"
                href={item.link}
                className="relative w-full text-lg text-white transition-colors duration-300 after:absolute after:-bottom-1 after:left-0 after:h-[3px] after:w-0 after:rounded-full after:bg-accent after:transition-all after:duration-300 hover:after:w-full"
              >
                {item.title}
              </Link>
            </NavbarItem>
          ))}
        </div>

        <NavbarItem className="flex flex-1 justify-end">
          <Button
            as={Link}
            href="#"
            variant="flat"
            className={`text-md rounded-full bg-accent bg-opacity-50 text-white`}
          >
            Регистрация
          </Button>
        </NavbarItem>
      </NavbarContent>

      <NavbarMenu className="bg-swhite px-10 pt-10">
        <div className="flex flex-col gap-2">
          {layoutConfig.header.menuTabs.map((item, index) => (
            <NavbarMenuItem key={`${item}-${index}`}>
              <Link
                className="w-full"
                color="foreground"
                href={item.link}
                size="lg"
              >
                {item.title}
              </Link>
            </NavbarMenuItem>
          ))}
        </div>
      </NavbarMenu>
    </Navbar>
  );
};
