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
import UpdateForm, { FormValueType } from '../components/UpdateForm';
import styles from './index.less';
import { queryConfigList, addUser, deleteUser, modifyUser } from './service';
import { history } from 'umi';
import CommonNamespace from '../components/CommonNamespace';
import { useModel } from '@umijs/max';

// const { addUser, queryUserList, deleteUser, modifyUser } =
//   services.UserController;

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
  const namespace = localStorage.getItem('use-local-storage-state-namespace'); // 默认namespace
  console.log(namespace);
  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      width: 150,
      title: '配置项名称',
      dataIndex: 'name',
      tip: '名称是唯一的 key',
    },
    {
      title: '唯一标识',
      width: 150,
      dataIndex: 'nickName',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      valueType: 'date',
      hideInSearch: true,
    },
    {
      title: '最近更新时间',
      hideInSearch: true,
      dataIndex: 'gender',
      hideInForm: true,
      valueEnum: {
        0: { text: '男', status: 'MALE' },
        1: { text: '女', status: 'FEMALE' },
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
            查看
          </a>
          <a
            style={{ marginRight: 8 }}
            onClick={() => {
              handleUpdateModalVisible(true);
              setFormValues(record);
              setFormType('编辑');
            }}
          >
            编辑
          </a>
          <Popconfirm
            placement="topLeft"
            title={'确认删除吗'}
            onConfirm={() => {
              handleRemove();
            }}
            okText="Yes"
            cancelText="No"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          <TableDropdown
            key="actionGroup"
            onSelect={(e) => history.push({ pathname: `/dash-board/${e}` })}
            menus={[
              { key: 'history', name: '查看历史' },
              { key: 'list', name: '监听查询' },
            ]}
          />
        </>
      ),
    },
  ];

  return (
    <>
      <CommonNamespace />
      <PageContainer
        header={{
          title: '配置列表',
        }}
        style={{ paddingTop: 0 }}
      >
        <ProTable<API.UserInfo>
          headerTitle=""
          actionRef={actionRef}
          rowKey="id"
          search={{
            labelWidth: 120,
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
            const { data, success } = await queryConfigList({
              namespace: namespace,
              ...params,
              // namespace: 'default',
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
          <UpdateForm
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
    </>
  );
};

export default TableList;
