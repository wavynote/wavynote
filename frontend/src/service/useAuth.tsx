'use client'

import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";

export default function useAuth() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const pathname = usePathname();
  const router = useRouter();

  useEffect(() => {
    const checkLoggedIn = async () => {
      if (typeof window !== 'undefined') {
        const loggedIn = localStorage.getItem("isLoggedIn");
        if (loggedIn === "true") {
          setIsLoggedIn(true);
        } else {
          setIsLoggedIn(false);
          if (!pathname.startsWith("/login") && !pathname.startsWith("/intro")) {
            router.push("/login");
          }
        }
      }
    };

    checkLoggedIn();
  }, [pathname]);

  const handleLogout = () => {
    localStorage.removeItem("isLoggedIn");
    setIsLoggedIn(false);
    router.push("/login");
  };

  return { isLoggedIn, handleLogout };
}