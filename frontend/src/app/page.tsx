"use client";

import Image from 'next/image'
import styles from './page.module.scss';
import { BrowserRouter, Link } from 'react-router-dom';
import myNote from './assets/images/icons/mynote.svg'
import box from './assets/images/icons/box.svg'
import openNote from './assets/images/icons/opennote.svg'
import myPage from './assets/images/icons/mypage.svg'



export default function Home() {
  return (
		<>
			<BrowserRouter>
			<div className={styles.outContent}>
				끝내보자 웨이비노트!🤩
			</div>
			<div className={styles.contentWrap}>
				<div>
					<header className="header">
						<div className="titleWrap">
							<button className="icBtn prevBtn">이전으로</button>
							<h2>나의노트</h2>
						</div>
						<div className="headerBtnWrap">
							<button>노트선택</button>
							<button className="icBtn searchBtn"></button>
						</div>
					</header>
				</div>
				<div>
					<div className={styles.btnWrap}>
						<div className={styles.newNoteBtn}>
							<button>새로운 노트 쓰기</button>
						</div>
					</div>
				</div>
				<div>	
					<ul className="noteListWrap">
						<li className="noteListMin">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 2줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
					</ul>
				</div>
				<nav className={styles.nav}>
					<div className={styles.navMin}>
						<Link to="">
							<Image
								src={myNote}
								width={20}
								height={20}
								alt="mypage image"
							/>
							<span>나의노트</span>
						</Link>
						<Link to="">
							<Image
								src={box}
								width={20}
								height={16}
								alt="mypage image"
							/>
							<span>받은노트</span>
						</Link>
						<Link to="">
							<Image
								src={openNote}
								width={14}
								height={20}
								alt="mypage image"
							/>
							<span>오픈노트</span>
						</Link>
						<Link to="">
							<Image
								src={myPage}
								width={17}
								height={18}
								alt="mypage image"
							/>
							<span>마이페이지</span>
						</Link>
					</div>
				</nav>
			</div>
			
			</BrowserRouter>
		</>
  	)
}


