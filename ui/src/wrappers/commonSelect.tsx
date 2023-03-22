import CommonNamespace from '@/pages/dash-board/components/CommonNamespace';
import { Outlet } from '@umijs/max';

export default () => {
  return (
    <>
      <CommonNamespace />
      <Outlet />
    </>
  );
};
