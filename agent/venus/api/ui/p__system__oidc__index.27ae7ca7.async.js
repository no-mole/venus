"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[445],{84313:function(e,r,n){n.r(r),n.d(r,{default:function(){return F}});var t=n(35290),i=n.n(t),u=n(30279),s=n.n(u),a=n(411),c=n.n(a),l=n(46686),o=n.n(l),d=n(10305),p=n(76405),f=n(56453),h=n(30486),m=n(62086),v=n(93236),C=n(12961),k=["fieldProps","unCheckedChildren","checkedChildren","proFieldProps"],w=v.forwardRef((function(e,r){var n=e.fieldProps,t=e.unCheckedChildren,i=e.checkedChildren,u=e.proFieldProps,s=(0,h.Z)(e,k);return(0,m.jsx)(C.Z,(0,f.Z)({valueType:"switch",fieldProps:(0,f.Z)({unCheckedChildren:t,checkedChildren:i},n),ref:r,valuePropName:"checked",proFieldProps:u,filedConfig:{valuePropName:"checked",ignoreWidth:!0}},s))})),x=n(1896),b=n(51955),y=n(62995),_=n(61388);function P(e){return Z.apply(this,arguments)}function Z(){return(Z=c()(i()().mark((function e(r){return i()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,_.request)("/api/v1/sys_config",{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function j(e){return g.apply(this,arguments)}function g(){return(g=c()(i()().mark((function e(r){return i()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,_.request)("/api/v1/sys_config",{method:"POST",data:r}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var q={labelCol:{span:2},wrapperCol:{span:16}},F=function(){var e=(0,v.useState)({}),r=o()(e,2),n=(r[0],r[1],(0,v.useRef)()),t=function(){var e=c()(i()().mark((function e(){var r,t;return i()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,P({});case 2:0==(null==(t=e.sent)?void 0:t.code)&&null!=t&&null!==(r=t.data)&&void 0!==r&&r.oidc?null==n||n.current.setFieldsValue(s()(s()({},t.data.oidc),{},{oidc_status:1==t.data.oidc.oidc_status})):null==n||n.current.setFieldsValue({});case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return(0,v.useEffect)((function(){t()}),[]),(0,m.jsx)(d._z,{header:{title:"OIDC"},children:(0,m.jsxs)(p.A,s()(s()({formRef:n},q),{},{layout:"horizontal",onFinish:function(){var e=c()(i()().mark((function e(r){var n,u;return i()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return(n=r).oidc_status=r.oidc_status?1:-1,e.next=4,j({oidc:n});case 4:0==(null==(u=e.sent)?void 0:u.code)?(b.ZP.success("操作成功"),t()):b.ZP.error("操作失败，请稍后再试");case 6:case"end":return e.stop()}}),e)})));return function(r){return e.apply(this,arguments)}}(),submitter:{render:function(e,r){return[(0,m.jsx)(y.Z,{type:"primary",onClick:function(){var r,n;return null===(r=e.form)||void 0===r||null===(n=r.submit)||void 0===n?void 0:n.call(r)},children:"更新"},"submit")]}},children:[(0,m.jsx)(w,{name:"oidc_status",label:"开启OIDC",required:!0}),(0,m.jsx)(x.Z,{width:"md",name:"oauth_server",label:"OAuthServer",rules:[{required:!0,message:"请填写OAuthServer"}]}),(0,m.jsx)(x.Z,{width:"md",name:"client_id",label:"ClientID",rules:[{required:!0,message:"请填写ClientID"}]}),(0,m.jsx)(x.Z,{width:"md",name:"client_secret",label:"ClientSecret",rules:[{required:!0,message:"请填写ClientSecret"}]}),(0,m.jsx)(x.Z,{width:"md",name:"redirect_uri",label:"RedireUri",rules:[{required:!0,message:"请填写RedireUri"}]})]}))})}}}]);