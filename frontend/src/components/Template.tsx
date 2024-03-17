'use client'

import { ReactNode } from "react";
import useAuth from '@/service/useAuth';

type TemplateProps = {
  children: ReactNode;
};

export default function Template({ children }: TemplateProps) {

  if (typeof window !== 'undefined'){
    const { isLoggedIn } = useAuth();
  }

  return (
    <html lang="en">
      <body>
        <main className="contentWrap">
          {children}
        </main>
      </body>
    </html>
  );
}