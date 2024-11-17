import type { Config } from "tailwindcss";

import { nextui } from "@nextui-org/react";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        accent: "#45693C",
        mutedAccent: "#A2B287",
        swhite: "#FEFAE0",
        olive: "#626F47",
      },
    },
  },
  darkMode: "class",
  plugins: [nextui()],
} satisfies Config;
