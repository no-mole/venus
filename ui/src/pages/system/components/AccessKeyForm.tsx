import {
  ModalForm,
  ProForm,
  ProFormRadio,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React from 'react';

export interface FormValueType extends Partial<API.UserInfo> {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
}

export interface UpdateFormProps {
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<void>;
  updateModalVisible: boolean;
  values: Partial<API.UserInfo>;
  formType: string;
}

const AccessKeyForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
    visible={props.updateModalVisible}
    autoFocusFirstInput
    modalProps={{
      destroyOnClose: true,
      onCancel: () => props.onCancel(),
    }}
    submitTimeout={2000}
    onFinish={async (values) => {
      console.log(values.name);
      message.success('提交成功');
      return true;
    }}
    width={440}
  >
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="name"
        label="AccessKey"
        rules={[{ required: true, message: '请输入AccessKey！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="id"
        label="AccessSecret"
        tooltip="请谨慎保存AccessKey和AccessSecret，关闭后不可再查看AccessSecret"
        rules={[{ required: true, message: '请输入AccessSecret！' }]}
      />
    </ProForm.Group>
  </ModalForm>
);

export default AccessKeyForm;
