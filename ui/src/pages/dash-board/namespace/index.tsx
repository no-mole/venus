import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useRef, useState } from 'react';
import NameSpaceForm from './NameSpaceForm';

import styles from './../config/index.less';
import { getList, postAddUser, postDeleteUser } from './service'

const namespace = 'comos';
const uid = localStorage.getItem('uid');

/**
 * 添加节点
 * @param fields
 */
const handleAddAndUpdate = async (fields: any) => {
  const hide = message.loading('正在添加');
  console.log(fields);

  try {
    await postAddUser({ ...fields, uid, namespace: fields?.namespace_en });
    hide();
    message.success('添加成功');
    return true;
  } catch (error) {
    hide();
    message.error('添加失败请重试！');
    return false;
  }
};

/**
 *  删除节点
 * @param selectedRows
 */
const handleRemove = async (record: any) => {
  const hide = message.loading('正在删除');
  try {
    await postDeleteUser({
      uid: record?.uid,
      namespace: record?.namespace,
    });
    hide();
    message.success('删除成功');
    return true;
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑
  const actionRef = useRef<ActionType>();
  // const history = useHistory();

  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '用户名称',
      dataIndex: 'user_name',
    },
    {
      title: '用户邮箱',
      dataIndex: 'uid',
      valueType: 'text',
    },
    {
      title: '角色',
      dataIndex: 'role',
      valueType: 'text',
      valueEnum: {
        'wr': { text: '空间管理员' },
        'r': { text: '只读成员' },
      },
    },
    {
      title: '创建时间',
      dataIndex: 'update_time',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record: any) => (
        <>
          <Popconfirm
            placement="topLeft"
            title={`删除用户${record?.user_name}对命名空间${record?.namespace}的访问授权？`}
            onConfirm={() => handleRemove(record)}
            okText="确定"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          <a
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues(record);
              setFormType('修改');
            }}
            rel="noopener noreferrer"
            style={{ marginRight: 8 }}
          >
            修改授权
          </a>
        </>
      ),
    },
  ];

  return (
    <PageContainer
      header={{
        title: '命名空间用户权限管理',
      }}
    >
      <ProTable<API.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={false}
        options={false}
        toolBarRender={() => [
          <Button
            key="1"
            type="primary"
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues('');
              setFormType('新建');
            }}
          >
            新建
          </Button>,
        ]}
        request={async (params, sorter, filter) => {
          const { data, success } = await getList({
            ...params,
            // FIXME: remove @ts-ignore
            // @ts-ignore
            sorter,
            filter,
            namespace
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

      {/* 新增/修改 */}
      {
        <NameSpaceForm
          formType={formType}
          onSubmit={async (value) => {
            let success = await handleAddAndUpdate(value);
            if (success) {
              handleUpdateModalVisible(false);
              setFormValues({});
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
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
