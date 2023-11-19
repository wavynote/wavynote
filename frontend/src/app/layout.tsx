import { ReactNode } from 'react';
import type { Metadata } from 'next';
import localFont from 'next/font/local';

import './globals.css';

export const metadata: Metadata = {
  title: 'WavyNote',
  description: '',
};

const nanumSquareNeo = localFont({
  src: '../fonts/NanumSquareNeo-Variable.woff2',
  display: 'swap',
});

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en" className={nanumSquareNeo.className}>
      <body className="flex justify-center items-center w-full h-screen">
        <main className="w-[390px] h-[844px] bg-gray-100">{children}</main>
      </body>
    </html>
  );
}
