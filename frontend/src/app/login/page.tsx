'use client';

import { useState } from "react";
import { loginUser } from "@/api/auth/auth";

import useValidation from "./userValidation";
import { useRouter } from "next/navigation";


import "@/assets/scss/style.scss";
import Image from 'next/image';
import TextBtn from "@/components/TextBtn";

export default function loginPage() {

  const { id, setId, password, setPassword, idError, passwordError, setPasswordError, validateInput } = useValidation();
  const router = useRouter();

  const handleLogin = async (e:any) => {
    e.preventDefault();

    // 유효성 검사
    const isValid = validateInput(id, password);
    if (!isValid) {
      return;
    }

    try {
      const data = await loginUser(id, password);
      console.log(data);
      return router.push(`/main?${data.user_id, data.folder_id}`);

    } catch (error) {
      console.error(error);
      setPasswordError('로그인에 실패했습니다. 아이디와 비밀번호를 확인해주세요.');
    }
  };

  return (
    <>
      <div className="loginWrap">
        <div className="loginMin">
          <div className="stepInfo">
            <h2>이메일과 비밀번호만으로 웨이비노트를 이용할 수 있어요!</h2>
          </div>
          <div className="loginInfo">
            <form action="" method="POST">
              <div className="inputWrap">
                <div className="inputMin">
                  {/* <label for="userId">이메일 입력</label> */}
                  <input id="userId" type="email" required placeholder="이메일을 입력해주세요" value={id} onChange={(e) => setId(e.target.value)}/>
                  {idError && <p className="errorText">{idError}</p>}
                </div>
                <div className="inputMin">
                  {/* <label for="userId">이메일 입력</label> */}
                  <input id="userPw" type="password" required placeholder="비밀번호를 입력해주세요" value={password} onChange={(e) => setPassword(e.target.value)}/>
                  {passwordError && <p className="errorText">{passwordError}</p>}

                  {/* {passwordError && <p className="errorText">비밀번호는 8~20자 이내로 영문 대소문자, 숫자, 특수문자 중 3가지 이상 혼용하여 입력해 주세요.</p>} */}
                </div>
              </div>              
              <div className="btnWrap">
                <button className="textBtn dark" type="submit" onClick={handleLogin}>웨이비노트 시작하기</button>
              </div>
            </form>
            <div className="loginHelpWrap">
              <button disabled>아이디 찾기</button>
              <button disabled>비밀번호 재설정</button>
            </div>
          </div>
          <div className="otherLoginWrap">
            <div className="text">
              또는<br/> 
              다른 서비스 계정으로 시작하기
            </div>
            <div className="selectLogin">
              <button className="kakao" disabled>카카오 로그인</button>
              <button className="naver" disabled>네이버 로그인</button>
              <button className="google" disabled>구글 로그인</button>
            </div>
          </div>          
        </div>        
      </div>
    </>
  );
}