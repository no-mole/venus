import React, { useEffect, useRef, useState } from 'react';
// import jsonlint from 'jsonlint'; //全局化

import { Controlled as CodeMirror } from 'react-codemirror2';

import 'codemirror/lib/codemirror.css';
import 'codemirror/theme/material.css';

import 'codemirror/mode/xml/xml';
// 设置代码语言模式（比如JS，SQL，python，java等）
import 'codemirror/mode/javascript/javascript';
import 'codemirror/mode/yaml/yaml.js';
import 'codemirror/mode/toml/toml.js';
import 'codemirror/mode/properties/properties.js';
// 引入背景主题
import 'codemirror/theme/yonce.css';
//ctrl+空格代码提示补全
import 'codemirror/addon/hint/show-hint.css'; // start-ctrl+空格代码提示补全
import 'codemirror/addon/hint/show-hint.js';
import 'codemirror/addon/hint/sql-hint';
import 'codemirror/addon/hint/anyword-hint.js'; // end
//代码高亮
import 'codemirror/addon/selection/active-line';
//折叠代码
import 'codemirror/addon/fold/foldgutter.css';
import 'codemirror/addon/fold/foldcode.js';
import 'codemirror/addon/fold/foldgutter.js';
import 'codemirror/addon/fold/brace-fold.js';
import 'codemirror/addon/fold/comment-fold.js';

const EditOrViewCode = ({ codeValue, type, formType, changeCode }: any) => {
  const [code, setCode] = useState('');
  let cm: any;
  useEffect(() => {
    // 解决行号错误的问题
    setTimeout(() => {
      // @ts-ignore
      cm = document?.querySelector('div.CodeMirror')?.CodeMirror;
      cm.refresh();
      if (typeof codeValue === 'string') {
        setCode(JSON.parse(codeValue));
      } else {
        setCode(JSON.stringify(codeValue, null, 2));
      }
    }, 500);
  }, []);

  useEffect(() => {
    cm = document?.querySelector('div.CodeMirror')?.CodeMirror;
    setTimeout(() => {
      // @ts-ignore
      cm.refresh();
    }, 100);
  }, [type]);

  return (
    <>
      <CodeMirror
        value={code}
        onBeforeChange={(editor, data, value) => {
          if (!value) return;
          changeCode(editor, data, value);
          if (formType === '新建' || formType === '编辑') {
            setCode(value);
          }
        }}
        // onKeyPress={() => {
        //   cm.showHint();
        // }}
        // onChange={(editor, newValue, changeObj) => {
        //   changeCode(editor, newValue, changeObj);
        // }}
        options={{
          lineWrapping: true,
          autofocus: true,
          // lint: true,
          indentWithTabs: true,
          mode: {
            name:
              type === 'yaml' || type === 'toml' || type === 'properties'
                ? 'text/x-' + type
                : 'text/jsx',
            json: true,
          },
          theme: 'material',
          lineNumbers: true,
          readOnly: formType === '详情',
          smartIndent: true, //自动缩进
          indentUnit: 4,
          // autoRefresh: true,
          //start-设置支持代码折叠
          foldGutter: true,
          styleActiveLine: true, // 当前行背景高亮
          matchBrackets: true, // 括号匹配
          autoRefresh: true,
          gutters: [
            'CodeMirror-linenumbers',
            'CodeMirror-foldgutter',
            'CodeMirror-lint-markers',
          ], //end
          extraKeys: {
            Ctrl: 'autocomplete',
            'Ctrl-S': function (editor) {
              // codeSave(editor);
            },
            'Ctrl-Z': function (editor) {
              editor.undo();
            }, //undo
            F8: function (editor) {
              editor.redo();
            }, //Redo
          },
          autoCloseBrackets: true, //键入时将自动关闭()[]{}''""
        }}
      />
    </>
  );
};

export default EditOrViewCode;
