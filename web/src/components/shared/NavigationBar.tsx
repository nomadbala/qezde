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
      shouldHideOnScroll
      isBordered
      className={`bg-mutedAccent w-11/12 rounded-full py-4 shadow-lg`}
    >
      <NavbarContent className="basis-1 pl-4 sm:hidden" justify="start">
        <NavbarMenuToggle
          aria-label={isMenuOpen ? "Close menu" : "Open menu"}
        />
      </NavbarContent>

      <NavbarContent className="sm:hidden" justify="center">
        <NavbarBrand>
          <Link href="/">
            <Logo size={152} />
          </Link>
        </NavbarBrand>
      </NavbarContent>

      <NavbarContent className="hidden gap-4 sm:flex" justify="center">
        <NavbarBrand>
          <Link href="/">
            <Logo size={152} />
          </Link>
        </NavbarBrand>
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
      </NavbarContent>

      <NavbarContent justify="end">
        <NavbarItem>
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

      <NavbarMenu className="">
        <div className="mx-4 mt-2 flex flex-col gap-2">
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
