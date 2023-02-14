import services from '@/services/demo';
import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { Button, message } from 'antd';
import React, { useRef, useState } from 'react';
import NameSpaceForm, { FormValueType } from '../components/NameSpaceForm';

import styles from './../config/index.less';

const { addUser, queryUserList, deleteUser, modifyUser } =
  services.UserController;

/**
 * 添加节点
 * @param fields
 */
const handleAdd = async (fields: API.UserInfo) => {
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
const handleRemove = async (selectedRows: API.UserInfo[]) => {
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
  // const history = useHistory();
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '用户名称',
      dataIndex: 'name',
      hideInSearch: true,
    },
    {
      title: '唯一标识',
      dataIndex: 'email',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '角色',
      dataIndex: 'role',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '权限',
      dataIndex: 'limit',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '创建时间',
      hideInSearch: true,
      dataIndex: 'time',
      hideInForm: true,
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record, _, action) => (
        <>
          <a
            style={{ marginRight: 8 }}
            onClick={() => {
              handleRemove();
            }}
          >
            删除
          </a>
          <a
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues(record);
              setFormType('详情');
            }}
            rel="noopener noreferrer"
            style={{ marginRight: 8 }}
          >
            修改权限
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

      {/* 更新 */}
      {
        <NameSpaceForm
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
