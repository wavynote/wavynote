"use client";

import Link from "next/link";
import Image from "next/image";

import React, { useState } from "react";
import { usePathname, useSearchParams } from 'next/navigation';

// import { ReactComponent as MyNote} from '@/assets/images/icons/myNote.svg'

import myNote from "@/assets/images/icons/myNote.svg";
import box from "@/assets/images/icons/box.svg";
import openNote from "@/assets/images/icons/opennote.svg";
import myPage from "@/assets/images/icons/mypage.svg";

const Nav = () => {
  const pathname = usePathname();
  if( pathname === '/intro' ||  pathname === '/login') return null;
  
  const [activeNav, setActNav] = useState(1);

  return (
    <nav className="nav">
      <div className="navMin">
        <Link href="/main" onClick={() => setActNav(1)}>
          <div className={`myNote ${activeNav === 1 ? "active" : "navItem"}`}>
            <span>나의노트</span>
          </div>          
        </Link>
        <Link href="/box" onClick={() => setActNav(2)}>
          <div className={`box ${activeNav === 2 ? "active" : "navItem"}`}>
            <span>받은노트</span>
          </div> 
        </Link>
        <Link href="/openNote" onClick={() => setActNav(3)}>
          <div className={`openNote ${activeNav === 3 ? "active" : "navItem"}`}>
            <span>오픈노트</span>
          </div> 
        </Link>
        <Link href="/myPage" onClick={() => setActNav(4)}>
          <div className={`myPage ${activeNav === 4 ? "active" : "navItem"}`}>
            <span>마이페이지</span>
          </div> 
        </Link>
      </div>
    </nav>
  );
};

export default Nav;
