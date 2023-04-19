import { request } from '@umijs/max';

//获取服务列表
export async function getList(params: any) {
  return request(`/api/v1/service/${params.namespace}`, {
    method: 'GET',
  });
}

//获取服务版本
export async function getListVersion(params: any) {
  return request(`/api/v1/service/${params.namespace}/${params.name}`, {
    method: 'GET',
  });
}

//获取服务入口
export async function getListAddr(params: any) {
  return request(
    `/api/v1/service/${params.namespace}/${params.name}/${params.version}`,
    {
      method: 'GET',
    },
  );
}

//获取服务所有信息
export async function getDetailInfo(params: any) {
  return request(
    `/api/v1/service/${params.namespace}/${params.name}/${params.version}/${params.addr}`,
    {
      method: 'GET',
    },
  );
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
