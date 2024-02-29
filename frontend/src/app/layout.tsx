import { ReactNode } from "react";
import type { Metadata } from "next";
import localFont from "next/font/local";

import "./globals.css";
import "@/assets/scss/style.scss";

import Nav from "@/components/Nav";

export const metadata: Metadata = {
  title: "WavyNote",
  description: "",
  icons: {
    icon: "/favicon.ico",
  },
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
        {/* 구조 바꿔야할것 같음 for popup 이랑 bottomSheet 때문에 */}
      </body>
    </html>
  );
}
