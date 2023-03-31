import { request } from '@umijs/max';

//获取cluster node列表
export async function getList() {
    return request(`/api/v1/cluster`, {
      method: 'GET',
    });
}


//获取cluster node详情
export async function getDetail(params:{id:string|null}) {
  return request(`/api/v1/cluster/${params.id}`, {
    method: 'GET',
  });
}