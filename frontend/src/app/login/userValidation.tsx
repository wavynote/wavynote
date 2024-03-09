import { useState } from "react";

export default function useValidation() {
  const [id, setId] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [idError, setIdError] = useState<string>('');
  const [passwordError, setPasswordError] = useState<string>('');

  function validateInput(id: string, password: string): boolean {
    // 초기화
    setIdError('');
    setPasswordError('');

    function isValidEmail(id: string): boolean {
      // 간단한 형식의 @ 문자를 포함하는지만 확인하는 이메일 유효성 검사를 수행합니다.
      return /\S+@\S+\.\S+/.test(id);
    }

    // 유효성 검사
    let isValid = true;
    if (!id) {
      setIdError('ID를 입력해주세요.');
      isValid = false;
    } else if (!isValidEmail(id)) {
      setIdError('올바른 이메일 형식이 아닙니다.');
      isValid = false;
    }

    if (!password) {
      setPasswordError('비밀번호를 입력해주세요.');
      isValid = false;
    }

    return isValid;
  }

  return { id, setId, password, setPassword, idError, passwordError, setPasswordError, validateInput };
}