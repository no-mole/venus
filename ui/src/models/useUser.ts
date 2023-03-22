/* eslint-disable react-hooks/rules-of-hooks */
// 全局共享数据示例

import { getCommonNamespace } from '@/pages/dash-board/config/service';
import { useModel } from '@umijs/max';
import { useRequest } from '@umijs/max';
import { message, notification } from 'antd';
import { useEffect, useState } from 'react';

const useUser = () => {
  const { initialState } = useModel('@@initialState');
  const [select, setSelect] = useState({});
  const [list, setList] = useState([]);
  let uid: string = '';
  const userinfo = localStorage.getItem('userinfo');

  // 取出uid
  if (userinfo) {
    uid = JSON.parse(userinfo).uid;
  }

  const requestUser = async (uid: any) => {
    const res: any = await getCommonNamespace({ uid: uid });
    if (res?.code === 0) {
      if (res?.data?.length > 0) {
        setList(res?.data);
        setSelect(res?.data[0]);
        return res;
      } else if (!res.data || res?.data?.length === 0) {
        // 如果空间为空
        notification.warning({
          message: `提醒 `,
          description:
            '您还没有被授权任何空间访问权限，请联系对应空间管理员授权或系统管理员创建新的命名空间！',
          duration: 10,
        });
      } else {
        return {};
      }
    } else {
      return {};
    }
  };

  // 获取namespace接口
  const { loading: loading } = useRequest(async () => {
    if (uid !== '') {
      requestUser(uid);
    }
  });

  const increment = async (id: string) => {
    console.log('increment', increment);
    let userId = !id || id === '' ? initialState?.uid : id;
    requestUser(userId);
  };

  return {
    list,
    loading,
    setList,
    select,
    setSelect,
    increment,
  };
};

export default useUser;
