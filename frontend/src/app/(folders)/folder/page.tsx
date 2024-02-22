'use client';
import { useEffect, useState } from 'react';

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

// import { getFolderList }  from "@/service/folder";

// const folderListTest = ['기본폴더','나의폴더','좋아요'];

interface Folder {
  folder_id: string;
  folder_name: string;
  note_count: number;
}

export default function FolderList() {

  // const [folders, setFolders] = useState<Folder[]>([]);

  // useEffect(() => {
  //   const fetchFolderList = async (userId:string) => {
  //     try {
  //       const response = await fetch(`/wavynote/v1.0/main/folderlist?id=${userId}`, {
  //         method: 'GET',
  //         cache: 'no-store',
  //         headers: {
  //           'Authorization': 'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
  //           'Content-Type': 'application/json',
  //         },
  //       });
  //       const responseData = await response.json();
  //       // responseData의 형태가 { "data": [...] } 이므로 data 속성을 가져옴
  //       const folderData = responseData.data;
  //       // folderData의 각 요소에서 folder_name과 note_count를 추출하여 새로운 배열을 생성
  //       const formattedFolders = folderData.map((folder:Folder) => ({
  //         folderName: folder.folder_name,
  //         noteCount: folder.note_count
  //       }));
  //       // 새로 가공된 배열을 상태로 설정
  //       setFolders(formattedFolders);
  //     } catch (error) {
  //       console.error('Error fetching folder list:', error);
  //     }
  //   };

  //   fetchFolderList('somebody@naver.com');
  // }, []);

  const [folders, setFolders] = useState<Folder[]>([]);

  useEffect(() => {
    const fetchFolderList = async (userId: string) => {
      try {
        const response = await fetch(`/wavynote/v1.0/main/folderlist?id=${userId}`, {
          method: 'GET',
          cache: 'no-store',
          headers: {
            'Authorization': 'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
            'Content-Type': 'application/json',
          },
        });
        const data = await response.json();
        setFolders(data.data); // Here, we extract the 'data' array from the response
      } catch (error) {
        console.error('Error fetching folder list:', error);
      }
    };

    fetchFolderList('somebody@naver.com');
  }, []);

  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <h2>폴더</h2>
            </div>
            <div className="headerBtnWrap">
              <TextBtn name="편집" type="light"></TextBtn>
              <IconBtn name="" type="search"></IconBtn>
            </div>
          </header>
        </section>
        <section className="bgScroll">
          <ul className="folderWrap">
            <li className="folderMin">
              <button name="폴더추가" className="FolderBtn dark"></button>
            </li>
            {folders.map((folder, index) => (
              <li key={folder.folder_id} className="folderMin">
                <button className="FolderBtn light">
                  {folder.folder_name}
                </button>
              </li>
            ))}
             
            {/*<li className="folderMin">
              <FolderBtn name="기본 폴더" type="light"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="나의 폴더" type="light"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="보낸노트" type="light"></FolderBtn>
            </li> */}
          </ul>    
        </section>
      </div>
    </div>
  );
}
