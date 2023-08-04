import { request } from '@umijs/max';

// 登录
export async function login(params: any) {
  return request(`/api/v1/login/`, {
    method: 'POST',
    data: params,
  });
}

// OIDC登录
export async function oidclogin() {
  return request(`/api/v1/oidc_login`, {
    method: 'get',
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
  return request(`/api/v1/change_password`, {
    method: 'put',
    data: params,
  });
}

// 判断是否OIDC登录
export async function getOIDC() {
  return request(`/api/v1/user/venus`, {
    method: 'get',
  });
}
