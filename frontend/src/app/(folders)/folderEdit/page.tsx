import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";
import FolderBtn from "@/components/FolderBtn";


import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function folderEdit() {

  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <h2>폴더</h2>
            </div>
            <div className="headerBtnWrap">
              <TextBtn name="완료" type="dark"></TextBtn>
              <IconBtn name="" type="search"></IconBtn>
            </div>
          </header>
        </section>
        <section className="bgScroll">
          <ul className="folderWrap">
            <li className="folderMin">
              <FolderBtn name="" type="disable"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="" type="focused"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="나의 폴더" type="editable"></FolderBtn>
            </li>
            <li className="folderMin">
              <FolderBtn name="보낸노트" type="editable"></FolderBtn>
            </li>
          </ul>    
        </section>
      </div>
    </div>
  );
}
