import { defineConfig } from '@umijs/max';
// import defaultSettings from './defaultSettings';
import proxy from './proxy';
import { routes } from './routes';
import ico from './../src/assets/honeycomb.png';

const { NODE_ENV } = process.env;
let routerRoot =
  NODE_ENV === 'development' ? { publicPath: '/' } : { publicPath: '/ui/' };

export default defineConfig({
  ...routerRoot,
  hash: true,
  access: {},
  antd: {},
  model: {},
  initialState: {},
  // plugins: ['@umijs/plugins/dist/initial-state', '@umijs/plugins/dist/model'],
  request: {},
  layout: {
    title: '配置中心',
    locale: false, // 默认开启，如无需菜单国际化可关闭
  },
  // 路由前缀，部署到非根目录
  base: NODE_ENV === 'development' ? '/' : '/ui/',
  // links: [{ rel: 'icon', href: '/' }],
  jsMinifier: 'terser',
  favicons: ['./../public/assets/honeycomb.png'],
  // 兼容性配置
  // targets: {
  //   ie: 11,
  // },
  routes,
  ignoreMomentLocale: true,
  proxy: proxy['dev'],
  exportStatic: {},
});
