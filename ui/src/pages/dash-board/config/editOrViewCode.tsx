import React from 'react';
// import jsonlint from 'jsonlint'; //全局化

import CodeMirror from 'react-codemirror';
import 'codemirror/lib/codemirror.js';
import 'codemirror/lib/codemirror.css';
// 引入代码模式
import 'codemirror/mode/javascript/javascript';
// 引入背景主题
import 'codemirror/theme/yonce.css';
const EditOrViewCode = ({ codeValue }: any) => {
  return (
    <>
      <CodeMirror
        value={
          typeof codeValue === 'string'
            ? codeValue
            : JSON.stringify(codeValue, null, 2)
        }
        // 设置CodeMirror标签的初始值
        options={{
          mode: {
            // 实现代码高亮
            name: 'text/x-yaml' || 'javascript',
            // name: "javascript", // 没错，需要先引入 javascript
            json: true,
          },
          tabSize: 4,
          autofocus: true, // 自动获取焦点
          lineWrapping: true, // 代码自动换行
          theme: 'yonce', // 代码编译器主题
          // lineNumbers: true, // 显示行号
          readOnly: true,
        }}
      />
    </>
  );
};

export default EditOrViewCode;
