import { DownloadOutlined } from '@ant-design/icons';
import {
  ModalForm,
  ProForm,
  ProFormRadio,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { Button, message } from 'antd';
import React from 'react';

export interface UpdateFormProps {
  onCancel: (flag?: boolean, formVals?: any) => void;
  onSubmit: (values?: any) => Promise<void>;
  onDownLoad: () => void;
  updateModalVisible: boolean;
  values: any;
  formType: string;
}

const formItemLayout = {
  labelCol: { span: 4 },
  wrapperCol: { span: 18 },
};

const AccessKeyForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
    title="新建"
    {...formItemLayout}
    visible={props.updateModalVisible}
    layout="horizontal"
    autoFocusFirstInput
    modalProps={{
      destroyOnClose: true,
      onCancel: () => props.onCancel(),
    }}
    submitter={{
      render: () => {
        return [
          // ...doms,
          <Button.Group key="refs" style={{ display: 'block' }}>
            <Button
              htmlType="button"
              key="sure"
              onClick={() => {
                props.onSubmit();
              }}
            >
              确定
            </Button>
            <Button
              htmlType="button"
              key="down"
              icon={<DownloadOutlined />}
              type="primary"
              onClick={() => props.onDownLoad()}
            >
              下载
            </Button>
          </Button.Group>,
        ];
      },
    }}
    initialValues={props.values}
    submitTimeout={2000}
    width={700}
  >
    <ProFormText width="xl" name="ak" label="AccessKey" disabled />
    <ProFormText
      width="xl"
      name="password"
      label="AccessSecret"
      extra="请谨慎保存AccessKey和AccessSecret，关闭后不可再查看AccessSecret"
      disabled
    />
  </ModalForm>
);

export default AccessKeyForm;
