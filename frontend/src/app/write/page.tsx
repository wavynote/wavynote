"use client"
import React from "react";
import { useState } from "react";

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function folderEdit() {

  const [isDone, setIsOpen] = useState(false);
  const [isReadOnly, setReadOnly] = useState(false);
  
  // 수정, 저장 버튼 토글
  function toggleBtn(){
    setIsOpen((isDone) => !isDone);
    setReadOnly((isReadOnly) => !isReadOnly);
  }

  // 토픽 닫기
  const [isClose, setClose] = useState(true);
  function closeTopic(){
    setClose((isClose) => !isClose);
  }

  // 바텀 시트 토글
  const [isBottomOpen, setBottom] = useState(false);
  function toggleBottom(){
    
    setBottom((isBottomOpen) => !isBottomOpen);
  }

  return (
    <div className="contentMin">
      <div className="">
        <section>
        <header className="header">
            <div className="titleWrap">
              <Link href="/" className="prev">
              </Link>
              <h2></h2>
            </div>
            <div className="headerBtnWrap">
              <IconBtn name="" type="etc"></IconBtn>
              {!isDone && <button className="textBtn light" onClick={toggleBtn}>저장</button>}
              {isDone && <div className="headerBtnMin">
                <button type="button" className="textBtn dark" onClick={toggleBtn}>수정</button>
                <button type="submit" className="textBtn dark" onClick={toggleBottom}>보내기</button>
              </div>}
            </div>            
          </header>
        </section>
        {isClose && <section className="topicWrap">
          <p className="topic">
          요즘 읽고 있는 책이 있나요? <br/> 또는 좋아하는 책에 대한 이야기를 들려주세요.
          </p>
          <p className="icBtnWrap">
            <button type="button" className="icBtn closeBtn" onClick={closeTopic}></button></p>
        </section>}
        <section className="writeWrap">
          <textarea name="postContent" placeholder="노트를 시작해주세요" readOnly={isReadOnly}></textarea> 
        </section>
      </div>
      {isBottomOpen && <div className="bottomSheetWrap">
        <div className="bottomSheetMin">
          <div className="bottomSheetTitle">
            <Link href="#" className="prev" onClick={toggleBottom}></Link>
            <p><b>누구</b>에게 보낼까요?</p>
          </div>
          <ul className="bottomList">
            <li><button>랜덤으로 보내기</button></li>
            <li><button>오픈노트에 공유하기</button></li>
            <li><button>'나의 친구'에게 보내기</button></li>
            <li><button>'나의 친구'에게 보내기</button></li>
          </ul>
        </div>
      </div>}
    </div>
  );
}
