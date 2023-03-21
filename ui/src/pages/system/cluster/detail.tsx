import {
    PageContainer,
    ProCard
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useEffect, useState } from 'react';
import { useLocation } from 'umi';
import { getDetail } from './service'

const TableList: React.FC<unknown> = () => {
    const { search } = useLocation();
    let searchParams = new URLSearchParams(search);
    const id = searchParams.get('id');
    const nodeInfo = searchParams.get('nodeInfo');
    const [data, setData] = useState({})

    const getData = async () => {
        const res = await getDetail({
            id: id
        });
        if (res?.code === 0) {
            setData(res?.data || {})
        } else {
            message.error('节点详情接口数据获取失败')
        }
    }

    useEffect(() => {
        getData();
    }, [])

    return (
        <PageContainer
            header={{
                title: '节点详情-' + nodeInfo,
            }}
        >
            <ProCard style={{ marginBlockEnd: 16 }}>
                <pre>{JSON.stringify(data, null, 2)}</pre>
            </ProCard>

        </PageContainer>
    );
};

export default TableList;
