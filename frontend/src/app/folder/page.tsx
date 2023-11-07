"use client";

import styles from './page.module.scss';
import { BrowserRouter, Link } from 'react-router-dom';

import Image from 'next/image'
import Nav from './../components/nav'
import BtnDel from './../assets/images/icons/ic_del_white.svg'


export default function Home() {
  return (
		<>
			<BrowserRouter>
			<div className={styles.contentWrap}>
				<div className={styles.headerWrap}>
					<header className="header">
						<div className="titleWrap">
							<h2>폴더</h2>
						</div>
						<div className="headerBtnWrap">
                            <button>편집</button>
							<button className="icBtn searchBtn"></button>
						</div>
					</header>
				</div>
				<div className={styles.folderWrap}>
					<div className={styles.folderMin}>
						<div className={styles.addFolder}>
							<h3 className={styles.title}>폴더추가</h3>
						</div>
						<div className={styles.folder}>
							<h3 className={styles.title}>
							폴더이름 폴더이름 폴더이름 폴더이름 폴더이름 폴더이름 폴더이름 폴더이름 
							</h3>
							<p className={styles.noteCount}>노트 5개</p>
						</div>
						<div className={styles.folder}>
							<h3 className={styles.title}>나의 노트</h3>
							<p className={styles.noteCount}>노트 5개</p>
						</div>
						<div className={styles.folder}>
							<h3 className={styles.title}>좋아요</h3>
							<p className={styles.noteCount}>노트 5개</p>
						</div>
						<div className={styles.editFolder}>
							<input type="text" className={styles.editTitle} placeholder="제목 입력"/>
							<p className={styles.noteCount}>노트 5개</p>
							<button className={styles.btnDelFolder}>
							<Image
								src={BtnDel}
								width={8}
								height={8}
								alt="del image"
							/>
							</button>
						</div>
					</div>
				</div>
			</div>
			<Nav/>
			</BrowserRouter>
		</>
  	)
}