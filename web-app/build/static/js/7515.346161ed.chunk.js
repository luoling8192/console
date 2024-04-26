"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[7515],{67515:(e,t,n)=>{n.r(t),n.d(t,{default:()=>b});var a=n(65043),i=n(89923),o=n(56629),l=n(85330),r=n(53518),s=n(77403),c=n(64159),d=n(20554),u=n(94141),m=n(70579);const b=e=>{let{open:t,bucketName:n,closeModalAndRefresh:b}=e;const p=(0,d.jL)(),[h,g]=(0,a.useState)(!1),[j,x]=(0,a.useState)(!0),[v,f]=(0,a.useState)(l.BT.Compliance),[y,k]=(0,a.useState)(l.wg.Days),[C,S]=(0,a.useState)(1),[w,_]=(0,a.useState)(!1);return(0,a.useEffect)((()=>{Number.isNaN(C)||(C||1)<1?_(!1):_(!0)}),[C]),(0,a.useEffect)((()=>{j&&o.F.buckets.getBucketRetentionConfig(n).then((e=>{x(!1),f(e.data.mode),S(e.data.validity),k(e.data.unit)})).catch((()=>{x(!1)}))}),[j,n]),(0,m.jsx)(u.A,{title:"Set Retention Configuration",modalOpen:t,onClose:()=>{b()},children:j?(0,m.jsx)(i.aHM,{style:{width:16,height:16}}):(0,m.jsx)("form",{noValidate:!0,autoComplete:"off",onSubmit:e=>{e.preventDefault(),h||(g(!0),o.F.buckets.setBucketRetentionConfig(n,{mode:v||l.BT.Compliance,unit:y||l.wg.Days,validity:C||1}).then((()=>{g(!1),b()})).catch((e=>{g(!1),p((0,c.Dy)((0,r.S)(e.error)))})))},children:(0,m.jsxs)(i.Hbc,{containerPadding:!1,withBorders:!1,children:[(0,m.jsx)(i.z6M,{currentValue:v,id:"retention_mode",name:"retention_mode",label:"Retention Mode",onChange:e=>{f(e.target.value)},selectorOptions:[{value:"compliance",label:"Compliance"},{value:"governance",label:"Governance"}],helpTip:(0,m.jsxs)(a.Fragment,{children:[" ",(0,m.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#minio-object-locking-compliance",target:"blank",children:"Compliance"})," ","lock protects Objects from write operations by all users, including the FST root user.",(0,m.jsx)("br",{}),(0,m.jsx)("br",{}),(0,m.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#minio-object-locking-governance",target:"blank",children:"Governance"})," ","lock protects Objects from write operations by non-privileged users."]}),helpTipPlacement:"right"}),(0,m.jsx)(i.z6M,{currentValue:y,id:"retention_unit",name:"retention_unit",label:"Retention Unit",onChange:e=>{k(e.target.value)},selectorOptions:[{value:"days",label:"Days"},{value:"years",label:"Years"}]}),(0,m.jsx)(i.cl_,{type:"number",id:"retention_validity",name:"retention_validity",onChange:e=>{S(e.target.valueAsNumber)},label:"Retention Validity",value:String(C),required:!0,min:"1"}),(0,m.jsxs)(i.xA9,{item:!0,xs:12,sx:s.Uz.modalButtonBar,children:[(0,m.jsx)(i.$nd,{id:"cancel",type:"button",variant:"regular",disabled:h,onClick:()=>{b()},label:"Cancel"}),(0,m.jsx)(i.$nd,{id:"set",type:"submit",variant:"callAction",color:"primary",disabled:h||!w,label:"Set"})]}),h&&(0,m.jsx)(i.xA9,{item:!0,xs:12,children:(0,m.jsx)(i.z21,{})})]})})})}}}]);
//# sourceMappingURL=7515.346161ed.chunk.js.map
