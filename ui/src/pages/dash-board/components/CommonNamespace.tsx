import { ProFormSelect } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import { useLocalStorageState } from 'ahooks';
import { Modal } from 'antd';
import React, { PropsWithChildren } from 'react';

const CommonNamespace: React.FC<any> = () => {
  // const { modalVisible, onCancel } = props;

  const { list, loading, select, setSelect } = useModel('useUser');
  const [message, setMessage] = useLocalStorageState(
    'use-local-storage-state-namespace',
    {
      defaultValue: select,
    },
  );

  return (
    <div style={{ marginTop: 10, marginLeft: 40, marginBottom: '-24px' }}>
      <ProFormSelect
        options={list}
        width={'xs'}
        style={{ width: 180 }}
        fieldProps={{
          value: message.value,
          onChange: (e: any, option: any) => {
            setSelect({ label: option.label, value: option.value });
            setMessage({ label: option.label, value: option.value });
          },
        }}
      />
    </div>
  );
};

export default CommonNamespace;
