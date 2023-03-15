import { request } from '@umijs/max';

//获取命名空间下用户列表
export async function getList() {
    return request(`/api/v1/namespace`, {
      method: 'GET',
    });
}

//新增用户
export async function postAddNamespace(params: any) {
    return request(`/api/v1/namespace/${params.namespace_uid}`, {
      method: 'POST',
      data: params,
    });
}

//删除用户
export async function postDeleteUser(params: any) {
    return request(`/api/v1/namespace/${params.namespace}`, {
      method: 'DELETE',
    });
}