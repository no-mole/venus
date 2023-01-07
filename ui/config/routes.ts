export const routes = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    name: '仪表盘',
    path: '/home',
    routes: [
      {
        name: 'DashBoard',
        path: '/home',
        component: '@/pages/Home',
      },
    ],
  },
  {
    name: '工作台',
    path: '/access',
    routes: [
      { name: '配置管理', path: '/access', component: '@/pages/Access' },
      { name: '服务管理', path: '/access', component: '@/pages/Access' },
      { name: 'AccessKey', path: '/access', component: '@/pages/Access' },
      { name: '命名空间', path: '/access', component: '@/pages/Access' },
    ],
  },
  {
    name: '系统管理',
    path: '/table',
    component: '@/pages/Table',
  },
  {
    name: '用户管理',
    path: '/table',
    component: '@/pages/Table',
  },
  {
    name: 'AccessKey',
    path: '/table',
    component: '@/pages/Table',
  },
  {
    name: '命名空间',
    path: '/table',
    component: '@/pages/Table',
  },
  {
    name: '集群管理',
    path: '/table',
    component: '@/pages/Table',
  },
  {
    name: '系统设置',
    path: '/table',
    component: '@/pages/Table',
  },
];
