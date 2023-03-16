import {
  ModalForm,
  ProForm,
  ProFormRadio,
  ProFormText,
  ProFormSelect,
} from '@ant-design/pro-components';
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
  values: any;
  userList: any;
  formType: string;
  namespace_alias: string;
  namespace_uid: string;
}

const NameSpaceForm: React.FC<UpdateFormProps> = (props) => {
  return (
    <ModalForm
      title={props?.formType}
      visible={props.updateModalVisible}
      autoFocusFirstInput
      modalProps={{
        destroyOnClose: true,
        onCancel: () => props.onCancel(),
      }}
      submitTimeout={2000}
      onFinish={props?.onSubmit}
      width={440}
      initialValues={{
        namespace_alias: props?.namespace_alias,
        namespace_uid: props?.namespace_uid,
        uid: props?.values?.uid,
        role: props?.values?.role
      }}
    >
      <ProForm.Group>
        <ProFormSelect
          width="xl"
          name="uid"
          label="用户名称"
          fieldProps={{
            fieldNames: {
              label: 'name',
              value: 'uid',
            },
          }}
          options={props?.userList}
          rules={[{ required: true, message: '请输入用户名称！' }]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="namespace_alias"
          label="命名空间名称"
          disabled
          rules={[{ required: true, message: '请输入命名空间名称！' }]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="namespace_uid"
          label="命名空间标识"
          disabled
          rules={[{ required: true, message: '命名空间标识！' }]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormRadio.Group
          name="role"
          label="角色"
          options={[
            {
              label: '只读成员',
              value: 'r',
            },
            {
              label: '空间管理员',
              value: 'wr',
            },
          ]}
          rules={[{ required: true, message: '请选择角色！' }]}
        />
      </ProForm.Group>
    </ModalForm>
  )
};

export default NameSpaceForm;
