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
  const formRef = useRef<any>();
  const initData = async () => {
    let res = await getSYSconfig({});
    if (res?.code == 0 && res?.data?.oidc) {
      formRef?.current.setFieldsValue({
        ...res.data.oidc,
        oidc_status: res.data.oidc.oidc_status == 1 ? true : false,
      });
    } else {
      formRef?.current.setFieldsValue({});
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
        formRef={formRef}
        {...formItemLayout}
        layout="horizontal"
        onFinish={async (values: any) => {
          let params: any = values;
          params.oidc_status = values.oidc_status ? 1 : -1;
          let res = await postSYSconfig({ oidc: params });
          if (res?.code == 0) {
            message.success('操作成功');
            initData();
          } else {
            message.error('操作失败，请稍后再试');
          }
        }}
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
