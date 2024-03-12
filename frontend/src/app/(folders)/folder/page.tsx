'use client';
import { useEffect, useState } from 'react';

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

// import { getFolderList }  from "@/service/folder";

interface Folder {
  folder_id: string;
  folder_name: string;
  note_count: number;
}

export default function FolderList() {

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

    // fetchFolderList(userId);
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
              <button className="textBtn default">편집</button>
              {/* <TextBtn name="편집" type="light" onClick={undefined}></TextBtn> */}
              <IconBtn name="" type="search"></IconBtn>
            </div>
          </header>
        </section>
        <section className="bgScroll">
          <ul className="folderWrap">
            <li className="folderMin">
              <button name="폴더추가" className="FolderBtn dark">폴더추가</button>
            </li>
            {folders.map((folder, index) => (
              <li key={folder.folder_id} className="folderMin">
                <button className="FolderBtn light">
                  {folder.folder_name}
                </button>
              </li>
            ))}
          </ul>    
        </section>
      </div>
    </div>
  );
}
