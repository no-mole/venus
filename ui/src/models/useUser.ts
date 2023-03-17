/* eslint-disable react-hooks/rules-of-hooks */
// 全局共享数据示例

import { getCommonNamespace } from '@/pages/dash-board/config/service';
import { useModel } from '@umijs/max';
import { useRequest } from '@umijs/max';
import { useState } from 'react';

const useUser = () => {
  // const { initialState } = useModel('@@initialState');
  const [select, setSelect] = useState({});
  const [list, setList] = useState([]);
  let uid: string = '';
  const userinfo = localStorage.getItem('userinfo');

  // 取出uid
  if (userinfo) {
    uid = JSON.parse(userinfo).uid;
  } else {
    return false;
  }

  // 获取namespace接口
  const { loading: loading } = useRequest(async () => {
    const res: any = await getCommonNamespace({ uid: uid });
    if (res) {
      setList(res?.data);
      setSelect(res?.data[0]);
      return res;
    }
    return {};
  });
  return {
    list,
    loading,
    select,
    setSelect,
  };
};

export default useUser;
