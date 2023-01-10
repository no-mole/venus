import {
  ModalForm,
  ProForm,
  ProFormDateTimePicker,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  StepsForm,
} from '@ant-design/pro-components';
import { message, Modal } from 'antd';
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
}

const UpdateForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
    title="新建表单"
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
    width={640}
  >
    <ProForm.Group>
      <ProFormText
        width="md"
        name="name"
        label="规则名称"
        rules={[{ required: true, message: '请输入规则名称！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormTextArea
        name="desc"
        width="md"
        label="规则描述"
        placeholder="请输入至少五个字符"
        rules={[
          { required: true, message: '请输入至少五个字符的规则描述！', min: 5 },
        ]}
      />
    </ProForm.Group>
  </ModalForm>
);

export default UpdateForm;
