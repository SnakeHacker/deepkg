
export function setAuthenticated(userInfo: any) {
    localStorage.setItem('isLoggedIn', 'true');
    localStorage.setItem('userInfo', JSON.stringify(userInfo));
    return
}

// 检查用户是否已登录
export function isAuthenticated(): boolean {
  return localStorage.getItem('isLoggedIn') === 'true';
}

export function setAccessToken(token: string) {
    localStorage.setItem('accessToken', token);
    return
}

export function getAccessToken() {
    return localStorage.getItem('accessToken');
}

// 获取当前登录用户名
export function getCurrentUser(): any {
  const storedUserInfo = localStorage.getItem('userInfo');
  const userInfoObject = storedUserInfo ? JSON.parse(storedUserInfo) : null;
  return userInfoObject || null;
}
