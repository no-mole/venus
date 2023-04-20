import {
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import React from 'react';
import { history } from 'umi';
import { getList } from './service'
import styles from './index.less'

const TableList: React.FC<unknown> = () => {
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '节点ID',
      dataIndex: 'id',
    },
    {
      title: '节点入口',
      dataIndex: 'address',
    },
    {
      title: '角色',
      dataIndex: 'state',
    },
    {
      title: '是否在线',
      dataIndex: 'online',
      valueEnum: {
        true: { text: '在线' },
        false: { text: '离线' },
      },
    },
    {
      title: '选举权',
      dataIndex: 'suffrage',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_: any, record: any) => (
        <a onClick={() => history.push({ pathname: `/system/cluster/detail?id=${record?.id}&nodeInfo=${record?.address}` })}>查看</a>
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
        request={async () => {
          const { data, success } = await getList();
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
    </PageContainer>
  );
};

export default TableList;
