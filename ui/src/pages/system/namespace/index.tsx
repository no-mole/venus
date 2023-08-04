import {
  ActionType,
  PageContainer,
  ProCard,
  ProDescriptionsItemProps,
  ProTable,
} from '@ant-design/pro-components';
import { Button, Input, Space, message } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { history } from 'umi';
import CreateForm from './CreateForm';
import { getList, postAddNamespace } from './service';

const TableList: React.FC<unknown> = () => {
  const [createModalVisible, handleModalVisible] = useState<boolean>(false);
  const actionRef = useRef<ActionType>();
  const [data, setData] = useState([]);
  const [copyData, setCopyData] = useState([]); //备份数据
  const [searchVal, setSearchVal] = useState('');

  const getListData = async () => {
    const res = await getList();
    if (res?.code === 0) {
      setData(res?.data?.items);
      setCopyData(res?.data?.items);
    }
  };

  useEffect(() => {
    getListData();
  }, []);

  const handleAdd = async (fields: API.UserInfo) => {
    const hide = message.loading('正在添加');
    try {
      await postAddNamespace({ ...fields });
      hide();
      message.success('添加成功');
      getListData();
      return true;
    } catch (error) {
      hide();
      message.error('添加失败请重试！');
      return false;
    }
  };

  // const handleRemove = async (record: any) => {
  //   const hide = message.loading('正在删除');
  //   if (!record) return true;
  //   try {
  //     await postDeleteUser({
  //       namespace: record?.namespace_uid,
  //     });
  //     hide();
  //     message.success('删除成功');
  //     getListData();
  //     return true;
  //   } catch (error) {
  //     hide();
  //     message.error('删除失败，请重试');
  //     return false;
  //   }
  // };

  const columns: ProDescriptionsItemProps<API.UserInfo>[] = [
    {
      title: '空间名称',
      dataIndex: 'namespace_alias',
    },
    {
      title: '唯一标识',
      dataIndex: 'namespace_uid',
    },
    {
      title: '创建时间',
      valueType: 'dateTime',
      dataIndex: 'create_time',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record: any) => (
        <>
          <a
            onClick={() => {
              history.push({
                pathname: `/system/namespace/detail`,
                search: `?namespaceUid=${record?.namespace_uid}&namespaceAlias=${record?.namespace_alias}`,
              });
            }}
          >
            查看
          </a>
          {/* <Divider type="vertical" />
          <Popconfirm
            title={`删除空间 ${record?.namespace_alias}（${record?.namespace_uid}）?`}
            description={
              () => <div style={{ color: 'red' }}>该操作会清空空间下的所有配置项<br />和注册的服务，请谨慎操作</div>
            }
            okText="确定"
            cancelText="取消"
            onConfirm={() => handleRemove(record)}
          >
            <a>删除</a>
          </Popconfirm> */}
        </>
      ),
    },
  ];

  const handleSearch = () => {
    const filterData = copyData?.filter((item: any) => {
      return item?.namespace_uid?.indexOf(searchVal) > -1;
    });
    setData(filterData);
  };

  const handleReset = () => {
    setSearchVal('');
    getListData();
  };

  return (
    <PageContainer
      header={{
        title: '命名空间管理',
      }}
    >
      <ProCard style={{ marginBlockEnd: 16 }}>
        <Space>
          <ProCard>
            <div>
              关键词：
              <Input
                style={{ width: '240px' }}
                value={searchVal}
                onChange={(e) => setSearchVal(e.target.value)}
                placeholder="请输入命名空间唯一标识"
              />
            </div>
          </ProCard>
          <Button type="primary" onClick={handleSearch}>
            查询
          </Button>
          <Button onClick={handleReset}>重置</Button>
        </Space>
      </ProCard>
      <ProTable<API.UserInfo>
        actionRef={actionRef}
        rowKey="namespace_uid"
        toolBarRender={() => [
          <Button
            key="1"
            type="primary"
            onClick={() => handleModalVisible(true)}
          >
            新建
          </Button>,
        ]}
        options={false}
        dataSource={data}
        columns={columns}
        search={false}
      />
      <CreateForm
        onCancel={() => handleModalVisible(false)}
        modalVisible={createModalVisible}
        onSubmit={async (value) => {
          const success = await handleAdd(value);
          if (success) {
            handleModalVisible(false);
          }
        }}
      ></CreateForm>
    </PageContainer>
  );
};

export default TableList;
