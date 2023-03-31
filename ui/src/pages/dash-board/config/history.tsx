import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { useLocation } from '@umijs/max';
import React, { useRef, useState } from 'react';
import UpdateForm from '../components/UpdateForm';
import styles from './index.less';
import { getHistoryList } from './service';

const TableList: React.FC<unknown> = () => {
  const { search } = useLocation();
  let searchParams = new URLSearchParams(search);
  const namespace = searchParams.get('namespace');
  const key = searchParams.get('key');
  const alias = searchParams.get('alias');
  const actionRef = useRef<ActionType>();
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '配置项名称',
      dataIndex: 'alias',
      hideInSearch: true,
    },
    {
      title: '唯一标识',
      dataIndex: 'key',
      hideInSearch: true,
    },
    {
      title: '版本',
      hideInSearch: true,
      dataIndex: 'version',
      hideInForm: true,
    },
    {
      title: '最近更新时间',
      hideInSearch: true,
      dataIndex: 'update_time',
      valueType: 'dateTime',
      hideInForm: true,
    },
    {
      title: '操作',
      render: (text, record: any) => (
        <>
          <a
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues(record);
              setFormType('详情');
            }}
            rel="noopener noreferrer"
            style={{ marginRight: 8 }}
          >
            查看
          </a>
        </>
      ),
    },
  ];

  return (
    <PageContainer
      header={{
        title: `${alias}配置历史列表`,
      }}
    >
      <ProTable<API.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={false}
        request={async (params, sorter, filter) => {
          const { data, success } = await getHistoryList({
            namespace: namespace,
            key: key,
            ...params,
            // FIXME: remove @ts-ignore
            // @ts-ignore
            sorter,
            filter,
          });
          return {
            data: data || [],
            success,
          };
        }}
        columns={columns}
        rowClassName={(record, index) => {
          let className = styles.lightRow;

          if (index % 2 === 1) className = styles.darkRow;
          return className;
        }}
      />

      {/* 更新 */}
      {
        <UpdateForm
          formType={formType}
          onSubmit={async () => {}}
          onCancel={() => {
            handleUpdateModalVisible(false);
            setFormValues({});
          }}
          updateModalVisible={updateModalVisible}
          values={formValues}
        />
      }
    </PageContainer>
  );
};

export default TableList;
