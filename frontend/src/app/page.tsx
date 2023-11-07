"use client";


import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';

import Image from 'next/image'
import styles from './page.module.scss';

import Nav from './components/nav'


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
				<div className="noteListWrap">
					<p className="noList">새로운 노트를<br/>써보세요</p>
				</div>
			</div>
			<Nav/>
			</BrowserRouter>
		</>
  	)
}
