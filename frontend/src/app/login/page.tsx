'use client';

import { useState } from "react";

import "@/assets/scss/style.scss";
import Image from 'next/image';
import TextBtn from "@/components/TextBtn";

export default function loginPage() {

  const handleLogin = () => {
    console.log('로그인시도');
  }

  return (
    <>
      <div className="loginWrap">
        <div className="loginMin">
          <div className="stepInfo">
            <h2>이메일과 비밀번호만으로 웨이비노트를 이용할 수 있어요!</h2>
          </div>
          <div className="loginInfo">
            <form action="" method="">
              <div className="inputWrap">
                <div className="inputMin">
                  {/* <label for="userId">이메일 입력</label> */}
                  <input id="userId" type="email" required placeholder="이메일을 입력해주세요"></input>
                  <p className="errorText">정확한 이메일 주소를 입력해주세요</p>
                </div>
                <div className="inputMin">
                  {/* <label for="userId">이메일 입력</label> */}
                  <input id="userPw" type="password" required placeholder="비밀번호를 입력해주세요" ></input>
                  <p className="errorText">비밀번호는 8~20자 이내로 영문 대소문자, 숫자, 특수문자 중 3가지 이상 혼용하여 입력해 주세요.</p>
                </div>
              </div>              
              <div className="btnWrap">
                <TextBtn name="웨이비노트 시작하기" type="dark" onClick={handleLogin}></TextBtn>
              </div>
            </form>
            <div className="loginHelpWrap">
              <button>아이디 찾기</button>
              <button>비밀번호 재설정</button>
            </div>
          </div>
          <div className="otherLoginWrap">
            <div className="text">
              또는<br/> 
              다른 서비스 계정으로 시작하기
            </div>
            <div className="selectLogin">
              <button className="kakao">카카오 로그인</button>
              <button className="naver">네이버 로그인</button>
              <button className="google">구글 로그인</button>
            </div>
          </div>          
        </div>        
      </div>
    </>
  );
}
  