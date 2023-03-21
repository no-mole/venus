import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import NameSpaceForm from './NameSpaceForm';

import styles from './../config/index.less';
import { getList, postAddUser, postDeleteUser, getUserList } from './service';
import { useLocation } from 'umi';
import { useModel } from '@umijs/max';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑
  const actionRef = useRef<ActionType>();
  const { search } = useLocation();
  const [namespace_uid, setNamespace_uid] = useState('');
  const [namespace_alias, setNamespace_alias] = useState<any>('');
  const [userList, setUserList] = useState([]);
  //@ts-ignore
  const { select, add } = useModel('useUser', (model: any) => ({
    select: model.select,
    add: model.increment,
  }));
  let searchParams = new URLSearchParams(search);
  const namespaceUid = searchParams.get('namespaceUid');
  const namespaceAlias = searchParams.get('namespaceAlias');

  const getUserListData = async () => {
    const res = await getUserList();
    if (res?.code === 0) {
      setUserList(res?.data?.items || []);
    } else {
      message.error('用户列表数据获取失败');
    }
  };

  //新增/编辑
  const handleAddAndUpdate = async (fields: any) => {
    const hide = message.loading('正在添加');
    const filterData: any = userList?.filter((item: any) => {
      return item?.uid === fields?.uid;
    });
    try {
      await postAddUser({ ...fields, user_name: filterData[0]?.name });
      hide();
      message.success('添加成功');
      return true;
    } catch (error) {
      hide();
      message.error('添加失败请重试！');
      return false;
    }
  };

  //删除
  const handleRemove = async (record: any) => {
    const hide = message.loading('正在删除');
    try {
      await postDeleteUser({
        uid: record?.uid,
        namespace: record?.namespace_uid,
      });
      hide();
      add(''); // 删除后需要更新namespace下拉列表
      actionRef?.current?.reload();
      message.success('删除成功');
      return true;
    } catch (error) {
      hide();
      message.error('删除失败，请重试');
      return false;
    }
  };

  useEffect(() => {
    console.log('select', select);
    if (select) {
      setNamespace_uid(select?.value);
      setNamespace_alias(select?.label);
      if (actionRef.current) {
        actionRef.current.reload();
      }
    }
  }, [select]);

  useEffect(() => {
    //namespaceUid存在，代表页面从系统管理-命名空间-查看 跳转过来的
    if (namespaceUid) {
      setNamespace_uid(namespaceUid);
      setNamespace_alias(namespaceAlias);
    }
  }, []);

  //打开弹窗
  const handleClickAddAndUpdateBtn = (type: string, record: any) => {
    handleUpdateModalVisible(true);
    setFormValues(record);
    setFormType(type);
    getUserListData();
  };

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
        wr: { text: '空间管理员' },
        r: { text: '只读成员' },
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
            title={`删除用户${record?.user_name}对命名空间${record?.namespace_alias}的访问授权？`}
            onConfirm={() => handleRemove(record)}
            okText="确定"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          <a
            onClick={() => handleClickAddAndUpdateBtn('修改', record)}
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
    <>
      {/* {!namespaceUid && <CommonNamespace />} */}
      <PageContainer
        header={{
          title: '命名空间用户权限管理',
        }}
      >
        <ProTable<API.UserInfo>
          headerTitle=""
          actionRef={actionRef}
          rowKey="uid"
          search={false}
          options={false}
          toolBarRender={() => [
            <Button
              key="1"
              type="primary"
              onClick={() => handleClickAddAndUpdateBtn('新建', '')}
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
              namespace: namespace_uid,
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
            userList={userList}
            namespace_alias={namespace_alias}
            namespace_uid={namespace_uid}
            onSubmit={async (value) => {
              let success = await handleAddAndUpdate(value);
              if (success) {
                handleUpdateModalVisible(false);
                setFormValues({});
                add('');
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
