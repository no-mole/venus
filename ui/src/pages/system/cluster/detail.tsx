import {
    PageContainer,
    ProCard
} from '@ant-design/pro-components';
import { message } from 'antd';
import React, { useState } from 'react';

const TableList: React.FC<unknown> = () => {
    const a = {
        "namespace": "string",
        "role": "string",
        "uid": "string",
        "update_time": "string",
        "updater": "string",
        "user_name": "string"
    }
    return (
        <PageContainer
            header={{
                title: '节点详情',
            }}
        >
            <ProCard style={{ marginBlockEnd: 16 }}>
                <pre>{JSON.stringify(a, null, 2)}</pre>
            </ProCard>

        </PageContainer>
    );
};

export default TableList;
