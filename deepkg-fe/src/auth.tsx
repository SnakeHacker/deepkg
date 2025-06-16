import { Navigate, useLocation } from 'react-router-dom';
import React from 'react';
import { isAuthenticated, setAccessToken, getCurrentUser } from './service/auth';

export const AuthRoute = ({ children }: { children: React.ReactNode }) => {
  const location = useLocation();

  const userInfo = getCurrentUser();
  if (userInfo?.accessToken) {
    setAccessToken(userInfo?.accessToken);
  }

  if (!isAuthenticated()) {
    // 如果用户未登录，重定向到登录页面
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
};
