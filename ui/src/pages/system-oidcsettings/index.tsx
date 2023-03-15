import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProForm,
  ProFormSwitch,
  ProFormText,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { history, useLocation } from 'umi';
import styles from './../config/index.less';

const formItemLayout = {
  labelCol: { span: 2 },
  wrapperCol: { span: 16 },
};

const SystemOIDCSettings: React.FC<unknown> = () => {
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
        onFinish={async (values) => {}}
        params={{}}
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
        <ProFormSwitch name="switch" label="开启OIDC" required />
        <ProFormText
          width="md"
          name="name"
          label="OAuthServer"
          rules={[{ required: true, message: '请填写OAuthServer' }]}
        />
        <ProFormText
          width="md"
          name="name"
          label="ClientID"
          rules={[{ required: true, message: '请填写ClientID' }]}
        />
        <ProFormText
          width="md"
          name="company"
          label="ClientSecret"
          rules={[{ required: true, message: '请填写ClientSecret' }]}
        />
        <ProFormText
          width="md"
          name="company"
          label="RedireUri"
          rules={[{ required: true, message: '请填写RedireUri' }]}
        />
      </ProForm>
    </PageContainer>
  );
};

export default SystemOIDCSettings;
