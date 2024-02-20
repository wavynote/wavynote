'use client';
import { useEffect, useState } from 'react';

import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

const folderListTest = ['기본폴더','나의폴더','좋아요'];

export default function folderList() {

  const [folderName, getFolderName] = useState();

  useEffect(() => {
    fetch('/main/folderlist?id=wavynote', {
      method:'GET',
      headers: {
        'Authorization':'Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0',
        'Content-Type':'application/json',
      },})
    .then((res) => res.json())
    .then((data) => alert(data[0]));
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
            { folderListTest.map((folderName,index)=><li className="folderMin">
              <button className="FolderBtn light">
                {folderName}
              </button>
            </li>)}
{/*             
            <li className="folderMin">
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
