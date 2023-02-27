import { request } from '@umijs/max';

// 登录
export async function login(params: any) {
  return request(`/api/v1/user/login/${params.uid}`, {
    method: 'POST',
    data: params,
  });
}
