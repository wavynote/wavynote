"use client";

import Image from 'next/image'
import styles from './nav.scss';
import { Link } from 'react-router-dom';
import myNote from './../assets/images/icons/mynote.svg'
import box from './../assets/images/icons/box.svg'
import openNote from './../assets/images/icons/opennote.svg'
import myPage from './../assets/images/icons/mypage.svg'

export default function Home() {
  return (
		<>
			<nav className="nav">
				<div className="navMin">
					<Link to="/myNote">
						<Image
							src={myNote}
							width={20}
							height={20}
							alt="mypage image"
						/>
						<span>나의노트</span>
					</Link>
					<Link to="/box">
						<Image
							src={box}
							width={20}
							height={16}
							alt="mypage image"
						/>
						<span>받은노트</span>
					</Link>
					<Link to="/openNote">
						<Image
							src={openNote}
							width={14}
							height={20}
							alt="mypage image"
						/>
						<span>오픈노트</span>
					</Link>
					<Link to="/myPage">
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
		</>
  	)
}