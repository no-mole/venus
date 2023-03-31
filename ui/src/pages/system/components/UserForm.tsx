import {
  ModalForm,
  ProForm,
  ProFormRadio,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React from 'react';
import { creatNewUser } from '../user/service';

export interface FormValueType extends Partial<API.UserInfo> {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
  namespace_alias?: string;
}

export interface UpdateFormProps {
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<void>;
  updateModalVisible: boolean;
  values: Partial<API.UserInfo>;
  formType: string;
}

export interface userProps {
  name: string;
  uid: string;
  password: string;
}

const UserForm: React.FC<UpdateFormProps> = (props) => {
  const finish = async (values: userProps) => {
    let res = await creatNewUser({ ...values });
    if (res.code === 0) {
      message.success('新增成功');
      // props.updateModalVisible(false);
    } else {
      message.error('添加失败');
    }
  };

  return (
    <ModalForm
      initialValues={props.values}
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
      width={440}
    >
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="name"
          label="用户名"
          rules={[{ required: true, message: '请输入用户名！' }]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormText
          width="xl"
          name="uid"
          label="邮箱"
          rules={[
            {
              type: 'email',
              message: '请输入正确的电子邮箱',
            },
            { required: true, message: '请输入邮箱！' },
          ]}
        />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormText.Password width="xl" name="password" label="密码" />
      </ProForm.Group>
      <ProForm.Group>
        <ProFormRadio.Group
          name="role"
          label="权限"
          options={['普通成员', '管理员']}
          rules={[{ required: true, message: '请选择权限！' }]}
        />
      </ProForm.Group>
    </ModalForm>
  );
};

export default UserForm;
