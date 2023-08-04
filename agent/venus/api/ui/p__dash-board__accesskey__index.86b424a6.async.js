"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[452],{84481:function(e,t,n){n.r(t),n.d(t,{default:function(){return E}});var r=n(35290),a=n.n(r),c=n(411),s=n.n(c),u=n(46686),o=n.n(u),i=n(10305),l=n(13246),p=n(10571),d=n(1896),f=n(51955),h=n(59019),m=n(62995),y=n(93236),k=n(61388),v=n(30279),x=n.n(v),w=n(56453),b={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M505.7 661a8 8 0 0012.6 0l112-141.7c4.1-5.2.4-12.9-6.3-12.9h-74.1V168c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v338.3H400c-6.7 0-10.4 7.7-6.3 12.9l112 141.8zM878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8z"}}]},name:"download",theme:"outlined"},j=n(967),T=function(e,t){return y.createElement(j.Z,(0,w.Z)((0,w.Z)({},e),{},{ref:t,icon:b}))};T.displayName="DownloadOutlined";var C=y.forwardRef(T),S=n(62086),g={labelCol:{span:4},wrapperCol:{span:18}},Z=function(e){return(0,S.jsxs)(p.Y,x()(x()({title:"新建"},g),{},{visible:e.updateModalVisible,layout:"horizontal",autoFocusFirstInput:!0,modalProps:{destroyOnClose:!0,onCancel:function(){return e.onCancel()}},submitter:{render:function(){return[(0,S.jsxs)(m.Z.Group,{style:{display:"block"},children:[(0,S.jsx)(m.Z,{htmlType:"button",onClick:function(){e.onSubmit()},children:"确定"},"sure"),(0,S.jsx)(m.Z,{htmlType:"button",icon:(0,S.jsx)(C,{}),type:"primary",onClick:function(){return e.onDownLoad()},children:"下载"},"down")]},"refs")]}},initialValues:e.values,submitTimeout:2e3,width:700,children:[(0,S.jsx)(d.Z,{width:"xl",name:"ak",label:"AccessKey",disabled:!0}),(0,S.jsx)(d.Z,{width:"xl",name:"password",label:"AccessSecret",extra:"请谨慎保存AccessKey和AccessSecret，关闭后不可再查看AccessSecret",disabled:!0})]}))},I=n(78076),_=n(14362),E=function(){var e=(0,y.useState)(!1),t=o()(e,2),n=t[0],r=t[1],c=(0,y.useState)({}),u=o()(c,2),v=u[0],x=u[1],w=(0,y.useState)(""),b=o()(w,2),j=b[0],T=b[1],C=(0,y.useState)(!1),g=o()(C,2),E=g[0],K=g[1],R=(0,y.useState)(!0),A=o()(R,2),q=A[0],F=A[1],P=(0,y.useState)([]),O=o()(P,2),N=O[0],L=O[1],V=JSON.parse(localStorage.getItem("use-local-storage-state-namespace")),G=(0,k.useModel)("useUser").select,z=(0,y.useRef)(),D=function(){var e=s()(a()().mark((function e(t){var n,r,c;return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!q){e.next=6;break}return e.next=3,(0,_.JF)({namespace:V.value,ak:t});case 3:r=e.sent,e.next=9;break;case 6:return e.next=8,(0,_.N4)({ak:t});case 8:r=e.sent;case 9:0==(null===(n=r)||void 0===n?void 0:n.code)?(f.ZP.success("删除成功"),null==z||z.current.reload()):f.ZP.error((null===(c=r)||void 0===c?void 0:c.mes)||"操作失败，请稍后再试");case 10:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),M=[{title:"关键词",dataIndex:"keyword",hideInTable:!0},{title:"AccessKeyName",dataIndex:q?"ak_alias":"alias",hideInSearch:!0},{title:"AccessKey",dataIndex:"ak",valueType:"text",hideInSearch:!0},{title:"空间状态",dataIndex:"status",valueType:"text",hideInSearch:!0,hideInTable:q,render:function(e){return 1==e?"启用":"禁用"}},{title:"更新人",hideInSearch:!0,dataIndex:"updater",hideInForm:!0},{title:"更新时间",hideInSearch:!0,dataIndex:"update_time",hideInForm:!0,valueType:"dateTime"},{title:"操作",dataIndex:"option",valueType:"option",render:function(e,t,n,r){return(0,S.jsxs)(S.Fragment,{children:[(0,S.jsx)("a",{onClick:function(){k.history.push({pathname:q?"/dash-board/accesskey/detail":"/system/accesskey/detail",search:t.ak},{ak:t.ak})},rel:"noopener noreferrer",style:{marginRight:8},children:"查看"}),(0,S.jsx)(h.Z,{placement:"topLeft",title:"确认删除吗",onConfirm:function(){D(t.ak)},okText:"删除",cancelText:"取消",children:(0,S.jsx)("a",{style:{marginRight:8},children:"删除"})}),!q&&(0,S.jsx)("a",{onClick:s()(a()().mark((function e(){var n;return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,_.Ct)({ak:t.ak,status:1==t.status?-1:1});case 2:0==(null==(n=e.sent)?void 0:n.code)?(f.ZP.success("操作成功"),null==z||z.current.reload()):f.ZP.error((null==n?void 0:n.msg)||"操作失败，请稍后再试");case 4:case"end":return e.stop()}}),e)}))),children:1==t.status?"禁用":"启用"})]})}}];return(0,y.useEffect)((function(){"/dash-board/accesskey"==k.history.location.pathname||"/ui/dash-board/accesskey"==k.history.location.pathname?F(!0):F(!1)}),[k.history.location.pathname]),(0,y.useEffect)((function(){null==z||z.current.reload()}),[G]),(0,S.jsx)(S.Fragment,{children:(0,S.jsxs)(i._z,{header:{title:"AccessKey管理"},children:[(0,S.jsx)(l.Z,{actionRef:z,rowKey:"ak",search:{labelWidth:60},options:!1,toolBarRender:function(){return[(0,S.jsx)(m.Z,{type:"primary",onClick:function(){K(!0),T("新建")},style:{display:q?"block":"none"},children:"新建"},"ketnew")]},request:function(){var e=s()(a()().mark((function e(t,n,r){var c,s,u,o,i,l;return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(c=[],s=[],t.keyword){e.next=17;break}if(!q){e.next=10;break}return e.next=6,(0,_.Kn)({namespace:V.value});case 6:0==(null==(o=e.sent)?void 0:o.code)&&(null==o||null===(u=o.data)||void 0===u?void 0:u.items.length)>0&&(c=o.data.items),e.next=14;break;case 10:return e.next=12,(0,_.ab)({});case 12:0==(null==(l=e.sent)?void 0:l.code)&&(null==l||null===(i=l.data)||void 0===i?void 0:i.items.length)>0&&(c=l.data.items);case 14:L(c),e.next=18;break;case 17:N.map((function(e){-1!=e.ak_alias.indexOf(t.keyword)&&s.push(e)}));case 18:return e.abrupt("return",{data:t.keyword?s||[]:c});case 19:case"end":return e.stop()}}),e)})));return function(t,n,r){return e.apply(this,arguments)}}(),columns:M,rowClassName:function(e,t){var n=I.Z.lightRow;return t%2==1&&(n=I.Z.darkRow),n}}),(0,S.jsx)(p.Y,{layout:"horizontal",open:E,title:"新建AccessKey",autoFocusFirstInput:!0,modalProps:{destroyOnClose:!0,onCancel:function(){return K(!1)}},submitTimeout:2e3,onFinish:function(){var e=s()(a()().mark((function e(t){var n;return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,_.ik)({namespace:V.value,alias:t.alias});case 2:0==(null==(n=e.sent)?void 0:n.code)?(x({ak:n.data.ak,password:n.data.password}),K(!1),r(!0),null==z||z.current.reload()):f.ZP.error((null==n?void 0:n.msg)||"操作失败，请稍后再试");case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),children:(0,S.jsx)(d.Z,{width:"md",name:"alias",label:"AccessKey别名",placeholder:"请填写AccessKey别名",fieldProps:{max:16},rules:[{required:!0,message:"请填写AccessKey别名"}]})}),(0,S.jsx)(Z,{formType:j,onSubmit:function(){var e=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:r(!1);case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),onCancel:function(){r(!1)},updateModalVisible:n,values:v,onDownLoad:function(){var e=v,t=JSON.stringify(e);!function(e,t){var n=document.createElement("a");n.download=t||"文件",n.style.display="none";var r=new Blob([e]);n.href=URL.createObjectURL(r),document.body.appendChild(n),n.click(),document.body.removeChild(n)}(new Blob([t],{type:"application/json,charset=utf-8;"}),"config.json"),r(!1)}})]})})}},14362:function(e,t,n){n.d(t,{Ct:function(){return C},JF:function(){return h},Kn:function(){return o},N4:function(){return j},R8:function(){return d},VK:function(){return v},Xd:function(){return y},ab:function(){return w},ik:function(){return l}});var r=n(35290),a=n.n(r),c=n(411),s=n.n(c),u=n(61388);function o(e){return i.apply(this,arguments)}function i(){return(i=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/namespace/".concat(t.namespace,"/access_key"),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function l(e){return p.apply(this,arguments)}function p(){return(p=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/access_key/".concat(t.namespace,"/").concat(t.alias),{method:"POST",data:t}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function d(e){return f.apply(this,arguments)}function f(){return(f=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/access_key/".concat(t.ak,"/namespace"),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function h(e){return m.apply(this,arguments)}function m(){return(m=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/namespace/".concat(t.namespace,"/access_key/").concat(t.ak),{method:"DELETE"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function y(e){return k.apply(this,arguments)}function k(){return(k=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/namespace",{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function v(e){return x.apply(this,arguments)}function x(){return(x=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/namespace/".concat(t.namespace,"/access_key/").concat(t.ak),{method:"POST",data:t}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function w(e){return b.apply(this,arguments)}function b(){return(b=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/access_key",{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function j(e){return T.apply(this,arguments)}function T(){return(T=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/access_key/".concat(t.ak),{method:"DELETE"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function C(e){return S.apply(this,arguments)}function S(){return(S=s()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,u.request)("/api/v1/access_key/".concat(t.ak),{method:"PUT",data:t}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}},78076:function(e,t){t.Z={darkRow:"darkRow___94Git","panel-table":"panel-table___tjWYN"}}}]);