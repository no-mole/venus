import {
  ActionType,
  ModalForm,
  ProForm,
  ProFormInstance,
  ProFormRadio,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useEffect, useLayoutEffect, useRef, useState } from 'react';

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

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
  const formRef = useRef<ProFormInstance>();

  return (
    <ModalForm
      initialValues={props.values}
      formRef={formRef}
      title={`配置${props.formType}`}
      open={props.updateModalVisible}
      autoFocusFirstInput
      modalProps={{
        destroyOnClose: true,
        onCancel: () => props.onCancel(),
      }}
      submitTimeout={2000}
      onFinish={async (values) => {
        props.onSubmit(values);
      }}
      width={640}
      disabled={props.formType === '详情'}
    >
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="alias"
          label="配置名称"
          rules={[{ required: true, message: '请输入配置名称！' }]}
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
        <ProFormText width="xl" name="description" label="描述" />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormRadio.Group
          name="data_type"
          label="数据类型"
          options={['text', 'json', 'yaml', 'toml', 'properties', 'ini']}
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
};

export default UpdateForm;
