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
import EditOrViewCode from '../config/editOrViewCode';

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

const UpdateForm: React.FC<UpdateFormProps> = (props: any) => {
  const formRef = useRef<ProFormInstance>();
  const [state, setState] = useState('');
  const [dataType, setDataType] = useState('');

  useEffect(() => {
    setDataType(formRef?.current?.getFieldValue(type));
  }, []);

  const changeCode = (editor, changeObj: any, code: any) => {
    if (!code) return;
    // 获取 CodeMirror.doc.getValue()
    // 赋值 CodeMirror.doc.setValue(value) // 会触发 onChange 事件，小心进入无线递归。
    // console.log('code', editor, changeObj, code);
    setState(JSON.stringify(code));
  };
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
        let new_value = { ...values, value: state };
        props.onSubmit(new_value);
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
          fieldProps={{
            onChange: (e: any) => {
              formRef?.current.setFieldsValue({
                type: e.target.value,
              });
              setDataType(e.target.value);

              console.log(formRef?.current?.getFieldValue('type'));
              // return val;
            },
          }}
        />
      </ProForm.Group>
      {props.formType === '详情' ? (
        <ProForm.Group>
          <ProFormText
            disabled
            width="xl"
            name="version"
            label="数据版本"
            rules={[{ required: true, message: '请输入数据版本！' }]}
          />
        </ProForm.Group>
      ) : (
        ''
      )}
      {/* <ProForm.Group>
        <ProFormTextArea
          name="value"
          width="xl"
          label="配置内容"
          rules={[{ required: true, message: '请输入配置内容！', min: 5 }]}
        />
      </ProForm.Group> */}
      <p style={{ marginBottom: 20 }}>
        <em style={{ color: '#ff4d4f', marginInlineEnd: 4 }}>*</em>配置内容:
      </p>
      <EditOrViewCode
        changeCode={changeCode}
        codeValue={props?.values?.value}
        type={dataType}
        setDataType={setDataType}
        formType={props.formType}
      />
    </ModalForm>
  );
};

export default UpdateForm;
