import { Settings as ProSettings } from '@ant-design/pro-layout';

type DefaultSettings = Partial<ProSettings> & {
  pwa: boolean;
  // logo: string;
};

const proSettings: DefaultSettings = {
  navTheme: 'light',
  // primaryColor: '#2A61EE',
  layout: 'mix',
  contentWidth: 'Fluid',
  fixedHeader: false,
  fixSiderbar: true,
  pwa: false,
  title: '1111配置中心',
  // headerHeight: 64,
  splitMenus: false,
};

export type { DefaultSettings };
export default proSettings;
