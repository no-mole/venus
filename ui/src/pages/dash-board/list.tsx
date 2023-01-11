import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import React, { useRef, useState } from 'react';
import styles from './index.less';

const { queryUserList } = services.UserController;

const TableList: React.FC<unknown> = () => {
  const actionRef = useRef<ActionType>();
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      width: 250,
      title: 'MD5',
      dataIndex: 'name',
      tip: '名称是唯一的 key',
      hideInSearch: true,
    },
    {
      title: '操作者',
      width: 150,
      dataIndex: 'user',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '操作者邮箱',
      hideInSearch: true,
      dataIndex: 'email',
      hideInForm: true,
    },
    {
      title: '操作时间',
      hideInSearch: true,
      dataIndex: 'gender',
      hideInForm: true,
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record, _, action) => (
        <>
          <a rel="noopener noreferrer" style={{ marginRight: 8 }}>
            查看
          </a>
          <a style={{ marginRight: 8 }}>DIFF</a>
          <a style={{ marginRight: 8 }}>回滚</a>
        </>
      ),
    },
  ];

  return (
    <PageContainer
      header={{
        title: '配置项监听列表-mysql',
      }}
    >
      <ProTable<API.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={false}
        request={async (params, sorter, filter) => {
          const { data, success } = await queryUserList({
            ...params,
            // FIXME: remove @ts-ignore
            // @ts-ignore
            sorter,
            filter,
          });
          return {
            data: data?.list || [],
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
    </PageContainer>
  );
};
export default TableList;
