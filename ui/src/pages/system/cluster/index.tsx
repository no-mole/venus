import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useRef, useState } from 'react';
import { history } from 'umi';

const { queryUserList } =
  services.UserController;

const TableList: React.FC<unknown> = () => {
  const actionRef = useRef<ActionType>();
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '节点ID',
      dataIndex: 'name',
    },
    {
      title: '节点IP',
      dataIndex: 'nickName',
    },
    {
      title: '端口',
      dataIndex: 'nickName',
    },
    {
      title: 'HOSTNAME',
      dataIndex: 'nickName',
    },
    {
      title: '角色',
      dataIndex: 'nickName',
    },
    {
      title: '是否在线',
      dataIndex: 'nickName',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_: any, record: any) => (
        <a onClick={() => history.push({ pathname: '/system/cluster/detail' })}>查看</a>
      ),
    },
  ];

  return (
    <PageContainer
      header={{
        title: '集群管理',
      }}
    >
      <ProTable<API.UserInfo>
        rowKey="id"
        search={false}
        toolBarRender={false}
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
      />
    </PageContainer>
  );
};

export default TableList;
