import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";


import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function HomePage() {
  return (
    <div className="contentMin">
      <div className="">
        <section>
          <header className="header">
            <div className="titleWrap">
              <Link href="/" className="prev">
              </Link>
              <h2>나의노트</h2>
            </div>
            <div className="headerBtnWrap">
              <TextBtn name="노트선택" type="light"></TextBtn>
              <IconBtn name="" type="search"></IconBtn>
            </div>
          </header>
        </section>
        <section>
          <div className="btnWrap">
            <div className="newNoteBtn">
              <TextBtn name="새로운 노트 쓰기" type="newNote"></TextBtn>
            </div>
          </div>
        </section>
        <section className="noteListWrap">	
          {/* <div className="noteListBtnWrap">
            <div className="noteListBtnMin">
              <TextBtn name="폴더이동" type="light"></TextBtn>
              <TextBtn name="삭제" type="dark"></TextBtn>
            </div>
          </div> */}
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
            </ul>
          </div>
        </section>
        {/* <div className="noteListWrap">
          <p className="noList">
            새로운 노트를
            <br />
            써보세요
          </p>
        </div> */}
      </div>
    </div>
  );
}
