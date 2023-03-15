import { request } from '@umijs/max';

// 新增用户
export async function creatNewUser(params: any) {
  return request(`/api/v1/user/default`, {
    method: 'post',
    data: params,
  });
}

// 获取用户列表
export async function getUserList(params: any) {
  return request(`/api/v1/user`, {
    method: 'get',
    params,
  });
}

// uer namespace列表  /user/{ak}/namespace
export async function getUserNamespace(params: any) {
  return request(`/api/v1/user/${params.uid}/namespace`, {
    method: 'get',
    params,
  });
}
