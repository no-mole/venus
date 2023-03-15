import { request } from '@umijs/max';

// 登录
export async function queryConfigList(params: any) {
  return request(`/api/v1/kv/${params.namespace}`, {
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

// 删除配置
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

// 获取namespace 列表
export async function getCommonNamespace(params: any) {
  return request(`/api/v1/user/${params.uid}/namespace`, {
    method: 'get',
    params,
  });
}
