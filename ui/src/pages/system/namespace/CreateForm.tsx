import React, { PropsWithChildren } from 'react';
import {
  ProForm,
  ProFormText,
  ModalForm
} from '@ant-design/pro-components';

export interface FormValueType extends Partial<API.UserInfo> {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
}

interface CreateFormProps {
  title?: string,
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<void>;
  modalVisible: boolean;
}

const CreateForm: React.FC<PropsWithChildren<CreateFormProps>> = (props) => {
  const { modalVisible, onCancel, onSubmit } = props;

  return (
    <ModalForm
      visible={modalVisible}
      autoFocusFirstInput
      modalProps={{
        destroyOnClose: true,
        onCancel: () => onCancel(),
      }}
      submitTimeout={2000}
      onFinish={onSubmit}
      width={440}
    >
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="namespace_alias"
          label="空间名称"
          rules={[{ required: true, message: '请输入空间名称！' }]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="namespace_uid"
          label="唯一标识"
          rules={[{ required: true, message: '请输入唯一标识！' }]}
        />
      </ProForm.Group>
    </ModalForm>
  );
};

export default CreateForm;
