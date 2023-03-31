// import { outLogin } from '@/services/ant-design-pro/api';
import { outLogin, upDatePassWord } from '@/pages/login/service';
import {
  LogoutOutlined,
  SettingOutlined,
  UserOutlined,
} from '@ant-design/icons';
import {
  ModalForm,
  ProForm,
  ProFormInstance,
  ProFormText,
} from '@ant-design/pro-components';
import { Avatar, Menu, message, notification, Spin } from 'antd';
// import { stringify } from 'querystring';
import type { MenuInfo } from 'rc-menu/lib/interface';
import React, { useCallback, useEffect, useRef, useState } from 'react';
import { history, useModel } from 'umi';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';

export type GlobalHeaderRightProps = {
  menu?: boolean;
};

export type UserInfo = {
  uid?: string | null;
  name?: string | null;
  role?: string | null;
  access_token?: string | null;
  token_type?: string | null;
};
/**
 * 退出登录，并且将当前的 url 保存
 */
const loginOut = async () => {
  await outLogin();
  if (window.location.pathname !== '/login') {
    history.replace({
      pathname: '/login',
    });
  }
};

const AvatarDropdown: React.FC<GlobalHeaderRightProps> = ({ menu }) => {
  const userinfo: any = localStorage.getItem('userinfo');
  // let username: string = '',
  //   uid: string = '';
  const [username, setUsername] = useState('');
  const [uid, setUid] = useState('');

  const { initialState, setInitialState } = useModel('@@initialState');
  const [passWordVisible, setPassWordVisible] = useState(false);
  const formPassRef = useRef<ProFormInstance>();
  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

  useEffect(() => {
    if (userinfo) {
      const info = JSON.parse(userinfo);
      setUsername(info?.name);
      setUid(info?.uid);

      // 判断是否修改过密码，如未修改过提示用户修改密码
      if (
        info?.change_password_status !== 1 ||
        info?.change_password_status == 0
      ) {
        message.warning('您还没有修改密码，请及时修改密码~');
        setPassWordVisible(true);
      }
    } else {
      history.push('/login');
    }
  }, [userinfo]);

  // 菜单点击事件
  const onMenuClick = useCallback(
    (event: MenuInfo) => {
      const { key } = event;
      if (key === 'logout') {
        setInitialState({});
        localStorage.removeItem('use-local-storage-state-namespace');
        localStorage.removeItem('userinfo');
        loginOut();
        history.push(`/login`);
        return;
      } else if (key === 'setpassword') {
        setPassWordVisible(true);
      }
    },
    [setInitialState],
  );

  // 密码校验
  const validatePassWord = (rule: any, value: any, callback: any) => {
    if (!value || value.length < 4 || value.length > 20) {
      return Promise.reject(
        '密码为6-20位，可以是字母、数字、特殊字符或它们的组合',
      );
    } else {
      return Promise.resolve();
    }
  };

  const loading = (
    <span className={`${styles.action} ${styles.account}`}>
      <Spin
        size="small"
        style={{
          marginLeft: 8,
          marginRight: 8,
        }}
      />
    </span>
  );

  if (!initialState) {
    return loading;
  }

  const menuHeaderDropdown = (
    <Menu className={styles.menu} selectedKeys={[]} onClick={onMenuClick}>
      {/* {menu && (
        <Menu.Item key="center">
          <UserOutlined />
          个人中心
        </Menu.Item>
      )}
      {menu && (
        <Menu.Item key="settings">
          <SettingOutlined />
          个人设置
        </Menu.Item>
      )}
      {menu && <Menu.Divider />} */}

      <Menu.Item key="setpassword">
        <SettingOutlined />
        修改密码
      </Menu.Item>

      <Menu.Item key="logout">
        <LogoutOutlined />
        退出登录
      </Menu.Item>
    </Menu>
  );
  return (
    <>
      <HeaderDropdown overlay={menuHeaderDropdown}>
        <span className={`${styles.action} ${styles.account}`}>
          {/* <Avatar
            size="default"
            className={styles.avatar}
            src={''}
            alt="avatar"
            style={{ marginRight: 5 }}
          /> */}
          <span className={`${styles.name} anticon`}>{username}</span>
        </span>
      </HeaderDropdown>

      {/* 修改密码 */}
      <ModalForm
        visible={passWordVisible}
        title="修改密码"
        width={560}
        labelAlign="right"
        autoFocusFirstInput
        {...layout}
        layout="horizontal"
        formRef={formPassRef}
        initialValues={{
          email: userinfo?.email,
        }}
        labelCol={{ span: 4 }}
        wrapperCol={{ span: 18 }}
        grid={true}
        rowProps={{
          gutter: [12, 12],
        }}
        modalProps={{
          okText: '保存',
          className: 'model',
          destroyOnClose: true,
          onCancel: () => setPassWordVisible(false),
        }}
        onFinish={async (values) => {
          if (values.new_password !== values.confirm_password) {
            message.error('新密码输入不一致，请检查');
            return;
          } else {
            let res = await upDatePassWord({
              uid: uid,
              // name: username,
              old_password: values.old_password,
              new_password: values.new_password,
            });
            if (res?.code == 0) {
              let userparse = JSON.parse(userinfo);
              userparse.change_password_status =
                res?.data?.change_password_status;
              localStorage.setItem('userinfo', JSON.stringify(userparse));
              message.success('修改成功');
              setPassWordVisible(false);
            } else {
              message.error(res?.msg || '操作失败，请稍后再试！');
            }
          }
        }}
      >
        <ProForm.Group>
          <ProFormText
            width="md"
            name="uid"
            initialValue={username}
            label="用户名："
            disabled
            fieldProps={{
              autoComplete: 'off',
            }}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText.Password
            width="md"
            name="old_password"
            label="原密码："
            fieldProps={{
              autoComplete: 'off',
            }}
            rules={[
              {
                required: true,
                message: '请填写原密码',
              },
            ]}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText.Password
            width="md"
            name="new_password"
            label="新密码："
            fieldProps={{
              autoComplete: 'off',
            }}
            rules={[
              {
                required: true,
                validator: validatePassWord,
              },
            ]}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText.Password
            fieldProps={{
              autoComplete: 'off',
            }}
            width="md"
            name="confirm_password"
            label="确认密码："
            help={'密码为6-20位，可以是字母、数字、特殊字符或它们的组合'}
            rules={[
              {
                required: true,
                validator: validatePassWord,
              },
            ]}
          />
        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default AvatarDropdown;
