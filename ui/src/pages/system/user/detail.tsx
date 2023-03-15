import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useRef, useState } from 'react';
import { history } from 'umi';
import UserForm from '../components/UserForm';
import { FormValueType } from '../components/UserForm';
import styles from './index.less';
import { getUserNamespace } from './service';

import { parse } from 'query-string';

const { addUser, deleteUser, modifyUser } = services.UserController;

/**
 * 添加节点
 * @param fields
 */
const handleAdd = async (fields: User.UserInfo) => {
  const hide = message.loading('正在添加');
  try {
    await addUser({ ...fields });
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
 * 更新节点
 * @param fields
 */
const handleUpdate = async (fields: FormValueType) => {
  const hide = message.loading('正在配置');
  try {
    await modifyUser(
      {
        userId: fields.id || '',
      },
      {
        name: fields.name || '',
        nickName: fields.nickName || '',
        email: fields.email || '',
      },
    );
    hide();

    message.success('配置成功');
    return true;
  } catch (error) {
    hide();
    message.error('配置失败请重试！');
    return false;
  }
};

/**
 *  删除节点
 * @param selectedRows
 */
const handleRemove = async (selectedRows: User.UserInfo[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    await deleteUser({
      userId: selectedRows.find((row) => row.id)?.id || '',
    });
    hide();
    message.success('删除成功，即将刷新');
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
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<ActionType>();
  const query = parse(history.location.search);
  const { uid } = query;
  const columns: ProDescriptionsItemProps<User.UserInfo>[] = [
    {
      title: '命名空间名称',
      dataIndex: 'name',
      hideInSearch: true,
    },
    {
      title: '命名空间标识',
      dataIndex: 'uid',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '权限',
      hideInSearch: true,
      dataIndex: 'create_time',
      hideInForm: true,
    },
    {
      title: '创建时间',
      hideInSearch: true,
      dataIndex: 'creator',
      hideInForm: true,
    },
    {
      title: '角色',
      hideInSearch: true,
      dataIndex: 'role',
      hideInForm: true,
      render: (record) => {
        return record === 'UserRoleAdministrator'
          ? '管理员'
          : record === 'UserRoleMember'
          ? '普通成员'
          : record;
      },
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record, _, action) => (
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
            修改授权
          </a>
          <Popconfirm
            placement="topLeft"
            title={'确认删除吗'}
            onConfirm={() => {
              handleRemove();
            }}
            okText="删除"
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
        title: '用户空间权限',
      }}
    >
      <ProTable<User.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={false}
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
          const { data, success } = await getUserNamespace({
            uid,
            ...params,
            // FIXME: remove @ts-ignore
            // @ts-ignore
            sorter,
            filter,
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

      {/* 更新 */}
      {
        <UserForm
          formType={formType}
          onSubmit={async (value) => {
            const success = await handleUpdate(value);
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
