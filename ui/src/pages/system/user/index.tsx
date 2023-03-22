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
import styles from './index.less';
import { creatNewUser, deleteUser, getUserList, resetUser } from './service';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<ActionType>();
  const uid = localStorage?.getItem('uid');

  const columns: ProDescriptionsItemProps<User.UserInfo>[] = [
    {
      title: '关键词',
      dataIndex: 'keyword',
      hideInTable: true,
    },
    {
      title: '用户名',
      dataIndex: 'name',
      hideInSearch: true,
    },
    {
      title: '邮箱',
      dataIndex: 'uid',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '创建时间',
      hideInSearch: true,
      dataIndex: 'update_time',
      hideInForm: true,
      valueType: 'date',
    },
    {
      title: '创建者',
      hideInSearch: true,
      dataIndex: 'updater',
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
              console.log('record', record);
              history.push({
                pathname: `/system/user/detail?uid=${record?.uid}`,
              });
            }}
            rel="noopener noreferrer"
            style={{ marginRight: 8 }}
          >
            查看
          </a>
          <Popconfirm
            placement="topLeft"
            title={'确认删除吗'}
            onConfirm={() => {
              handleRemove(record);
            }}
            okText="删除"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          <TableDropdown
            key="actionGroup"
            onSelect={async (e) => {
              if (e == 'reset-pasword') {
                try {
                  await resetUser({
                    uid: record?.uid,
                  });
                  message.success('重置成功');
                  actionRef.current?.reload(); // 表格刷新
                  return true;
                } catch (error) {
                  message.error('重置失败，请重试');
                  return false;
                }
              }
            }}
            menus={[
              { key: 'reset-pasword', name: '重置密码' },
              { key: 'disable-login', name: '禁止登录' },
            ]}
          />
        </>
      ),
    },
  ];

  /**
   *  删除节点
   * @param selectedRows
   */
  const handleRemove = async (selectedRows: CONFIG.TableColumns) => {
    const hide = message.loading('正在删除');
    if (!selectedRows) return true;
    try {
      await deleteUser({
        uid: selectedRows?.uid,
      });
      hide();
      message.success('删除成功');
      actionRef.current?.reload(); // 表格刷新
      return true;
    } catch (error) {
      hide();
      message.error('删除失败，请重试');
      return false;
    }
  };

  return (
    <PageContainer
      header={{
        title: '用户管理',
      }}
    >
      <ProTable<User.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey={(record: any) => record?.namespace_uid}
        search={{
          labelWidth: 60,
        }}
        toolBarRender={() => [
          <Button
            key="new"
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
          const { data, success } = await getUserList({
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
            const success = await creatNewUser({
              ...value,
              role:
                value?.role === '普通用户'
                  ? 'UserRoleMember'
                  : 'UserRoleAdministrator',
            });
            if (success) {
              message.success('成功');
              handleUpdateModalVisible(false);
              setFormValues('');
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
          onCancel={() => {
            handleUpdateModalVisible(false);
            setFormValues('');
          }}
          updateModalVisible={updateModalVisible}
          values={formValues}
        />
      }
    </PageContainer>
  );
};

export default TableList;
