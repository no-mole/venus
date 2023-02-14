import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import React, { useRef, useState } from 'react';
import styles from './index.less';
import { Modal } from 'antd';
import DiffPanel from './diff';
import EditOrViewCode from './editOrViewCode';

const { queryUserList } = services.UserController;

// 测试数据
const testData = { name: '', age: 11 };

const TableList: React.FC<unknown> = () => {
  const actionRef = useRef<ActionType>();
  const [diffModalVisible, setDiffModalVisible] = useState(false); // DIFF弹层
  const [viewModalVisible, setViewModalVisible] = useState(false); // 查看代码弹层

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
          <a
            rel="noopener noreferrer"
            style={{ marginRight: 8 }}
            onClick={() => {
              setTimeout(() => {
                setViewModalVisible(true);
              }, 100);

              // history.push({ pathname: `/dash-board/diff` });
            }}
          >
            查看
          </a>
          <a
            style={{ marginRight: 8 }}
            onClick={() => {
              setDiffModalVisible(true);
              // history.push({ pathname: `/dash-board/diff` });
            }}
          >
            DIFF
          </a>
          <a style={{ marginRight: 8 }}>回滚</a>
        </>
      ),
    },
  ];

  return (
    <>
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

      {/* 测试回滚功能 */}

      {/* 查看弹层 */}
      <Modal
        title={'diff'}
        visible={viewModalVisible}
        width={1200}
        footer={false}
        onCancel={() => setViewModalVisible(false)}
      >
        <EditOrViewCode codeValue={testData} />
      </Modal>

      {/* diff弹层 */}
      <Modal
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
      </Modal>
    </>
  );
};
export default TableList;
