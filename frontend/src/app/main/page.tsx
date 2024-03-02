"use client"
import React from "react";
import { useState } from "react";

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function notePage() {

  // 토글 리스트 수정
  const [isOpenHeaderBtn, setToggleBtn] = useState(true);
  function toggleHeaderBtn(){
    setToggleBtn((isOpenHeaderBtn) => !isOpenHeaderBtn);
  }

  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <Link href="/folder" className="prev">
              </Link>
              <h2>나의노트</h2>
            </div>
            { isOpenHeaderBtn === false ? ( <div className="headerBtnWrap">
                <TextBtn name="폴더이동" type="light"></TextBtn>
                <TextBtn name="삭제" type="light"></TextBtn>
                <button className="textBtn light" onClick={toggleHeaderBtn}>완료</button>
              </div> ) : ( <div className="headerBtnWrap">
                <button className="textBtn light" onClick={toggleHeaderBtn}>노트선택</button>
                <Link href="/search" className="searchLink"><IconBtn name="" type="search"></IconBtn></Link>
              </div> )
            }      
          </header>
        </section>
        <section>
          <div className="btnWrap">
            <div className="newNoteBtn">
              <Link href="/write">
                <TextBtn name="새로운 노트 쓰기" type="newNote"></TextBtn>
              </Link>
            </div>
          </div>
        </section>
        <section className="noteListWrap">
          <div className="noteListMin">
            <ul className="noteList">
              <li className="list">
                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li>
              {/* <li className="list">

                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li>
              <li className="list">

                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li>
              <li className="list">

                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li>
              <li className="list">

                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li>
              <li className="list">

                <Link href="/" className="noteCont">
                    <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                    <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                    <div className="tagArea">
                      <span className="tagDate">2023.01.01</span>
                    </div>
                </Link>							
              </li> */}
              
            </ul>
          </div>
        </section>
        {/* <div className="noteListWrap">
          <div className="noteListMin">
            <p className="noList">
              새로운 노트를
              <br />
              써보세요
            </p>
          </div>          
        </div>*/}
      </div>
    </div>
  );
}
