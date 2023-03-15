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
  values: Partial<API.UserInfo>;
  formType: string;
}


const namespace_cn = 'comos';
const namespace_en = 'comos';

const NameSpaceForm: React.FC<UpdateFormProps> = (props) => (
  <ModalForm
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
      namespace_cn,
      namespace_en
    }}
  >
    <ProForm.Group>
      <ProFormSelect
        width="xl"
        name="user_name"
        label="用户名称"
        request={async () => [
          { label: 'user1', value: 'user1' },
          { label: 'user2', value: 'user2' },
        ]}
        rules={[{ required: true, message: '请输入用户名称！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="namespace_cn"
        label="命名空间名称"
        disabled
        rules={[{ required: true, message: '请输入命名空间名称！' }]}
      />
    </ProForm.Group>
    <ProForm.Group>
      <ProFormText
        width="xl"
        name="namespace_en"
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
);

export default NameSpaceForm;
