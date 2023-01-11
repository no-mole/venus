import { QuestionCircleOutlined } from '@ant-design/icons';
import { Space } from 'antd';
import React from 'react';
import { useModel } from 'umi';
import Avatar from './AvatarDropdown';
import styles from './index.less';
export type SiderTheme = 'light' | 'dark';

const GlobalHeaderRight: React.FC = () => {
  // const { initialState } = useModel('@@initialState');

  // // @ts-ignore
  // if (!initialState || !initialState.settings) {
  //   return null;
  // }

  // // @ts-ignore
  // const { navTheme, layout } = initialState.settings;
  // let className = styles.right;

  // if ((navTheme === 'dark' && layout === 'top') || layout === 'mix') {
  //   className = `${styles.right}  ${styles.dark}`;
  // }

  return (
    <Space>
      {/* <NoticeIconView /> */}
      <Avatar menu />
    </Space>
  );
};

export default GlobalHeaderRight;
