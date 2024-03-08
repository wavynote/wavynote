'use client';

import { useState } from "react";

import "@/assets/scss/style.scss";
import Image from 'next/image';
import TextBtn from "@/components/TextBtn";

export default function loginPage() {

  const [id, setId] = useState('');
  const [password, setPassword] = useState('');
  const [idError, setIdError] = useState('');
  const [passwordError, setPasswordError] = useState('');

  const handleLogin = async (e:any) => {
    e.preventDefault();
    
    debugger;

    // 초기화
    setIdError('');
    setPasswordError('');

    // 유효성 검사
    let isValid = true;
    if (!id) {
      setIdError('ID를 입력해주세요.');
      isValid = false;
    }
    if (!password) {
      setPasswordError('Password를 입력해주세요.');
      isValid = false;
    }
    if (!isValid) {
      return;
    }

    const credentials = {
      id: id,
      password: password
    };

    try {
      const response = await fetch('/wavynote/v1.0/profile/signin', {
        method: 'POST',
        headers: {
          'Authorization': `Basic ${btoa('ID:PW')}`, // Replace 'ID:PW' with your actual credentials
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
      });

      if (!response.ok) {
        throw new Error('Failed to login');
      }

      const data = await response.json();
      // Handle successful login, e.g., redirect to another page
      console.log(data);
    } catch (error) {
      console.error(error);
      // 서버 오류 또는 인증 오류에 대한 메시지 표시
      setIdError('로그인에 실패했습니다. ID와 Password를 확인해주세요.');
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
  