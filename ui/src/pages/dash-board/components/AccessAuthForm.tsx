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

const AccessAuthForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
    visible={props.updateModalVisible}
    title={`新增用户爱奇艺对命名空间***的权限`}
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
        label="命名空间名称"
        rules={[{ required: true, message: '请输入命名空间名称！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="id"
        label="命名空间标识"
        rules={[{ required: true, message: '请输入命名空间标识！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormRadio.Group
        name="checkbox-group"
        label="权限"
        options={['只读']}
        rules={[{ required: true, message: '请选择权限！' }]}
      />
    </ProForm.Group>
  </ModalForm>
);

export default AccessAuthForm;
