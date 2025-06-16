import axios from 'axios';
import { getAccessToken } from '../service/auth';

export const baseUrl = (window as any).global_config.BASE_URL
export const ossUrl = (window as any).global_config.OSS_URL

console.log( import.meta.env.VITE_APP_ENV)

const instance = axios.create({
  baseURL: `${baseUrl}/api/`,
  timeout: 120000,
  withCredentials: false,
});

instance.interceptors.request.use(
  function (config) {
    // 请求携带token
    const token = getAccessToken();
    if (token) {
      config.headers.Authorization = token;
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  },
);

instance.interceptors.response.use(
  function (response) {
    const { returnCode, returnDesc, data } = response.data;

    if (returnCode !== 17000) {
      console.error(response.data)

      if (returnDesc.includes(401)) {
        console.error('认证失败:', response.data);

        // 检查当前路径是否已经是登录页，避免重定向循环
        const currentPath = window.location.hash;
        if (currentPath !== '#/login') {
          // 清除登录状态
          localStorage.removeItem('isLoggedIn');
          localStorage.removeItem('accessToken');
          localStorage.removeItem('userInfo');

          window.location.href = '#/login';
        }

        // 在认证失败的情况下，返回一个特殊标记的错误对象
        const authError = new Error('认证失败，请重新登录');
        (authError as any).isAuthError = true;
        return Promise.reject(authError);
      }

      return response.data;
    }

    return data || true;
  },
  function (error) {

    console.log(error);
    const { response } = error;
    if (!response) {
      return Promise.reject(error);
    }

    return Promise.reject(error);
  },
);

export default instance;
