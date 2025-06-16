import request from '../utils/req';


export async function GetCaptcha(): Promise<any> {
    return request.get(`/session/captcha`);
}

export async function GetPublicKey(): Promise<any> {
    return request.get(`/session/publickey`);
}

export interface LoginParams {
    account: string;
    password: string;
    captcha_id: string;
    captcha_value: string;
}

export async function Login(params: LoginParams): Promise<any> {
    return request.post(`/session/login`, params);
}


// 登出
export function Logout(): void {
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('userInfo');
    localStorage.removeItem('accessToken');
}
