"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[942],{37655:function(e,r,o){var l=o(56453),t=o(30486),n=o(62086),a=o(63502),u=o(93236),s=o(84650),i=o(12961),c=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],d=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],p=u.forwardRef((function(e,r){var o=e.fieldProps,d=e.children,p=e.params,f=e.proFieldProps,m=e.mode,v=e.valueEnum,h=e.request,g=e.showSearch,w=e.options,P=(0,t.Z)(e,c),S=(0,u.useContext)(s.Z);return(0,n.jsx)(i.Z,(0,l.Z)((0,l.Z)({valueEnum:(0,a.h)(v),request:h,params:p,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,l.Z)({options:w,mode:m,showSearch:g,getPopupContainer:S.getPopupContainer},o),ref:r,proFieldProps:f},P),{},{children:d}))})),f=u.forwardRef((function(e,r){var o=e.fieldProps,c=e.children,p=e.params,f=e.proFieldProps,m=e.mode,v=e.valueEnum,h=e.request,g=e.options,w=(0,t.Z)(e,d),P=(0,l.Z)({options:g,mode:m||"multiple",labelInValue:!0,showSearch:!0,showArrow:!1,autoClearSearchValue:!0,optionLabelProp:"label"},o),S=(0,u.useContext)(s.Z);return(0,n.jsx)(i.Z,(0,l.Z)((0,l.Z)({valueEnum:(0,a.h)(v),request:h,params:p,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,l.Z)({getPopupContainer:S.getPopupContainer},P),ref:r,proFieldProps:f},w),{},{children:c}))})),m=p;m.SearchSelect=f,m.displayName="ProFormComponent",r.Z=m},2259:function(e,r,o){o.r(r),o.d(r,{default:function(){return g}});var l=o(46686),t=o.n(l),n=o(37655),a=o(12188),u=o(93236);const s=e=>"function"==typeof e;var i=function(e){const r=(0,u.useRef)(e);r.current=(0,u.useMemo)((()=>e),[e]);const o=(0,u.useRef)();return o.current||(o.current=function(...e){return r.current.apply(this,e)}),o.current};const c=e=>(r,o)=>{const l=(0,u.useRef)(!1);e((()=>()=>{l.current=!1}),[]),e((()=>{if(l.current)return r();l.current=!0}),o)};var d=c(u.useEffect);var p=!("undefined"==typeof window||!window.document||!window.document.createElement);var f,m=(f=()=>p?localStorage:void 0,function(e,r){let o;try{o=f()}catch(e){console.error(e)}function l(){try{const t=null==o?void 0:o.getItem(e);if(t)return l=t,(null==r?void 0:r.deserializer)?null==r?void 0:r.deserializer(l):JSON.parse(l)}catch(e){console.error(e)}var l;return s(null==r?void 0:r.defaultValue)?null==r?void 0:r.defaultValue():null==r?void 0:r.defaultValue}const[t,n]=(0,u.useState)((()=>l()));return d((()=>{n(l())}),[e]),[t,i((l=>{const a=s(l)?l(t):l;if(n(a),(e=>void 0===e)(a))null==o||o.removeItem(e);else try{null==o||o.setItem(e,(e=>(null==r?void 0:r.serializer)?null==r?void 0:r.serializer(e):JSON.stringify(e))(a))}catch(e){console.error(e)}}))]}),v=o(62086),h=function(){var e=(0,a.useModel)("useUser"),r=e.list,o=e.select,l=e.setSelect,s=localStorage.getItem("use-local-storage-state-namespace");s&&"{}"!==s||localStorage.setItem("use-local-storage-state-namespace",JSON.stringify({label:null==o?void 0:o.namespace_alias,value:null==o?void 0:o.namespace_uid}));var i=m("use-local-storage-state-namespace",{defaultValue:{label:null==o?void 0:o.namespace_alias,value:null==o?void 0:o.namespace_uid}}),c=t()(i,2),d=c[0],p=c[1];return(0,u.useEffect)((function(){}),[o]),(0,v.jsx)("div",{style:{marginTop:10,marginLeft:40,marginBottom:"-24px"},children:(0,v.jsx)(n.Z,{allowClear:!1,options:r,width:"xs",style:{width:180},fieldProps:{fieldNames:{label:"namespace_alias",value:"namespace_uid"},value:d.value,onChange:function(e,r){console.log("option",r),l({label:r.label,value:r.value}),p({label:r.label,value:r.value})}}})})},g=function(){return(0,v.jsxs)(v.Fragment,{children:[(0,v.jsx)(h,{}),(0,v.jsx)(a.Outlet,{})]})}}}]);