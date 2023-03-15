import { request } from '@umijs/max';

// 获取工作台accesskey列表
export async function getAccessKeyList(params: any) {
  return request(`/api/v1/namespace/${params.namespace}/access_key`, {
    method: 'GET',
    // params,
  });
}

// 工作台创建accesskey
export async function postNewAccesskey(params: any) {
  return request(`/api/v1/access_key/${params.namespace}/${params.alias}`, {
    method: 'POST',
    data: params,
  });
}

// 详情查看accesskey列表
export async function getDeatilsAccessKeyList(params: any) {
  return request(`/api/v1/access_key/${params.ak}/namespace`, {
    method: 'GET',
  });
}

// 工作台及查看删除列表accesskey
export async function deleteAccessKey(params: any) {
  return request(
    `/api/v1/namespace/${params.namespace}/access_key/${params.ak}`,
    {
      method: 'DELETE',
      // params,
    },
  );
}

// 获取命名空间列表
export async function getNameSpaceList(params: any) {
  return request(`/api/v1/namespace`, {
    method: 'GET',
  });
}

// 命名空间下新增accesskey
export async function postNameSpaceAccessKey(params: any) {
  return request(
    `/api/v1/namespace/${params.namespace}/access_key/${params.ak}`,
    {
      method: 'POST',
      data: params,
    },
  );
}
