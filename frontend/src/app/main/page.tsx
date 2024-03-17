"use client"
import React from "react";
import { useState, useEffect } from "react";

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import { useSearchParams } from 'next/navigation';

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

// interface LandingProps {
//   userId: string;
//   folderId: string;
// }

interface Note {
  note_id: string;
  title: string;
  preview: string;
}

interface Folder {
  folder_id: string;
  folder_name: string;
}

export default function landingPage() {

  // 토글 리스트 수정
  const [isOpenHeaderBtn, setToggleBtn] = useState(true);
  function toggleHeaderBtn(){
    setToggleBtn((isOpenHeaderBtn) => !isOpenHeaderBtn);
  }

  //
  const [pageTitle, setPageTitle] = useState('');
  const [folders, setFolders] = useState<Folder[]>([]);
  const [notes, setNotes] = useState<Note[]>([]);

  useEffect(() => {
    // 로그인에서 받아온 데이터
    const userDataString = localStorage.getItem("userData");
    const userData = JSON.parse(userDataString || '{}');

    const userId = userData.user_id;
    fetchFolders(userId);
  }, []);

  const fetchFolders = async (userId:string) => {

    try {
      const response = await fetch(`/wavynote/v1.0/main/folderlist?id=${userId}`,{
        method: 'GET',
        cache: 'no-store',
        headers: {
          'Authorization': 'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
          'Content-Type': 'application/json',
        },
      });

      const data = await response.json();
      setFolders(data.data);
      console.info("fetchFolder 실행시 응답받은 값은 아마도 존재하는 모든 폴더목록 : "+ data.data);
    } catch (error) {
      console.error('Error fetching folders:', error);
    }
  };

  useEffect(() => {
    // 로그인에서 받아온 데이터
    const userDataString = localStorage.getItem("userData");
    const userData = JSON.parse(userDataString || '{}');
    const userId = userData.user_id;
    const folderId = userData.folder_id;

    if (folderId) {
      findFolder(userId, folderId);
    }
  }, [folders]);

  const findFolder = async(userId:string, folderId:string)=>{
    if (folderId) {
      console.log("folders" + folders);
      const selectedFolder = folders.find(folder => folder.folder_id === folderId);
      console.log("selectedFolder : "+ selectedFolder);
      
      if (selectedFolder) {
        setPageTitle(selectedFolder.folder_name);
        fetchNotes(userId, folderId);
      }     
    }
  }

  const fetchNotes = async (userId:string, folderId: string) => {
    try {
      const response = await fetch(`/wavynote/v1.0/main/notelist?uid=${userId}&fid=${folderId}`,{
        method: 'GET',
        cache: 'no-store',
        headers: {
          'Authorization': 'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
          'Content-Type': 'application/json',
        },
      });

      if (response.status === 404) {
        // 404 일때 빈 배열로 넘김
        setNotes([]);
      } else {
        const data = await response.json();
        setNotes(data.data);
      }
    } catch (error) {
      console.error('Error fetching notes:', error);
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
              <h2>{pageTitle}</h2>
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
            {notes.length === 0 ? (
              <p className="noList">
                새로운 노트를<br />써보세요
              </p>
              ) : ( 
              <ul className="noteList">
                {notes.map(note => (
                  <li key={note.note_id} className="list">
                    <Link href={`/note/${note.note_id}`} className="noteCont">
                      <h4>{note.title}</h4>
                      <p>{note.preview}</p>
                    </Link>							
                  </li>
                ))}      
              </ul> 
              )
            }
          </div>
        </section>        
      </div>
    </div>
  );
};
