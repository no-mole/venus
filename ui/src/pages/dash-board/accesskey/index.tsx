import {
  ModalForm,
  PageContainer,
  ProDescriptionsItemProps,
  ProFormText,
  ProTable,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useEffect, useRef, useState } from 'react';
import { history, useModel } from 'umi';
import AccessKeyForm from '../components/AccessKeyForm';
import CommonNamespace from '../components/CommonNamespace';
import styles from './../config/index.less';
import {
  deleteAccessKey,
  deleteSystermListAccessKey,
  getAccessKeyList,
  getSystermAccessKeyList,
  postNewAccesskey,
  putAccessKeyStatus,
} from './service';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const [makeakVisibel, setMakeakVisibel] = useState<boolean>(false);
  const [roleTypeBollean, setRoleTypeBollean] = useState<boolean>(true);

  const namespace = JSON.parse(
    // @ts-ignore
    localStorage.getItem('use-local-storage-state-namespace'),
  );
  // @ts-ignore
  const { select } = useModel('useUser');
  const actionRef = useRef<any>();

  const handleRemove = async (ak: string) => {
    let res;
    if (roleTypeBollean) {
      res = await deleteAccessKey({ namespace: namespace.value, ak: ak });
    } else {
      res = await deleteSystermListAccessKey({ ak: ak });
    }

    if (res?.code == 0) {
      message.success('删除成功');
      actionRef?.current.reload();
    } else {
      message.error(res?.mes || '操作失败，请稍后再试');
    }
  };

  const downloadFun = function (content: any, filename: any) {
    const dom = document.createElement('a');
    dom.download = filename || '文件';
    dom.style.display = 'none';
    const blob = new Blob([content]);
    dom.href = URL.createObjectURL(blob);
    document.body.appendChild(dom);
    dom.click();
    document.body.removeChild(dom);
  };

  const columns: ProDescriptionsItemProps[] = [
    {
      title: '关键词',
      dataIndex: 'keyword',
      hideInTable: true,
    },
    {
      title: 'AccessKeyName',
      dataIndex: roleTypeBollean ? 'ak_alias' : 'alias',
      hideInSearch: true,
    },
    {
      title: 'AccessKey',
      dataIndex: 'ak',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '空间状态',
      dataIndex: 'status',
      valueType: 'text',
      hideInSearch: true,
      hideInTable: roleTypeBollean,
      render: (text) => {
        return text == 1 ? '启用' : '禁用';
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
          <a
            onClick={() => {
              history.push(
                {
                  pathname: roleTypeBollean
                    ? `/dash-board/accesskey/detail`
                    : `/system/accesskey/detail`,
                  search: record.ak,
                },
                { ak: record.ak },
              );
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
              handleRemove(record.ak);
            }}
            okText="删除"
            cancelText="取消"
          >
            <a style={{ marginRight: 8 }}>删除</a>
          </Popconfirm>
          {!roleTypeBollean && (
            <a
              onClick={async () => {
                let res = await putAccessKeyStatus({
                  ak: record.ak,
                  status: record.status == 1 ? -1 : 1,
                });
                if (res?.code == 0) {
                  message.success('操作成功');
                  actionRef?.current.reload();
                } else {
                  message.error(res?.msg || '操作失败，请稍后再试');
                }
              }}
            >
              {record.status == 1 ? '禁用' : '启用'}
            </a>
          )}
        </>
      ),
    },
  ];

  useEffect(() => {
    if (history.location.pathname == '/dash-board/accesskey') {
      setRoleTypeBollean(true);
    } else {
      setRoleTypeBollean(false);
    }
  }, [history.location.pathname]);

  useEffect(() => {
    actionRef?.current.reload();
  }, [select]);

  return (
    <>
      <CommonNamespace />
      <PageContainer
        header={{
          title: 'AccessKey管理',
        }}
      >
        <ProTable<API.UserInfo>
          actionRef={actionRef}
          rowKey="ak"
          search={{
            labelWidth: 60,
          }}
          options={false}
          headerTitle={[
            <Button
              key="ketnew"
              type="primary"
              onClick={() => {
                setMakeakVisibel(true);
                setFormType('新建');
              }}
              style={{
                display: roleTypeBollean ? 'block' : 'none',
              }}
            >
              新建
            </Button>,
          ]}
          request={async (params, sorter, filter) => {
            let tableData = [];
            if (roleTypeBollean) {
              let res = await getAccessKeyList({
                // @ts-ignore
                namespace: namespace.value,
              });
              if (res?.code == 0 && res?.data?.items.length > 0) {
                tableData = res.data.items;
              }
            } else {
              let res = await getSystermAccessKeyList({});
              if (res?.code == 0 && res?.data?.items.length > 0) {
                tableData = res.data.items;
              }
            }

            return {
              data: tableData || [],
              // success,
            };
          }}
          columns={columns}
          rowClassName={(record, index) => {
            let className = styles.lightRow;
            if (index % 2 === 1) className = styles.darkRow;
            return className;
          }}
        />
        <ModalForm<{
          alias: string;
        }>
          layout="horizontal"
          open={makeakVisibel}
          title="新建AccessKey"
          autoFocusFirstInput
          modalProps={{
            destroyOnClose: true,
            onCancel: () => setMakeakVisibel(false),
          }}
          submitTimeout={2000}
          onFinish={async (values: { alias: string }) => {
            let res = await postNewAccesskey({
              namespace: namespace.value,
              alias: values.alias,
            });
            if (res?.code == 0) {
              setFormValues({
                ak: res.data.ak,
                password: res.data.password,
              });
              setMakeakVisibel(false);
              handleUpdateModalVisible(true);
              actionRef?.current.reload();
            } else {
              message.error(res?.msg || '操作失败，请稍后再试');
            }
          }}
        >
          <ProFormText
            width="md"
            name="alias"
            label="AccessKey"
            placeholder="请填写AccessKey"
            fieldProps={{
              max: 16,
            }}
            rules={[{ required: true, message: '请填写AccessKey' }]}
          />
        </ModalForm>

        {/* 更新 */}
        <AccessKeyForm
          formType={formType}
          onSubmit={async (value) => {
            handleUpdateModalVisible(false);
          }}
          onCancel={() => {
            handleUpdateModalVisible(false);
          }}
          updateModalVisible={updateModalVisible}
          values={formValues}
          onDownLoad={() => {
            const result = formValues;
            const content = JSON.stringify(result);
            const data = new Blob([content], {
              type: 'application/json,charset=utf-8;',
            });
            downloadFun(data, 'config.json');
            handleUpdateModalVisible(false);
          }}
        />
      </PageContainer>
    </>
  );
};

export default TableList;
