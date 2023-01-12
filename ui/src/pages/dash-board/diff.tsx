import React from 'react';
import PropTypes from 'prop-types';
import ReactDiffViewer, { DiffMethod } from 'react-diff-viewer';
import styles from './index.less';

function DiffPanel({ oldValue, newValue }: any) {
  const oldVerifyConf = JSON.stringify(oldValue, null, 2);
  const newVerifyConf = JSON.stringify(newValue, null, 2);

  const newStyles = {
    variables: {
      dark: {
        highlightBackground: '#fefed5',
        highlightGutterBackground: '#ffcd3c',
      },
    },
    line: {
      padding: '10px 2px',
      '&:hover': {
        background: '#a26ea1',
      },
    },
    contentText: {
      width: '390px',
    },
  };
  // const sameCode = React.useMemo(() => (
  //   newVerifyConf === oldVerifyConf
  // ), [newVerifyConf, oldVerifyConf]);
  // 折叠块显示内容
  const codeFoldMessageRenderer: any = React.useCallback(
    (num: any) =>
      newValue === oldValue
        ? '修改前后没有区别，点击展开查看'
        : `展开隐藏的 ${num} 行...`,
    [newValue, oldValue],
  );
  return (
    <div className="am-diff-panel">
      <div className={styles['panel-table']}>
        <p>修改前配置</p>
        <p>修改后配置</p>
      </div>
      <ReactDiffViewer
        oldValue={oldVerifyConf}
        newValue={newVerifyConf}
        styles={newStyles}
        codeFoldMessageRenderer={codeFoldMessageRenderer}
        splitView
        compareMethod={DiffMethod.LINES}
      />
    </div>
  );
}

DiffPanel.propTypes = {
  oldValue: PropTypes.object.isRequired,
  newValue: PropTypes.object.isRequired,
};
// 接口传入的数据
DiffPanel.defaultProps = {
  oldValue: { a: '111asdfsd', b: 'sdsdflslfs' },
  newValue: { a: '111asdfsd', b: 'sdsdflslfs' },
};

export default DiffPanel;
