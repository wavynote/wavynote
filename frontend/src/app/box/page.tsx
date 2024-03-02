"use client"
import React from "react";
import { useState } from "react";

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

// 이미지 테스트
import Image from 'next/image';
import userImg from "@/assets/images/content/userimg.png"

export default function boxPage() {

  // 토글 리스트 수정
  const [isActiveTab, setToggleBtn] = useState(true);
  
  function tabNewNote(){
    setToggleBtn((true));
  }

  function tabMatch(){
    setToggleBtn(false);
  }

  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <Link href="/folder" className="prev">
              </Link>
              <h2>받은노트</h2>
            </div>
            <div className="headerBtnWrap">
                {/* <button className="textBtn light" onClick={toggleHeaderBtn}>노트선택</button> */}
                <Link href="/search" className="searchLink"><IconBtn name="" type="search"></IconBtn></Link>
              </div>
          </header>
        </section>
        <section className="tabWrap">
          <div className="tabMin">
            <button className={isActiveTab ? 'active' : ''} onClick={tabNewNote}>새로운 노트</button>
            <button className={isActiveTab ? '' : 'active'} onClick={tabMatch}>나의 친구들</button>
          </div>
        </section>
        <section className="noteListWrap">
          <div className="noteListMin">
            { isActiveTab === true ? (<ul className="newNoteList">
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
              
            </ul>) : (<ul className="matchList">
              <li className="list">
                <Link href="/" className="matchCont">
                  <div className="imgWrap">
                    <Image src={userImg} alt="이미지테스트"/>
                  </div>
                  <div className="textWrap">
                    <h4>매칭된 유저 별명 매칭된 유저 별명매칭된 유저 별명매칭된 유저 별명매칭된 유저 별명</h4>
                    <p>주고받은 글 <span>2 개</span></p>
                  </div>
                </Link>
              </li>
              <li className="list emptyMatch">
                <div className="matchCont">
                  <div className="textWrap">
                    <h4>두번째 친구 자리</h4>
                    <p>?</p>
                    <p>
                    웨이비노트에서는 세명의 친구와 글을 주고 받을 수 있어요. 노트를 쓴 후 <b>랜덤으로 매칭하기</b>를 클릭해보세요!
                    </p>
                  </div>
                </div>
              </li>
              <li className="list emptyMatch">
                <div className="matchCont">
                  <div className="textWrap">
                    <h4>세번째 친구 자리</h4>
                    <p>?</p>
                    <p>
                    웨이비노트에서는 세명의 친구와 글을 주고 받을 수 있어요. 노트를 쓴 후 <b>랜덤으로 매칭하기</b>를 클릭해보세요!
                    </p>
                  </div>
                </div>
              </li>
            </ul>) 
            }
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
