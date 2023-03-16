import {
  PageContainer,
  ProForm,
  ProFormSwitch,
  ProFormText,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { getSYSconfig, postSYSconfig } from './service';

const formItemLayout = {
  labelCol: { span: 2 },
  wrapperCol: { span: 16 },
};

const SystemOIDCSettings: React.FC = () => {
  const [formValue, setFormValue] = useState<any>({});
  const initData = async () => {
    let res = await getSYSconfig({});
    if (res?.code == 0 && res?.data?.oidc) {
      setFormValue(res.data.odic);
    }
  };

  useEffect(() => {
    initData();
  }, []);

  return (
    <PageContainer
      header={{
        title: 'OIDC',
      }}
    >
      <ProForm<{
        name: string;
        company?: string;
        useMode?: string;
      }>
        {...formItemLayout}
        layout="horizontal"
        onFinish={async (values) => {
          console.log(values);
          let res = postSYSconfig(values);
          console.log(res);
        }}
        initialValues={formValue}
        submitter={{
          // 完全自定义整个区域
          render: (props, doms) => {
            return [
              <Button
                type="primary"
                key="submit"
                onClick={() => props.form?.submit?.()}
              >
                更新
              </Button>,
            ];
          },
        }}
      >
        <ProFormSwitch name="oidc_status" label="开启OIDC" required />
        <ProFormText
          width="md"
          name="oauth_server"
          label="OAuthServer"
          rules={[{ required: true, message: '请填写OAuthServer' }]}
        />
        <ProFormText
          width="md"
          name="client_id"
          label="ClientID"
          rules={[{ required: true, message: '请填写ClientID' }]}
        />
        <ProFormText
          width="md"
          name="client_secret"
          label="ClientSecret"
          rules={[{ required: true, message: '请填写ClientSecret' }]}
        />
        <ProFormText
          width="md"
          name="redirect_uri"
          label="RedireUri"
          rules={[{ required: true, message: '请填写RedireUri' }]}
        />
      </ProForm>
    </PageContainer>
  );
};

export default SystemOIDCSettings;
