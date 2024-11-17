import { layoutConfig } from "./layout";

export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  title: "Qezde",
  description:
    "Qezde is a platform where u can find the best places for you and your friends",
  socialLinks: {
    telegram: "https://t.me/qezdeapp",
    linkedin: "https://www.linkedin.com/company/qezde/",
    instagram: "https://www.instagram.com/qezde/",
  },
  icons: {
    icon: [
      {
        rel: "icon",
        type: "image/png",
        sizes: "96x96",
        url: "/favicon-96x96.png",
      },
      {
        rel: "icon",
        type: "image/svg+xml",
        sizes: undefined,
        url: "/favicon.svg",
      },
      {
        rel: "shortcut icon",
        type: "image/png",
        sizes: undefined,
        url: "/favicon.ico",
      },
    ],
    apple: "/apple-touch-icon.png",
  },
  appleWebApp: {
    title: "Qezde",
  },
  manifest: "/site.webmanifest",
  layout: layoutConfig,
};
