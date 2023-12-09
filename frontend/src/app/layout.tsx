import { ReactNode } from "react";
import type { Metadata } from "next";
import localFont from "next/font/local";

import "./globals.css";
import "@/assets/scss/style.scss";

import Nav from "@/components/Nav";

export const metadata: Metadata = {
  title: "WavyNote",
  description: "",
};

const nanumSquareNeo = localFont({
  src: "../fonts/NanumSquareNeo-Variable.woff2",
  display: "swap",
});

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en" className={nanumSquareNeo.className}>
      <body>
        <main className="contentWrap">
          {children}
        </main>
        <Nav />      
      </body>
    </html>
  );
}
