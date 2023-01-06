export const routes = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    name: '首页',
    path: '/home',
    component: '@/pages/Home',
  },
  {
    name: '权限演示',
    path: '/access',
    component: '@/pages/Access',
  },
  {
    name: ' CRUD 示例',
    path: '/table',
    component: '@/pages/Table',
  },
];
