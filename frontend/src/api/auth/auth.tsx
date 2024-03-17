// api/auth.ts
import crypto from "crypto";

export async function loginUser(id: string, password: string) {
  const hashedPassword = crypto.createHash('sha256').update(password).digest('hex');

  const credentials = {
    id: id,
    password: hashedPassword,
  };

  try {
    const response = await fetch('/wavynote/v1.0/profile/signin', {
      method: 'POST',
      headers: {
        'Authorization': `Basic d2F2eW5vdGU6d2F2eTIwMjMwOTE0`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    });

    if (!response.ok) {
      throw new Error('Failed to login');
    }

    const userData = await response.json();

    // 로그인 정보를 로컬 스토리지에 저장
    localStorage.setItem('userData', JSON.stringify(userData));

    return userData;

  } catch (error) {
    throw error;
  }
}