"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[706],{28077:function(e,n,t){t.r(n),t.d(n,{default:function(){return xe}});var r=t(93275),o=t.n(r),a=t(30279),c=t.n(a),i=t(35290),l=t.n(i),s=t(411),d=t.n(s),u=t(46686),p=t.n(u),v=t(10305),f=t(43609),m=t(84875),h=t.n(m),x=t(24255),b=t(94100),g=t(83235),y=t(52113),C=t(1423),$=t(55859),k=t(54465),I=t(7152),E=t(93236),Z=t(50631),w=t.n(Z),P=t(15882),S=t(30486),_=t(73066),N=t(6614),j=E.forwardRef((function(e,n){var t,r=e.prefixCls,o=e.forceRender,a=e.className,c=e.style,i=e.children,l=e.isActive,s=e.role,d=E.useState(l||o),u=(0,N.Z)(d,2),p=u[0],v=u[1];return E.useEffect((function(){(o||l)&&v(!0)}),[o,l]),p?E.createElement("div",{ref:n,className:h()("".concat(r,"-content"),(t={},(0,x.Z)(t,"".concat(r,"-content-active"),l),(0,x.Z)(t,"".concat(r,"-content-inactive"),!l),t),a),style:c,role:s},E.createElement("div",{className:"".concat(r,"-content-box")},i)):null}));j.displayName="PanelContent";var A=j,K=["className","id","style","prefixCls","headerClass","children","isActive","destroyInactivePanel","accordion","forceRender","openMotion","extra","collapsible"],T=function(e){(0,C.Z)(t,e);var n=(0,$.Z)(t);function t(){var e;(0,g.Z)(this,t);for(var r=arguments.length,o=new Array(r),a=0;a<r;a++)o[a]=arguments[a];return(e=n.call.apply(n,[this].concat(o))).onItemClick=function(){var n=e.props,t=n.onItemClick,r=n.panelKey;"function"==typeof t&&t(r)},e.handleKeyPress=function(n){"Enter"!==n.key&&13!==n.keyCode&&13!==n.which||e.onItemClick()},e.renderIcon=function(){var n=e.props,t=n.showArrow,r=n.expandIcon,o=n.prefixCls,a=n.collapsible;if(!t)return null;var c="function"==typeof r?r(e.props):E.createElement("i",{className:"arrow"});return c&&E.createElement("div",{className:"".concat(o,"-expand-icon"),onClick:"header"===a||"icon"===a?e.onItemClick:null},c)},e.renderTitle=function(){var n=e.props,t=n.header,r=n.prefixCls,o=n.collapsible;return E.createElement("span",{className:"".concat(r,"-header-text"),onClick:"header"===o?e.onItemClick:null},t)},e}return(0,y.Z)(t,[{key:"shouldComponentUpdate",value:function(e){return!w()(this.props,e)}},{key:"render",value:function(){var e,n,t=this.props,r=t.className,o=t.id,a=t.style,c=t.prefixCls,i=t.headerClass,l=t.children,s=t.isActive,d=t.destroyInactivePanel,u=t.accordion,p=t.forceRender,v=t.openMotion,f=t.extra,m=t.collapsible,b=(0,S.Z)(t,K),g="disabled"===m,y="header"===m,C="icon"===m,$=h()((e={},(0,x.Z)(e,"".concat(c,"-item"),!0),(0,x.Z)(e,"".concat(c,"-item-active"),s),(0,x.Z)(e,"".concat(c,"-item-disabled"),g),e),r),k={className:h()("".concat(c,"-header"),(n={},(0,x.Z)(n,i,i),(0,x.Z)(n,"".concat(c,"-header-collapsible-only"),y),(0,x.Z)(n,"".concat(c,"-icon-collapsible-only"),C),n)),"aria-expanded":s,"aria-disabled":g,onKeyPress:this.handleKeyPress};y||C||(k.onClick=this.onItemClick,k.role=u?"tab":"button",k.tabIndex=g?-1:0);var I=null!=f&&"boolean"!=typeof f;return delete b.header,delete b.panelKey,delete b.onItemClick,delete b.showArrow,delete b.expandIcon,E.createElement("div",(0,P.Z)({},b,{className:$,style:a,id:o}),E.createElement("div",k,this.renderIcon(),this.renderTitle(),I&&E.createElement("div",{className:"".concat(c,"-extra")},f)),E.createElement(_.Z,(0,P.Z)({visible:s,leavedClassName:"".concat(c,"-content-hidden")},v,{forceRender:p,removeOnLeave:d}),(function(e,n){var t=e.className,r=e.style;return E.createElement(A,{ref:n,prefixCls:c,className:t,style:r,isActive:s,forceRender:p,role:u?"tabpanel":null},l)})))}}]),t}(E.Component);T.defaultProps={showArrow:!0,isActive:!1,onItemClick:function(){},headerClass:"",forceRender:!1};var R=T;function M(e){var n=e;if(!Array.isArray(n)){var t=(0,k.Z)(n);n="number"===t||"string"===t?[n]:[]}return n.map((function(e){return String(e)}))}var B=function(e){(0,C.Z)(t,e);var n=(0,$.Z)(t);function t(e){var r;(0,g.Z)(this,t),(r=n.call(this,e)).onClickItem=function(e){var n=r.state.activeKey;if(r.props.accordion)n=n[0]===e?[]:[e];else{var t=(n=(0,b.Z)(n)).indexOf(e);t>-1?n.splice(t,1):n.push(e)}r.setActiveKey(n)},r.getNewChild=function(e,n){if(!e)return null;var t=r.state.activeKey,o=r.props,a=o.prefixCls,c=o.openMotion,i=o.accordion,l=o.destroyInactivePanel,s=o.expandIcon,d=o.collapsible,u=e.key||String(n),p=e.props,v=p.header,f=p.headerClass,m=p.destroyInactivePanel,h=p.collapsible,x=null!=h?h:d,b={key:u,panelKey:u,header:v,headerClass:f,isActive:i?t[0]===u:t.indexOf(u)>-1,prefixCls:a,destroyInactivePanel:null!=m?m:l,openMotion:c,accordion:i,children:e.props.children,onItemClick:"disabled"===x?null:r.onClickItem,expandIcon:s,collapsible:x};return"string"==typeof e.type?e:(Object.keys(b).forEach((function(e){void 0===b[e]&&delete b[e]})),E.cloneElement(e,b))},r.getItems=function(){var e=r.props.children;return(0,I.Z)(e).map(r.getNewChild)},r.setActiveKey=function(e){"activeKey"in r.props||r.setState({activeKey:e}),r.props.onChange(r.props.accordion?e[0]:e)};var o=e.activeKey,a=e.defaultActiveKey;return"activeKey"in e&&(a=o),r.state={activeKey:M(a)},r}return(0,y.Z)(t,[{key:"shouldComponentUpdate",value:function(e,n){return!w()(this.props,e)||!w()(this.state,n)}},{key:"render",value:function(){var e,n=this.props,t=n.prefixCls,r=n.className,o=n.style,a=n.accordion,c=h()((e={},(0,x.Z)(e,t,!0),(0,x.Z)(e,r,!!r),e));return E.createElement("div",{className:c,style:o,role:a?"tablist":null},this.getItems())}}],[{key:"getDerivedStateFromProps",value:function(e){var n={};return"activeKey"in e&&(n.activeKey=M(e.activeKey)),n}}]),t}(E.Component);B.defaultProps={prefixCls:"rc-collapse",onChange:function(){},accordion:!1,destroyInactivePanel:!1},B.Panel=R;var O=B,H=(B.Panel,t(1352)),G=t(38138),z=t(71),D=t(38892);var L=e=>{const{getPrefixCls:n}=E.useContext(G.E_),{prefixCls:t,className:r="",showArrow:o=!0}=e,a=n("collapse",t),c=h()({[`${a}-no-arrow`]:!o},r);return E.createElement(O.Panel,Object.assign({},e,{prefixCls:a,className:c}))},q=t(84917),W=t(92802),U=t(79468),F=t(15030);const X=e=>{const{componentCls:n,collapseContentBg:t,padding:r,collapseContentPaddingHorizontal:o,collapseHeaderBg:a,collapseHeaderPadding:c,collapsePanelBorderRadius:i,lineWidth:l,lineType:s,colorBorder:d,colorText:u,colorTextHeading:p,colorTextDisabled:v,fontSize:f,lineHeight:m,marginSM:h,paddingSM:x,motionDurationSlow:b,fontSizeIcon:g}=e,y=`${l}px ${s} ${d}`;return{[n]:Object.assign(Object.assign({},(0,F.Wf)(e)),{backgroundColor:a,border:y,borderBottom:0,borderRadius:`${i}px`,"&-rtl":{direction:"rtl"},[`& > ${n}-item`]:{borderBottom:y,"&:last-child":{[`\n            &,\n            & > ${n}-header`]:{borderRadius:`0 0 ${i}px ${i}px`}},[`> ${n}-header`]:{position:"relative",display:"flex",flexWrap:"nowrap",alignItems:"flex-start",padding:c,color:p,lineHeight:m,cursor:"pointer",transition:`all ${b}, visibility 0s`,[`> ${n}-header-text`]:{flex:"auto"},"&:focus":{outline:"none"},[`${n}-expand-icon`]:{height:f*m,display:"flex",alignItems:"center",paddingInlineEnd:h},[`${n}-arrow`]:Object.assign(Object.assign({},(0,F.Ro)()),{fontSize:g,svg:{transition:`transform ${b}`}}),[`${n}-header-text`]:{marginInlineEnd:"auto"}},[`${n}-header-collapsible-only`]:{cursor:"default",[`${n}-header-text`]:{flex:"none",cursor:"pointer"}},[`${n}-icon-collapsible-only`]:{cursor:"default",[`${n}-expand-icon`]:{cursor:"pointer"}},[`&${n}-no-arrow`]:{[`> ${n}-header`]:{paddingInlineStart:x}}},[`${n}-content`]:{color:u,backgroundColor:t,borderTop:y,[`& > ${n}-content-box`]:{padding:`${r}px ${o}px`},"&-hidden":{display:"none"}},[`${n}-item:last-child`]:{[`> ${n}-content`]:{borderRadius:`0 0 ${i}px ${i}px`}},[`& ${n}-item-disabled > ${n}-header`]:{"\n          &,\n          & > .arrow\n        ":{color:v,cursor:"not-allowed"}},[`&${n}-icon-position-end`]:{[`& > ${n}-item`]:{[`> ${n}-header`]:{[`${n}-expand-icon`]:{order:1,paddingInlineEnd:0,paddingInlineStart:h}}}}})}},J=e=>{const{componentCls:n}=e;return{[`${n}-rtl`]:{[`> ${n}-item > ${n}-header ${n}-arrow svg`]:{transform:"rotate(180deg)"}}}},V=e=>{const{componentCls:n,collapseHeaderBg:t,paddingXXS:r,colorBorder:o}=e;return{[`${n}-borderless`]:{backgroundColor:t,border:0,[`> ${n}-item`]:{borderBottom:`1px solid ${o}`},[`\n        > ${n}-item:last-child,\n        > ${n}-item:last-child ${n}-header\n      `]:{borderRadius:0},[`> ${n}-item:last-child`]:{borderBottom:0},[`> ${n}-item > ${n}-content`]:{backgroundColor:"transparent",borderTop:0},[`> ${n}-item > ${n}-content > ${n}-content-box`]:{paddingTop:r}}}},Y=e=>{const{componentCls:n,paddingSM:t}=e;return{[`${n}-ghost`]:{backgroundColor:"transparent",border:0,[`> ${n}-item`]:{borderBottom:0,[`> ${n}-content`]:{backgroundColor:"transparent",border:0,[`> ${n}-content-box`]:{paddingBlock:t}}}}}};var Q=(0,W.Z)("Collapse",(e=>{const n=(0,U.TS)(e,{collapseContentBg:e.colorBgContainer,collapseHeaderBg:e.colorFillAlter,collapseHeaderPadding:`${e.paddingSM}px ${e.padding}px`,collapsePanelBorderRadius:e.borderRadiusLG,collapseContentPaddingHorizontal:16});return[X(n),V(n),Y(n),J(n),(0,q.Z)(n)]}));const ee=e=>{const{getPrefixCls:n,direction:t}=E.useContext(G.E_),{prefixCls:r,className:o="",bordered:a=!0,ghost:c,expandIconPosition:i="start"}=e,l=n("collapse",r),s=n(),[d,u]=Q(l),p=E.useMemo((()=>"left"===i?"start":"right"===i?"end":i),[i]),v=h()(`${l}-icon-position-${p}`,{[`${l}-borderless`]:!a,[`${l}-rtl`]:"rtl"===t,[`${l}-ghost`]:!!c},o,u),m=Object.assign(Object.assign({},(0,z.ZP)(s)),{motionAppear:!1,leavedClassName:`${l}-content-hidden`});return d(E.createElement(O,Object.assign({openMotion:m},e,{expandIcon:function(){let n=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};const{expandIcon:t}=e,r=t?t(n):E.createElement(f.Z,{rotate:n.isActive?90:void 0});return(0,D.Tm)(r,(()=>({className:h()(r.props.className,`${l}-arrow`)})))},prefixCls:l,className:v}),(()=>{const{children:n}=e;return(0,I.Z)(n).map(((e,n)=>{var t;if(null===(t=e.props)||void 0===t?void 0:t.disabled){const t=e.key||String(n),{disabled:r,collapsible:o}=e.props,a=Object.assign(Object.assign({},(0,H.Z)(e.props,["disabled"])),{key:t,collapsible:null!=o?o:r?"disabled":void 0});return(0,D.Tm)(e,a)}return e}))})()))};ee.Panel=L;var ne=ee,te=t(51955),re=t(92662),oe=t(11511),ae=t(61388),ce="service___HEVHW";function ie(e){return le.apply(this,arguments)}function le(){return(le=d()(l()().mark((function e(n){return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,ae.request)("/api/v1/service/".concat(n.namespace),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function se(e){return de.apply(this,arguments)}function de(){return(de=d()(l()().mark((function e(n){return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,ae.request)("/api/v1/service/".concat(n.namespace,"/").concat(n.name),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function ue(e){return pe.apply(this,arguments)}function pe(){return(pe=d()(l()().mark((function e(n){return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,ae.request)("/api/v1/service/".concat(n.namespace,"/").concat(n.name,"/").concat(n.version),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function ve(e){return fe.apply(this,arguments)}function fe(){return(fe=d()(l()().mark((function e(n){return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,ae.request)("/api/v1/service/".concat(n.namespace,"/").concat(n.name,"/").concat(n.version,"/").concat(n.addr),{method:"GET"}));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var me=t(62086),he=ne.Panel,xe=function(){var e=(0,E.useState)([]),n=p()(e,2),t=n[0],r=n[1],a=(0,E.useState)({}),i=p()(a,2),s=i[0],u=i[1],f=(0,E.useState)({}),m=p()(f,2),h=m[0],x=m[1],b=(0,E.useState)({}),g=p()(b,2),y=g[0],C=g[1],$=(0,ae.useModel)("useUser").select,k=JSON.parse(localStorage.getItem("use-local-storage-state-namespace")||"{}"),I=(0,E.useState)([]),Z=p()(I,2),w=Z[0],P=Z[1],S=(0,E.useState)([]),_=p()(S,2),N=_[0],j=_[1],A=(0,E.useState)([]),K=p()(A,2),T=K[0],R=K[1],M=function(){var e=d()(l()().mark((function e(){var n;return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,ie({namespace:null==k?void 0:k.value});case 2:0===(null==(n=e.sent)?void 0:n.code)?r((null==n?void 0:n.data)||[]):te.ZP.error("服务列表数据获取失败");case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();(0,E.useEffect)((function(){k&&M()}),[$]);var B=function(){var e=d()(l()().mark((function e(n){var t;return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(null==w||!w.includes(n)){e.next=2;break}return e.abrupt("return");case 2:return e.next=4,se({namespace:null==k?void 0:k.value,name:n});case 4:if(0!==(null==(t=e.sent)?void 0:t.code)){e.next=9;break}u(c()(c()({},s),{},o()({},n,(null==t?void 0:t.data)||[]))),e.next=11;break;case 9:return te.ZP.error("服务版本数据获取失败"),e.abrupt("return",null);case 11:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}(),O=function(){var e=d()(l()().mark((function e(n,t,r){var a,i;return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(r.stopPropagation(),a="".concat(n,"_").concat(t),null==N||!N.includes(a)){e.next=4;break}return e.abrupt("return");case 4:return e.next=6,ue({namespace:null==k?void 0:k.value,name:n,version:t});case 6:if(0!==(null==(i=e.sent)?void 0:i.code)){e.next=11;break}x(c()(c()({},h),{},o()({},a,(null==i?void 0:i.data)||[]))),e.next=13;break;case 11:return te.ZP.error("服务入口数据获取失败"),e.abrupt("return",null);case 13:case"end":return e.stop()}}),e)})));return function(n,t,r){return e.apply(this,arguments)}}(),H=function(){var e=d()(l()().mark((function e(n,t,r,a){var i,s,d,u,p,v,f;return l()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(a.stopPropagation(),i="".concat(n,"_").concat(t,"_").concat(r),null==T||!T.includes(i)){e.next=4;break}return e.abrupt("return");case 4:return e.next=6,ve({namespace:null==k?void 0:k.value,name:n,version:t,addr:r});case 6:if(0!==(null==(s=e.sent)?void 0:s.code)){e.next=14;break}p=null==s||null===(d=s.data)||void 0===d?void 0:d.client_info,v=null==s||null===(u=s.data)||void 0===u?void 0:u.service_info,f=[{id:1,col1:"空间名称",col2:null==v?void 0:v.namespace,col3:"ACCESS_KEY",col4:null==p?void 0:p.register_access_key},{id:2,col1:"服务名称",col2:null==v?void 0:v.service_name,col3:"IP",col4:null==p?void 0:p.register_ip},{id:3,col1:"版本",col2:null==v?void 0:v.service_version,col3:"HOST",col4:null==p?void 0:p.register_host},{id:4,col1:"服务入口",col2:null==v?void 0:v.service_endpoint,col3:"注册时间",col4:null==p?void 0:p.register_time}],C(c()(c()({},y),{},o()({},i,f))),e.next=16;break;case 14:return te.ZP.error("服务详情信息数据获取失败"),e.abrupt("return",null);case 16:case"end":return e.stop()}}),e)})));return function(n,t,r,o){return e.apply(this,arguments)}}(),G=[{dataIndex:"col1",key:"col1"},{dataIndex:"col2",key:"col2"},{dataIndex:"col3",key:"col3"},{dataIndex:"col4",key:"col4"}];return(0,me.jsx)("div",{className:ce,children:(0,me.jsx)(v._z,{header:{title:"服务管理"},children:(null==t?void 0:t.length)>0?(0,me.jsx)(ne,{bordered:!1,onChange:function(e){return P(e)},children:t.map((function(e){var n,t;return(0,me.jsx)(he,{header:e,onClick:function(){return B(e)},children:(null===(n=s[e])||void 0===n?void 0:n.length)>0?(0,me.jsx)(ne,{bordered:!1,onChange:function(e){return j(e)},children:null===(t=s[e])||void 0===t?void 0:t.map((function(n){var t,r;return(0,me.jsx)(he,{header:n,onClick:function(t){return O(e,n,t)},children:(null===(t=h[e])||void 0===t?void 0:t.length)>0?(0,me.jsx)(ne,{bordered:!1,onChange:function(e){R(e)},children:null===(r=h["".concat(e,"_").concat(n)])||void 0===r?void 0:r.map((function(t){return(0,me.jsx)(he,{header:t,onClick:function(r){return H(e,n,t,r)},children:(0,me.jsx)(re.Z,{rowKey:"id",pagination:!1,bordered:!0,columns:G,dataSource:y["".concat(e,"_").concat(n,"_").concat(t)]})},"".concat(e,"_").concat(n,"_").concat(t))}))}):(0,me.jsx)(oe.Z,{image:oe.Z.PRESENTED_IMAGE_SIMPLE})},"".concat(e,"_").concat(n))}))}):(0,me.jsx)(oe.Z,{image:oe.Z.PRESENTED_IMAGE_SIMPLE})},e)}))}):(0,me.jsx)(oe.Z,{image:oe.Z.PRESENTED_IMAGE_SIMPLE})})})}}}]);