import Link from "next/link";

import "@/assets/scss/style.scss";

export default function openNote(){
  return(
    <div className="contentMin">
      <div>
        <section>
          <header className="header">
            <div className="titleWrap">
            </div>
            <div className="headerBtnWrap">
                <button className="textBtn light">수정</button>
              </div>
          </header>
        </section>
        <section className="userInfo">
          <div className="userImg"></div>
          <div className="userName">유저가 설정한 별명</div>
        </section>
        <div className="areaGap"></div>
        <section>
          <ul className="setList">
            <li><Link href="/">계정 관리</Link></li>
            <li><Link href="/">알림 설정</Link></li>
            <li><Link href="/">관심주제 설정</Link></li>
            <li><Link href="/">친구 관리</Link></li>
            <li><Link href="/">글쓰기 설정</Link></li>
            <li><Link href="/">문의하기</Link></li>
            <li><Link href="/">로그아웃</Link></li>
          </ul>
        </section>
      </div>        
    </div>
  )
}