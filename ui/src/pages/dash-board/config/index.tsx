import {
  ActionType,
  PageContainer,
  ProDescriptionsItemProps,
  ProTable,
  TableDropdown,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import UpdateForm, { FormValueType } from '../components/UpdateForm';
import styles from './index.less';
import {
  queryConfigList,
  addUser,
  deleteConfig,
  modifyConfig,
} from './service';
import { history } from 'umi';
import { useModel } from '@umijs/max';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const actionRef = useRef<ActionType>();
  const { select } = useModel('useUser'); //namespace 选中
  const namespace = localStorage.getItem('use-local-storage-state-namespace'); // 默认namespace
  let namespace_lable: string = '',
    namespace_value: string = ''; // 选中namespace拆解
  if (namespace) {
    namespace_value = JSON.parse(namespace).value;
  }

  useEffect(() => {
    if (actionRef.current) {
      actionRef.current.reload();
    }
  }, [select]);

  /**
   * 更新新增配置节点
   * @param fields
   */
  const handleAddOrUpdate = async (fields: FormValueType) => {
    const hide = message.loading('正在配置');
    try {
      await modifyConfig({ ...fields, namespace: namespace_value });
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
  const handleRemove = async (selectedRows: CONFIG.TableColumns) => {
    const hide = message.loading('正在删除');
    if (!selectedRows) return true;
    try {
      await deleteConfig({
        namespace: namespace_value,
        key: selectedRows?.key,
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

  const columns: ProDescriptionsItemProps<CONFIG.TableColumns>[] = [
    {
      title: '配置项名称',
      dataIndex: 'alias',
    },
    {
      title: '唯一标识',
      dataIndex: 'key',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '版本',
      dataIndex: 'version',
      hideInSearch: true,
    },
    {
      title: '最近更新时间',
      hideInSearch: true,
      valueType: 'date',
      dataIndex: 'update_time',
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (text, record) => (
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
            title={`确认删除配置项${record?.alias}吗`}
            onConfirm={() => {
              handleRemove(record);
            }}
            okText="确定"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          <TableDropdown
            key="actionGroup"
            onSelect={(e) =>
              history.push({
                pathname: `/dash-board/config/${e}?namespace=${record?.namespace}&key=${record?.key}&alias=${record?.alias}`,
              })
            }
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
      <PageContainer
        header={{
          title: '配置列表',
        }}
        style={{ paddingTop: 0 }}
      >
        <ProTable<CONFIG.TableColumns>
          headerTitle=""
          actionRef={actionRef}
          rowKey={(record: any) => record?.key}
          search={{
            labelWidth: 80,
          }}
          toolBarRender={() => [
            <Button
              key="create"
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
              namespace: namespace_value,
              ...params,
              // namespace: 'default',
              // FIXME: remove @ts-ignore
              // @ts-ignore
              sorter,
              filter,
            });
            return {
              data: data || [],
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
              const success = await handleAddOrUpdate(value);
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
