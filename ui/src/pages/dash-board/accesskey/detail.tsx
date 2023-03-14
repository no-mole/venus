import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { history, useLocation } from 'umi';
import AccessAuthForm from '../components/AccessAuthForm';
import styles from './../config/index.less';
import {
  getDeatilsAccessKeyList,
  deleteAccessKey,
  getNameSpaceList,
} from './service';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<any>();
  let location = useLocation();

  const handleRemove = async (obj: any) => {
    let res = await deleteAccessKey({ namespace: obj.namespace, ak: obj.ak });
    if (res?.code == 0) {
      message.success('删除成功');
      actionRef?.current.reload();
    } else {
      message.error(res?.mes || '操作失败，请稍后再试');
    }
  };

  const getSelectList = async () => {
    let res = await getNameSpaceList({});
    console.log(res);
  };

  useEffect(() => {
    getSelectList();
  }, []);

  const columns: ProDescriptionsItemProps[] = [
    {
      title: '命名空间名称',
      dataIndex: 'namespace',
    },
    {
      title: '命名空间标识',
      dataIndex: 'ak',
      hideInSearch: true,
    },
    {
      title: '权限',
      dataIndex: 'authority',
      valueType: 'text',
      hideInSearch: true,
      render: () => {
        return '只读';
      },
    },
    {
      title: '更新人',
      hideInSearch: true,
      dataIndex: 'updater',
      hideInForm: true,
    },
    {
      title: '更新时间',
      hideInSearch: true,
      dataIndex: 'update_time',
      hideInForm: true,
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record, _, action) => (
        <>
          <Popconfirm
            placement="topLeft"
            title={`确认删除${record?.namespace}对空间命名标识${record?.ak}的访问授权吗`}
            onConfirm={() => {
              handleRemove(record);
            }}
            okText="确定"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
        </>
      ),
    },
  ];

  return (
    <PageContainer
      header={{
        title: 'AccessKey空间权限',
      }}
    >
      <ProTable<API.UserInfo>
        actionRef={actionRef}
        rowKey="ak"
        search={false}
        options={false}
        headerTitle={[
          <Button
            key="keyadd"
            type="primary"
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues('');
              setFormType('添加');
            }}
          >
            添加
          </Button>,
        ]}
        request={async (params, sorter, filter) => {
          const { data, success } = await getDeatilsAccessKeyList({
            // @ts-ignore
            ak: location.state?.ak,
          });
          return {
            data: data?.items || [],
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

      {/* 添加，将某个accesskey加到某个命名空间下 */}
      <AccessAuthForm
        formType={formType}
        onSubmit={async (value) => {}}
        onCancel={() => {
          handleUpdateModalVisible(false);
          setFormValues({});
        }}
        updateModalVisible={updateModalVisible}
        values={formValues}
      />
    </PageContainer>
  );
};

export default TableList;
