import {
  PageContainer,
} from '@ant-design/pro-components';
import { message, Collapse, Empty, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import { useModel } from '@umijs/max';
import Styles from './index.less'
import { getList, getListVersion, getListAddr, getDetailInfo } from './service';
const { Panel } = Collapse;


const TableList: React.FC<unknown> = () => {
  const [list, setList] = useState<any>([]);
  const [versions, setVersions] = useState<any>({});
  const [address, setAddress] = useState<any>({});
  const [info, setInfo] = useState<any>({});
  //@ts-ignore
  const { select } = useModel('useUser');
  //@ts-ignore
  const namespace = JSON.parse(localStorage.getItem('use-local-storage-state-namespace') || '{}');
  const [nameArr, setNameArr] = useState<any>([]);
  const [versionArr, setVersionArr] = useState<any>([]);
  const [addrArr, setAddrArr] = useState<any>([]);


  const getListData = async () => {
    const res = await getList({
      namespace: namespace?.value
    });
    if (res?.code === 0) {
      setList(res?.data || []);
    } else {
      message.error('服务列表数据获取失败');
    }
  };

  useEffect(() => {
    if (namespace) {
      getListData()
    }
  }, [select]);

  const getVersions = async (name: string) => {
    if (nameArr?.includes(name)) {
      return
    }
    const res = await getListVersion({
      namespace: namespace?.value,
      name
    });
    if (res?.code === 0) {
      setVersions({ ...versions, [name]: res?.data || [] });
    } else {
      message.error('服务版本数据获取失败');
      return null
    }
  }

  const getAddr = async (name: string, version: string, e: any) => {
    e.stopPropagation();
    const str = `${name}_${version}`
    if (versionArr?.includes(str)) {
      return
    }
    const res = await getListAddr({
      namespace: namespace?.value,
      name,
      version
    });
    if (res?.code === 0) {
      setAddress({ ...address, [str]: res?.data || [] });
    } else {
      message.error('服务入口数据获取失败');
      return null
    }
  }

  const getInfo = async (name: string, version: string, addr: string, e: any) => {
    e.stopPropagation();
    const str = `${name}_${version}_${addr}`
    if (addrArr?.includes(str)) {
      return
    }
    const res = await getDetailInfo({
      namespace: namespace?.value,
      name,
      version,
      addr
    });
    if (res?.code === 0) {
      const client_info = res?.data?.client_info;
      const service_info = res?.data?.service_info;
      const data = [{
        id: 1,
        col1: '空间名称',
        col2: service_info?.namespace,
        col3: 'ACCESS_KEY',
        col4: client_info?.register_access_key
      }, {
        id: 2,
        col1: '服务名称',
        col2: service_info?.service_name,
        col3: 'IP',
        col4: client_info?.register_ip
      }, {
        id: 3,
        col1: '版本',
        col2: service_info?.service_version,
        col3: '注册时间',
        col4: client_info?.register_time
      }, {
        id: 4,
        col1: '服务入口',
        col2: service_info?.service_endpoint,
        col3: '',
        col4: ''
      }]
      setInfo({ ...info, [str]: data });
    } else {
      message.error('服务详情信息数据获取失败');
      return null
    }
  }

  const columns = [
    {
      dataIndex: 'col1',
      key: 'col1',
    },
    {
      dataIndex: 'col2',
      key: 'col2',
    },
    {
      dataIndex: 'col3',
      key: 'col3',
    },
    {
      dataIndex: 'col4',
      key: 'col4',
    },
  ];

  return (
    <div className={Styles.service}>
      <PageContainer
        header={{
          title: '服务管理',
        }}
      >
        {
          list?.length > 0 ?
            <Collapse bordered={false} onChange={(params) => setNameArr(params)}>
              {
                list.map(((name: any) => {
                  return <Panel header={name} key={name} onClick={() => getVersions(name)} >
                    <Collapse bordered={false} onChange={(params) => setVersionArr(params)}>
                      {
                        versions[name]?.map((v: string) => {
                          return <Panel header={v} key={`${name}_${v}`} onClick={(e: any) => getAddr(name, v, e)}>
                            <Collapse bordered={false} onChange={(params) => { setAddrArr(params) }}>
                              {
                                address[`${name}_${v}`]?.map((addr: string) => {
                                  return <Panel header={addr} key={`${name}_${v}_${addr}`} onClick={(e: any) => getInfo(name, v, addr, e)}>
                                    <Table rowKey='id' pagination={false} bordered columns={columns} dataSource={info[`${name}_${v}_${addr}`]} />
                                  </Panel>
                                })
                              }
                            </Collapse>
                          </Panel>
                        })
                      }
                    </Collapse>
                  </Panel>
                }))
              }
            </Collapse> : <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
        }
      </PageContainer>
    </div>
  );
};

export default TableList;
