import {
  ModalForm,
  ProForm,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useEffect } from 'react';

export interface UpdateFormProps {
  onCancel: (flag?: boolean, formVals?: any) => void;
  onSubmit: (values: any) => Promise<void>;
  updateModalVisible: boolean;
  values: any;
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
      <ProFormSelect.SearchSelect
        name="userQuery3"
        label="查询选择器 - options"
        fieldProps={{
          labelInValue: false,
          style: {
            minWidth: 140,
          },
        }}
        options={[
          { label: '全部', value: 'all' },
          { label: '未解决', value: 'open' },
          { label: '已解决', value: 'closed' },
          { label: '解决中', value: 'processing' },
        ]}
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
  </ModalForm>
);

export default AccessAuthForm;
