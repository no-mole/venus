import { request } from '@umijs/max';

// 登录
export async function queryConfigList(params: any) {
  return request(`/api/v1/kv/${params.namespace}`, {
    method: 'get',
    params,
  });
}

export async function addUser(params: any) {
  return request(`/api/v1/kv/${params?.uid}`, {
    method: 'get',
    data: params,
  });
}

// 删除配置
export async function deleteConfig(params: any) {
  return request(`/api/v1/kv/${params?.namespace}/${params?.key}`, {
    method: 'delete',
    data: params,
  });
}

// 新增、修改配置
export async function modifyConfig(params: any) {
  return request(`/api/v1/kv/${params?.namespace}/${params?.key}`, {
    method: 'put',
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

// 获取配置历史列表
export async function getHistoryList(params: any) {
  return request(`/api/v1/kv/history/${params?.namespace}/${params?.key}`, {
    method: 'get',
    params,
  });
}

// 获取配置历史列表详情
export async function getHistoryDetail(params: any) {
  return request(
    `/api/v1/kv/history/${params?.namespace}/${params?.key}/${params?.version}`,
    {
      method: 'get',
      params,
    },
  );
}
