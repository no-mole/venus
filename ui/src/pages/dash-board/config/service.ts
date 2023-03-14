import { request } from '@umijs/max';

// 登录
export async function queryConfigList(params: any) {
  return request(`/api/v1/kv/default`, {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    method: 'get',
    params,
  });
}

export async function addUser(params: any) {
  return request(`/api/v1/kv/default`, {
    method: 'get',
    data: params,
  });
}

export async function deleteUser(params: any) {
  return request(`/api/v1/kv/default`, {
    method: 'get',
    data: params,
  });
}

export async function modifyUser(params: any) {
  return request(`/api/v1/kv/default`, {
    method: 'get',
    data: params,
  });
}
