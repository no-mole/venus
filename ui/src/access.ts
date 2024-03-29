export default (initialState: any) => {
  if (!initialState?.role) return false;
  const { role } = initialState;
  // 在这里按照初始化数据定义项目中的权限，统一管理
  // 参考文档 https://next.umijs.org/docs/max/access
  // const canSeeAdmin = initialState && role === 'UserRoleAdministrator';
  return {
    canReadFoo: true,
    canUpdateFoo: role === 'UserRoleAdministrator',
  };
};
