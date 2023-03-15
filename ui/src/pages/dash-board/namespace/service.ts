import { request } from '@umijs/max';

//获取命名空间下用户列表
export async function getList(params: any) {
    return request(`/api/v1/namespace/${params.namespace}/user`, {
      method: 'GET',
    });
}

//新增用户
export async function postAddUser(params: any) {
    return request(`/api/v1/namespace/${params.namespace}/user/${params.uid}`, {
      method: 'POST',
      data: params,
    });
}

//删除用户
export async function postDeleteUser(params: any) {
    return request(`/api/v1/namespace/${params.namespace}/user/${params.uid}`, {
      method: 'DELETE',
    });
}

//获取用户列表
export async function getUserList() {
  return request(`/api/v1/user`, {
    method: 'GET',
  });
}