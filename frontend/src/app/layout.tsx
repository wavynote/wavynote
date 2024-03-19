import { ReactNode } from "react";
import type { Metadata } from "next";
import Nav from "@/components/Nav";
import Template from "@/components/Template";

import "./globals.css";
import "@/assets/scss/style.scss";


export const metadata: Metadata = {
  title: "WavyNote",
  description: "",
  icons: {
    icon: "/favicon.ico",
  },
};

// const nanumSquareNeo = localFont({
//   src: "../fonts/NanumSquareNeo-Variable.woff2",
//   display: 'swap',
// });

export default function RootLayout({ children } : { children: ReactNode}) {
  return (
    <Template>
      { children }
      <Nav></Nav>
    </Template>
  );
}
