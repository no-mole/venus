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

const UpdateForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
    title={`配置${props.formType}`}
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
        width="xl"
        name="name"
        label="namespace"
        rules={[{ required: true, message: '请输入配置名称名称！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="key"
        label="唯一标识"
        rules={[{ required: true, message: '请输入唯一标识！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText width="xl" name="desc" label="描述" />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormRadio.Group
        name="data_type"
        label="数据类型"
        options={['TEXT', 'JSON', 'YAML', 'TOML', 'PROPERTIES', 'INI']}
        rules={[{ required: true, message: '请选择数据类型！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="version"
        label="数据版本"
        rules={[{ required: true, message: '请输入数据版本！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormTextArea
        name="value"
        width="xl"
        label="配置内容"
        rules={[{ required: true, message: '请输入配置内容！', min: 5 }]}
      />
    </ProForm.Group>
  </ModalForm>
);

export default UpdateForm;
