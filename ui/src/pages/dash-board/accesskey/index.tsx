import services from '@/services/demo';
import {
  ActionType,
  ModalForm,
  PageContainer,
  ProDescriptionsItemProps,
  ProFormText,
  ProTable,
} from '@ant-design/pro-components';
import { Button, message, Popconfirm } from 'antd';
import React, { useRef, useState } from 'react';
import { history } from 'umi';
import AccessKeyForm from '../components/AccessKeyForm';
import styles from './../config/index.less';
import { deleteAccessKey, getAccessKeyList, postNewAccesskey } from './service';

const TableList: React.FC<unknown> = () => {
  const [updateModalVisible, handleUpdateModalVisible] =
    useState<boolean>(false);
  const [formValues, setFormValues] = useState({});
  const [formType, setFormType] = useState(''); // 弹窗类型，新建、编辑、查看
  const [makeakVisibel, setMakeakVisibel] = useState<boolean>(false);
  const actionRef = useRef<any>();

  const handleRemove = async (ak: string) => {
    let res = await deleteAccessKey({ namespace: 'comos', ak: ak });
    if (res?.code == 0) {
      message.success('删除成功');
      actionRef?.current.reload();
    } else {
      message.error(res?.mes || '操作失败，请稍后再试');
    }
  };

  const downloadFun = function (content: any, filename: any) {
    // 创建隐藏的可下载链接 A 标签
    const dom = document.createElement('a');
    dom.download = filename || '未命名文件';
    // 隐藏
    dom.style.display = 'none';
    // 将字符内容转换为成 blob 二进制
    const blob = new Blob([content]);
    // 创建对象 URL
    dom.href = URL.createObjectURL(blob);
    // 添加 A 标签到 DOM
    document.body.appendChild(dom);
    // 模拟触发点击
    dom.click();
    // 或
    // dom.dispatchEvent(new MouseEvent('click'))
    // 移除 A 标签
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
      dataIndex: 'ak',
      hideInSearch: true,
    },
    {
      title: 'AccessKey',
      dataIndex: 'key',
      valueType: 'text',
      hideInSearch: true,
    },
    {
      title: '创建时间',
      hideInSearch: true,
      dataIndex: 'update_time',
      hideInForm: true,
    },
    {
      title: '上次登陆时间',
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
            onClick={() => {
              history.push(
                {
                  pathname: `/dash-board/accesskey/detail`,
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
        </>
      ),
    },
  ];

  return (
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
          >
            新建
          </Button>,
        ]}
        request={async (params, sorter, filter) => {
          const { data, success } = await getAccessKeyList({
            namespace: 'comos',
          });
          console.log(data);
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
            namespace: 'comos',
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
          name="ak"
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
  );
};

export default TableList;
