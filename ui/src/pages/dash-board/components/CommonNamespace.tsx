import { ProFormSelect } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import { useLocalStorageState } from 'ahooks';
import React, { useEffect } from 'react';

const CommonNamespace: React.FC<any> = () => {
  const { list, select, setSelect } = useModel('useUser');
  let namespace = localStorage.getItem('use-local-storage-state-namespace');
  if (!namespace || namespace === '{}') {
    localStorage.setItem(
      'use-local-storage-state-namespace',
      JSON.stringify({
        label: select?.namespace_alias,
        value: select?.namespace_uid,
      }),
    );
  }

  const [message, setMessage] = useLocalStorageState(
    'use-local-storage-state-namespace',
    {
      defaultValue: {
        label: select?.namespace_alias,
        value: select?.namespace_uid,
      },
    },
  );

  useEffect(() => {
    // console.log('select', select);
  }, [select]);

  return (
    <div style={{ marginTop: 10, marginLeft: 40, marginBottom: '-24px' }}>
      <ProFormSelect
        allowClear={false}
        options={list}
        width={'xs'}
        style={{ width: 180 }}
        fieldProps={{
          fieldNames: {
            label: 'namespace_alias',
            value: 'namespace_uid',
          },
          value: message.value,
          onChange: (e: any, option: any) => {
            console.log('option', option);
            setSelect({ label: option.label, value: option.value });
            setMessage({ label: option.label, value: option.value });
          },
        }}
      />
    </div>
  );
};

export default CommonNamespace;
