
import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";

import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function folderList() {

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
              <FolderBtn name="폴더추가" type="dark"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="기본 폴더" type="light"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="나의 폴더" type="light"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="보낸노트" type="light"></FolderBtn>
            </li>
          </ul>    
        </section>
      </div>
    </div>
  );
}
