'use client';

import { useState, useEffect } from "react";
import Link from "next/link";

import "@/assets/scss/style.scss";
import Image from 'next/image';

export default function introPage() {

  const [isActiveTab, setTab] = useState(0);
  
  const classMap: { [key: number]: string }  = {
    0: 'slide01',
    1: 'slide02',
    2: 'slide03',
  };

  const getClassName = (isActiveTab:number) => {
    const additionalClass = classMap[isActiveTab] || '';
    return `${additionalClass}`;
  };

  // 버튼 클릭 시 탭 변경 함수
  const changeTab = (tab: number) => {
    setTab(tab);
  };

  // 자동 변경 타이머 설정
  useEffect(() => {
    const interval = setInterval(() => {
      // isActiveTab를 1씩 증가시키고 2를 넘어가면 0으로 초기화
      setTab((prevTab) => (prevTab + 1) % 3);
    }, 3000);
    
    // 컴포넌트가 언마운트 될 때 타이머 해제
    return () => clearInterval(interval);
  }, []); // 빈 배열을 전달하여 컴포넌트가 처음 렌더링 될 때만 타이머가 설정되도록 함

  return (
    <>
      <div className="introWrap">
        <div className={`introMin ${getClassName(isActiveTab)}`}>
          <div className="slideWrap">
            { isActiveTab === 0 && <section className="introText">
              메모장에 적어둔 글 하나씩은 있잖아요, <br/>
              익명의 누군가와 노트로 나누어보세요.
            </section> }
            { isActiveTab === 1 && <section className="introText">
              공통의 관심사를 기반으로 최대 3명과<br/> 
              매칭되어 노트를 주고 받을 수 있어요.
            </section> }
            { isActiveTab === 2 && <section className="introText">
              다양한 이야기를 읽고 <br/>
              마음에 드는 유저와 노트를<br/>
              주고 받을 수 있는 기회가 있어요.
            </section> }
          </div>
          <section className="introTabWrap">
            <button className={ isActiveTab === 0 ? 'active': ''} onClick={() => changeTab(0)}>1</button>
            <button className={ isActiveTab === 1 ? 'active': ''} onClick={() => changeTab(1)}>2</button>
            <button className={ isActiveTab === 2 ? 'active': ''} onClick={() => changeTab(2)}>3</button>
          </section>            
          <section className="btnWrap">
            <Link href="/login">
            <button className="textBtn dark">
              웨이비노트 시작하기
            </button></Link>
          </section>
        </div>
      </div>
    </>
  );
}
  