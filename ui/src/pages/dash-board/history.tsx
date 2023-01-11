import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message } from 'antd';
import React, { useRef, useState } from 'react';
import UpdateForm, { FormValueType } from './components/UpdateForm';
import styles from './index.less';

const { queryUserList } = services.UserController;

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<ActionType>();
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      width: 150,
      title: '客户端名称',
      dataIndex: 'name',
      tip: '名称是唯一的 key',
      hideInSearch: true,
    },
    {
      title: 'hostname',
      width: 150,
      dataIndex: 'nickName',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: 'ip',
      hideInSearch: true,
      dataIndex: 'gender',
      hideInForm: true,
    },
    {
      title: '创建监听事件',
      hideInSearch: true,
      dataIndex: 'gender',
      hideInForm: true,
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
