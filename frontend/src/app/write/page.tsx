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

  function toggleBtn(){
    setIsOpen((isDone) => !isDone);
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
                <button className="textBtn dark" onClick={toggleBtn}>수정</button>
                <button className="textBtn dark">보내기</button>
              </div>}
            </div>            
          </header>
        </section>
        <section className="topicWrap">
          <p className="topic">
          요즘 읽고 있는 책이 있나요? <br/> 또는 좋아하는 책에 대한 이야기를 들려주세요.
          </p>
        </section>
        <section className="writeWrap">
          <textarea placeholder="노트를 시작해주세요"></textarea> 
        </section>
      </div>
    </div>
  );
}
