/* eslint-disable react-hooks/rules-of-hooks */
// 全局共享数据示例

import { getCommonNamespace } from '@/pages/dash-board/config/service';
import { useModel } from '@umijs/max';
import { useRequest } from '@umijs/max';
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

  // 获取namespace接口
  const { loading: loading } = useRequest(async () => {
    if (uid !== '') {
      const res: any = await getCommonNamespace({ uid: uid });
      if (res?.code === 0 && res?.data?.length > 0) {
        setList(res?.data);
        setSelect(res?.data[0]);
        return res;
      } else {
        return {};
      }
    } else {
      return {};
    }
  });

  const increment = async (id: string) => {
    let userId = !id || id === '' ? initialState?.uid : id;
    const res: any = await getCommonNamespace({ uid: userId });
    if (res?.code === 0 && res?.data?.length > 0) {
      setList(res?.data);
      setSelect(res?.data[0]);
      return res;
    } else {
      return {};
    }
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
