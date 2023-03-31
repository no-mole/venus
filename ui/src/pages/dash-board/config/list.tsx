import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import React, { useRef, useState } from 'react';
import styles from './index.less';
import { Modal, Tooltip } from 'antd';
import DiffPanel from './diff';
import EditOrViewCode from './editOrViewCode';
import { getWatchList } from './service';
import { useLocation } from '@umijs/max';

const TableList: React.FC<unknown> = () => {
  const actionRef = useRef<ActionType>();
  const { search } = useLocation();
  let searchParams = new URLSearchParams(search);
  const namespace = searchParams.get('namespace');
  const key = searchParams.get('key');
  const alias = searchParams.get('alias');

  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '节点信息',
      dataIndex: 'node_id',
      hideInSearch: true,
      render: (_, record) => (
        <Tooltip title={record?.node_addr}>{record?.node_id}</Tooltip>
      ),
      //
    },
    {
      title: '客户端认证账户', //
      hideInSearch: true,
      dataIndex: ['client_info', 'register_access_key'],
      hideInForm: true,
    },
    {
      title: '客户端HOST',
      hideInSearch: true,
      dataIndex: ['client_info', 'register_host'],
      hideInForm: true,
    },
    {
      title: '客户端IP',
      hideInSearch: true,
      dataIndex: ['client_info', 'register_ip'],
      hideInForm: true,
    },
    {
      title: '注册时间',
      hideInSearch: true,
      dataIndex: ['client_info', 'register_time'],
      valueType: 'dateTime',
      hideInForm: true,
    },
    {
      title: '上次触发时间',
      hideInSearch: true,
      dataIndex: ['client_info', 'last_interaction_time'],
      valueType: 'dateTime',
      hideInForm: true,
    },
  ];

  return (
    <>
      <PageContainer
        header={{
          title: `配置项监听列表-${key}`,
        }}
      >
        <ProTable<API.UserInfo>
          headerTitle=""
          actionRef={actionRef}
          rowKey="id"
          search={false}
          request={async (params, sorter, filter) => {
            const { data, success } = await getWatchList({
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
      </PageContainer>

      {/* 测试回滚功能 */}

      {/* 查看弹层 */}
      {/* <Modal
        title={'diff'}
        visible={viewModalVisible}
        width={1200}
        footer={false}
        onCancel={() => setViewModalVisible(false)}
      >
        <EditOrViewCode codeValue={testData} />
      </Modal> */}

      {/* diff弹层 */}
      {/* <Modal
        title={'diff'}
        visible={diffModalVisible}
        width={1200}
        footer={false}
        onCancel={() => setDiffModalVisible(false)}
      >
        <DiffPanel
          oldValue={{ data: '11111' }}
          newValue={{ data: '111112' }}
        ></DiffPanel>
      </Modal> */}
    </>
  );
};
export default TableList;
