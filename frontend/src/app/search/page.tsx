import IconBtn from "@/components/IconBtn";
import TextBtn from "@/components/TextBtn";


import styles from "./page.module.scss";
import "@/assets/scss/style.scss";

import Link from "next/link";

export default function searchList() {
  return (
    <div className="contentMin">
      <div className="">
        <section className="searchWrap">
          <header className="header">
            <div className="titleWrap">
              <Link href="/" className="prev">
              </Link>
            </div>
            <input type="text" placeholder="검색어를 입력하세요"></input>
            <div className="headerBtnWrap">
              <IconBtn name="" type="searchBlack"></IconBtn>
            </div>
          </header>
        </section>
        {/* <div className="noteListWrap">
          <div className="noteListMin">
            <p className="noList">
              검색 결과가 없습니다.
            </p>
          </div>
        </div> */}
        <section className="noteListWrap">
          <div className="noteListMin">
            <ul className="searchList">
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
              <li className="list">
                <Link href="/" className="noteCont">
                  <h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
                  <p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
                  <div className="tagArea">
                    <span className="tagDate">2023.01.01</span>
                    <span className="toFrom">나의친구가 보냄</span>
                  </div>
                </Link>
              </li>
            </ul>
          </div>
        </section>
      </div>
    </div>
  );
}
