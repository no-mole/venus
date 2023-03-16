import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useId, useRef, useState } from 'react';
import { history } from 'umi';
import UserForm from '../components/UserForm';
import { FormValueType } from '../components/UserForm';
import styles from './index.less';
import { creatNewUser, getUserList } from './service';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<ActionType>();

  // 获取接口
  const initData = async () => {
    let res = await getUserList({});
    // eslint-disable-next-line eqeqeq
    if (res?.code == 0) {
      console.log(11);
    }
  };

  const confirm = () => {
    message.info('Clicked on Yes.');
  };

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
      dataIndex: 'create_time',
      hideInForm: true,
    },
    {
      title: '创建者',
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
            onSelect={(e) => history.push({ pathname: `/dash-board/${e}` })}
            menus={[
              { key: 'reset-pasword', name: '重置密码' },
              { key: 'disable-login', name: '禁止登录' },
            ]}
          />
        </>
      ),
    },
  ];

  useEffect(() => {
    initData();
  }, []);

  return (
    <PageContainer
      header={{
        title: '用户管理',
      }}
    >
      <ProTable<User.UserInfo>
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={{
          labelWidth: 60,
        }}
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
            const success = await creatNewUser(value);
            console.log('value', value);
            if (success) {
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
