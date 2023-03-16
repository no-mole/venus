import { request } from '@umijs/max';

// 获取系统配置
export async function getSYSconfig(params: any) {
  return request(`/api/v1/sys_config`, {
    method: 'GET',
    // params,
  });
}

// 更新系统配置
export async function postSYSconfig(params: any) {
  return request(`/api/v1/sys_config`, {
    method: 'POST',
    data: params,
  });
}
