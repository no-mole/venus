import { request } from '@umijs/max';

// 登录
export async function login(params: any) {
  return request(`/api/v1/login/`, {
    method: 'POST',
    data: params,
  });
}

// 退出登录
export async function outLogin() {
  return request(`/api/v1/logout/`, {
    method: 'delete',
  });
}

// 修改密码

export async function upDatePassWord(params: any) {
  return request(`/api/v1/user/${params?.uid}`, {
    method: 'put',
    data: params,
  });
}
