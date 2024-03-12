"use client"
import React from "react";
import { useState, useEffect } from "react";

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import { useSearchParams } from 'next/navigation';

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

interface LandingProps {
  userId: string;
  folderId: string;
}

export default function landingPage() {

  // 토글 리스트 수정
  const [isOpenHeaderBtn, setToggleBtn] = useState(true);
  function toggleHeaderBtn(){
    setToggleBtn((isOpenHeaderBtn) => !isOpenHeaderBtn);
  }

  // 
  const [pageTitle, setPageTitle] = useState('');

  const searchParams = useSearchParams();
  
  const userId = searchParams.get('userId');
  const folderId = searchParams.get('folderId');

  useEffect(() => {
    if (userId && folderId) {
      fetchPageTitle(userId, folderId);
    }
  }, [userId, folderId]);

  const fetchPageTitle = async (userId: string, folderId: string) => {
    try {
      const response = await fetch(`/main?userId=${userId}&folderId=${folderId}`);
      const data = await response.json();
      
      console.log(data);
      setPageTitle(data.title);

    } catch (error) {
      console.error('Error fetching page title:', error);
    }
  };

  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <Link href="/folder" className="prev">
              </Link>
              <h2>ddd{pageTitle}</h2>
            </div>
            { isOpenHeaderBtn === false ? ( <div className="headerBtnWrap">
                <TextBtn name="폴더이동" type="light" onClick={undefined}></TextBtn>
                <TextBtn name="삭제" type="light" onClick={undefined}></TextBtn>
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
                <button className="textBtn newNote dark">새로운 노트 쓰기</button>
                {/* <TextBtn name="새로운 노트 쓰기" type="newNote" onClick={undefined}></TextBtn> */}
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
};
