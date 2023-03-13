/**
 * @name umi 的路由配置
 * @description 只支持 path,component,routes,redirect,wrappers,name,icon 的配置
 * @param path  path 只支持两种占位符配置，第一种是动态参数 :id 的形式，第二种是 * 通配符，通配符只能出现路由字符串的最后。
 * @param component 配置 location 和 path 匹配后用于渲染的 React 组件路径。可以是绝对路径，也可以是相对路径，如果是相对路径，会从 src/pages 开始找起。
 * @param routes 配置子路由，通常在需要为多个路径增加 layout 组件时使用。
 * @param redirect 配置路由跳转
 * @param wrappers 配置路由组件的包装组件，通过包装组件可以为当前的路由组件组合进更多的功能。 比如，可以用于路由级别的权限校验
 * @param name 配置路由的标题，默认读取国际化文件 menu.ts 中 menu.xxxx 的值，如配置 name 为 login，则读取 menu.ts 中 menu.login 的取值作为标题
 * @param icon 配置路由的图标，取值参考 https://ant.design/components/icon-cn， 注意去除风格后缀和大小写，如想要配置图标为 <StepBackwardOutlined /> 则取值应为 stepBackward 或 StepBackward，如想要配置图标为 <UserOutlined /> 则取值应为 user 或者 User
 * @doc https://umijs.org/docs/guides/routes
 */
export const routes = [
  {
    path: '/login',
    component: '@/pages/login/index.tsx',
    layout: false,
    // 不展示顶栏
    headerRender: false,
    // 不展示页脚
    footerRender: false,
    // 不展示菜单
    menuRender: false,
    // 不展示菜单顶栏
    menuHeaderRender: false,
    // 权限配置，需要与 plugin-access 插件配合使用
    // access: 'canRead',
  },
  {
    path: '/',
    redirect: '/dash-board',
  },
  {
    name: '仪表盘',
    path: '/ddd',
    icon: 'FundProjectionScreenOutlined',
    component: '@/pages/Access',
  },
  {
    name: '工作台',
    path: '/dash-board',
    icon: 'ControlOutlined',
    routes: [
      {
        name: '配置管理',
        path: '/dash-board/config',
        routes: [
          {
            name: 'DashBoard',
            path: '/dash-board/config',
            component: '@/pages/dash-board/config/index',
            hideInMenu: true,
            icon: 'HomeOutlined',
          },
          {
            name: '查看历史',
            path: '/dash-board/config/history',
            component: '@/pages/dash-board/config/history',
            icon: 'HomeOutlined',
            hideInMenu: true,
          },
          {
            name: '监听列表',
            path: '/dash-board/config/list',
            component: '@/pages/dash-board/config/list',
            icon: 'HomeOutlined',
            hideInMenu: true,
          },
          {
            name: 'diff',
            path: '/dash-board/config/diff',
            component: '@/pages/dash-board/config/diff',
            icon: 'HomeOutlined',
            hideInMenu: true,
            hideInBreadcrumb: true,
          },
        ],
      },
      {
        name: '服务管理',
        path: '/dash-board/service',
        component: '@/pages/dash-board/config/list',
      },
      {
        name: 'AccessKey',
        path: '/dash-board/accesskey',
        routes: [
          {
            name: 'AccessKey列表',
            component: '@/pages/dash-board/accesskey/index',
            path: '/dash-board/accesskey',
            hideInMenu: true,
            icon: 'HomeOutlined',
          },
          {
            name: '查看详情',
            path: '/dash-board/accesskey/detail',
            component: '@/pages/dash-board/accesskey/detail',
            hideInMenu: true,
            icon: 'HomeOutlined',
          },
        ],
      },
      {
        name: '命名空间',
        path: '/dash-board/namespace',
        component: '@/pages/dash-board/namespace/index',
      },
    ],
  },
  {
    name: '系统管理',
    path: '/system',
    icon: 'SettingOutlined',
    routes: [
      {
        name: '用户管理',
        path: '/system/user',
        routes: [
          {
            name: 'DashBoard',
            path: '/system/user',
            component: '@/pages/system/user/index',
            hideInMenu: true,
            icon: 'HomeOutlined',
          },
        ],
      },
      {
        name: 'AccessKey',
        path: '/system/accesskey',
        component: '@/pages/Table',
      },
      {
        name: '命名空间',
        path: '/system/namespace',
        component: '@/pages/Table',
      },
      {
        name: '集群管理',
        path: '/system/jiqun', // 名称待定
        component: '@/pages/Table',
      },
      {
        name: '系统设置',
        path: '/system/config',
        component: '@/pages/Table',
      },
    ],
  },
];
