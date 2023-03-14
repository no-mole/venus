// 运行时配置
import { RunTimeLayoutConfig } from '@umijs/max';
import React from 'react';
import Footer from './components/Footer';
import RightContent from './components/RightContent';
import { theme } from 'antd';
import { errorConfig } from './requestErrorConfig';

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://next.umijs.org/docs/api/runtime-config#getinitialstate
export async function getInitialState(): Promise<{ name: string }> {
  return { name: 'name' };
}

export const layout: RunTimeLayoutConfig = () => {
  return {
    title: '配置中心',
    headerHeight: 20,
    layout: 'mix', // mix顶部才能展示
    splitMenus: false,
    navTheme: 'light',
    logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
    menu: {
      locale: false,
    },
    // 默认布局调整
    rightContentRender: () => <RightContent />,
    footerRender: () => <Footer />,
    menuHeaderRender: undefined,
  };
};

export const antd = (memo: any) => {
  memo.theme ||= {};
  memo.theme.algorithm = theme.darkAlgorithm;
  return memo;
};

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request = {
  ...errorConfig,
};
