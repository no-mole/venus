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
  return request(`/api/v1/service/${params.namespace}/${params.name}/${params.version}`, {
    method: 'GET',
  });
}

//获取服务所有信息
export async function getDetailInfo(params: any) {
  return request(`/api/v1/service/${params.namespace}/${params.name}/${params.version}/${params.addr}`, {
    method: 'GET',
  });
}
