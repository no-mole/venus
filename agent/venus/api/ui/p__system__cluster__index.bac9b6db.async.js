"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[574],{85460:function(t,e,n){n.r(e);var r=n(35290),a=n.n(r),u=n(411),s=n.n(u),c=n(10305),i=n(13246),o=(n(93236),n(61388)),d=n(69615),l=n(62086);e.default=function(){var t=[{title:"节点ID",dataIndex:"id"},{title:"节点入口",dataIndex:"address"},{title:"角色",dataIndex:"state"},{title:"是否在线",dataIndex:"online",valueEnum:{true:{text:"在线"},false:{text:"离线"}}},{title:"选举权",dataIndex:"suffrage"},{title:"操作",dataIndex:"option",valueType:"option",render:function(t,e){return(0,l.jsx)("a",{onClick:function(){return o.history.push({pathname:"/system/cluster/detail?id=".concat(null==e?void 0:e.id,"&nodeInfo=").concat(null==e?void 0:e.address)})},children:"查看"})}}];return(0,l.jsx)(c._z,{header:{title:"集群管理"},children:(0,l.jsx)(i.Z,{rowKey:"id",search:!1,toolBarRender:!1,request:s()(a()().mark((function t(){var e,n,r;return a()().wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,(0,d.g)();case 2:return e=t.sent,n=e.data,r=e.success,t.abrupt("return",{data:n||[],success:r});case 6:case"end":return t.stop()}}),t)}))),columns:t})})}},69615:function(t,e,n){n.d(e,{N:function(){return d},g:function(){return i}});var r=n(35290),a=n.n(r),u=n(411),s=n.n(u),c=n(61388);function i(){return o.apply(this,arguments)}function o(){return(o=s()(a()().mark((function t(){return a()().wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,c.request)("/api/v1/cluster",{method:"GET"}));case 1:case"end":return t.stop()}}),t)})))).apply(this,arguments)}function d(t){return l.apply(this,arguments)}function l(){return(l=s()(a()().mark((function t(e){return a()().wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.abrupt("return",(0,c.request)("/api/v1/cluster/".concat(e.id),{method:"GET"}));case 1:case"end":return t.stop()}}),t)})))).apply(this,arguments)}}}]);