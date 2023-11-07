"use client";

import { BrowserRouter, Link } from 'react-router-dom';
import Nav from './../components/nav'

import Image from 'next/image'
import styles from './page.module.scss';



export default function Home() {
  return (
		<>
			<BrowserRouter>
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
				<div className="noteListWrap">	
					<ul className="noteListMin">
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
						<li className="noteList">
							<div className="list">
								<h4>리스트 제목입니다. 최대 1줄까지 표시됩니다.</h4>
								<p>글쓰기 내용이 표시됩니다. 최대 1줄까지 표시되고 나머지 내용은 말줄임표로 표시됩니다...</p>
								<div className="tagArea">
									<span className="tagDate">2023.01.01</span>
								</div>
							</div>
						</li>
					</ul>
				</div>
			</div>
			<Nav/>
			</BrowserRouter>
		</>
  	)
}


