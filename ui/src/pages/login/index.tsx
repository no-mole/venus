import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Checkbox, Form, Input, message, Divider } from 'antd';
import { useEffect } from 'react';
import Styles from './index.module.less';
import { login } from './service';
import { history } from 'umi';
import Avatar from '../../assets/honeycomb.png';

export default () => {
  const [form] = Form.useForm();

  useEffect(() => {
    const userinfo = localStorage.getItem('userinfo');
    if (userinfo) {
      const info = JSON.parse(userinfo);
      form.setFieldsValue({
        name: info?.name,
        uid: info?.uid,
        password: info?.password,
      });
    } else {
      history.push('/login');
    }
  }, []);

  const onFinish = async (values: any) => {
    const res = await login({
      uid: values?.uid,
      password: values?.password,
    });
    if (res?.code !== 0) {
      message.error('登录失败，请稍后重试！');
    } else {
      if (values?.remember) {
        localStorage.setItem(
          'userinfo',
          JSON.stringify({
            name: res?.data?.name,
            uid: values?.uid,
            password: values?.password,
            token: res?.data?.token_type + ' ' + res?.data?.access_token,
            role: res?.data?.role,
          }),
        );
      } else {
        localStorage.setItem('userinfo', '');
      }
      localStorage.setItem('uid', res?.data?.uid);
      history.push({
        pathname: '/dash-board/config',
      });
    }
  };

  return (
    <div className={Styles.container}>
      <div className={Styles.title}>
        <img src={Avatar} alt="" />
        <Divider
          type="vertical"
          style={{ height: '47%', backgroundColor: '#959496', width: '2px' }}
        />
        <span>MOLE CLOUD</span>
      </div>
      <Form
        form={form}
        name="normal_login"
        className={Styles.loginForm}
        initialValues={{ remember: true }}
        onFinish={onFinish}
      >
        <Form.Item
          name="uid"
          rules={[{ required: true, message: 'Please input your Username!' }]}
        >
          <Input
            prefix={<UserOutlined className={Styles['site-form-item-icon']} />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: 'Please input your Password!' }]}
        >
          <Input.Password
            prefix={<LockOutlined className={Styles['site-form-item-icon']} />}
            type="password"
            placeholder="Password"
          />
        </Form.Item>
        <Form.Item>
          <Form.Item name="remember" valuePropName="checked" noStyle>
            <Checkbox>
              <span style={{ color: '#fff' }}>记住密码</span>
            </Checkbox>
          </Form.Item>
        </Form.Item>
        <Form.Item>
          <Button
            type="primary"
            htmlType="submit"
            className={Styles['login-form-button']}
          >
            登录
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
