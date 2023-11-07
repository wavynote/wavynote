"use client";

import styles from './page.module.scss';
import Nav from './../components/nav'

import { BrowserRouter, Link } from 'react-router-dom';



export default function Home() {
  return (
		<>
			<BrowserRouter>
			<div className={styles.outContent}>
				끝내보자 웨이비노트!🤩
			</div>
			<div className={styles.contentWrap}>
				<div className={styles.headerWrap}>
					<header className="header">
						<div className="titleWrap">
							<button className="icBtn prevBtn">이전으로</button>
						</div>
						<div className="headerBtnWrap">
                            <button>완료</button>
                            <button>수정</button>
                            <button>보내기</button>
						</div>
					</header>
				</div>
				<div className={styles.writeWrap}>
					<p className={styles.comments}>
                        <span>요즘 읽고 있는 책이 있나요? 또는 좋아하는 책에 대한 이야기를 들려주세요.</span>
                        <button className="icBtn commentsClose">닫기</button>
                    </p>	
					<textarea placeholder="노트를 시작해주세요">
                    </textarea>
				</div>
			</div>
			<Nav />
			</BrowserRouter>
		</>
  	)
}