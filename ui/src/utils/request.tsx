import type { RequestConfig } from '@umijs/max';

export const request: RequestConfig = {
  timeout: 2000,
  errorConfig: {
    // 错误抛出
    errorThrower: (res) => {
      console.log(res);
    },
    // 错误接收及处理
    errorHandler: (error: any) => {
      console.log(error);
    },
  },
  // 请求拦截器
  requestInterceptors: [
    (url, options) => {
      const headers: any = { ...options.headers };
      const token = localStorage.getItem('access_token');
      if (token) {
        headers['Authorization'] = token;
      }
      return { url, options: { ...options, headers } };
    },
  ],
  // 响应拦截器
  responseInterceptors: [
    (response) => {
      const { code: errorCode, message: errorMessage, data } = response.data;
      response.data = {
        errorCode,
        errorMessage,
        data,
        success: errorCode === 0,
      };
      return response;
    },
  ],
};

export default request;
