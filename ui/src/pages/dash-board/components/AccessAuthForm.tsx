import {
  ModalForm,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useRef } from 'react';

export interface UpdateFormProps {
  onCancel: (flag?: boolean, formVals?: any) => void;
  onSubmit: (values?: any) => Promise<void>;
  getChooseOption: (values: any) => void;
  updateModalVisible: boolean;
  values: any;
  formType: string;
  namespaceoptions: [];
}

const AccessAuthForm: React.FC<UpdateFormProps> = (props) => {
  const formRef = useRef<any>();
  return (
    <ModalForm
      formRef={formRef}
      visible={props.updateModalVisible}
      layout="horizontal"
      title="添加权限"
      autoFocusFirstInput
      modalProps={{
        destroyOnClose: true,
        onCancel: () => props.onCancel(),
      }}
      submitTimeout={2000}
      onFinish={async (values) => {
        props.onSubmit();
      }}
      width={440}
    >
      <ProFormSelect
        name="namespace_alias"
        label="命名空间名称"
        showSearch
        options={props.namespaceoptions}
        rules={[{ required: true, message: '请选择命名空间名称' }]}
        fieldProps={{
          onChange: (val: string, option: any) => {
            props.getChooseOption({ value: option.value, label: option.label });
            formRef?.current.setFieldsValue({
              namespace_uid: val,
            });
          },
        }}
      />
      <ProFormText
        width="xl"
        required
        name="namespace_uid"
        label="命名空间标识"
        disabled
      />
    </ModalForm>
  );
};

export default AccessAuthForm;
