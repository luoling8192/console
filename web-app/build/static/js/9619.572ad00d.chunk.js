"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[9619],{9619:(e,t,n)=>{n.r(t),n.d(t,{default:()=>E});var i=n(65043),s=n(94574),r=n(33097),a=n.n(r),o=n(73216),l=n(89923),c=n(56483),d=n(20554),u=n(22166),x=n(64159),h=n(23758),m=n(6681),b=n(57676),j=n(43596),g=n(70579);const p=e=>{let{hasErrors:t}=e;const n=(0,d.jL)(),i=(0,u.d4)((e=>e.addBucket.name));return(0,g.jsx)(l.cl_,{id:"bucket-name",name:"bucket-name",error:t?"Invalid bucket name":"",onFocus:()=>{n((0,b.Xb)(!0))},onChange:e=>{n((0,b.i)(e.target.value))},label:"Bucket Name",value:i,required:!0})};var f=n(99161),k=n(77938);const v=e=>{let{ruleText:t}=e;return(0,g.jsx)(i.Fragment,{children:(0,g.jsxs)(l.xA9,{container:!0,style:{display:"flex",justifyContent:"flex-start"},children:[(0,g.jsx)(l.xA9,{item:!0,xs:1,children:(0,g.jsx)(l.$rg,{width:"16px",height:"16px",style:{color:"#18BF42"}})}),(0,g.jsx)(l.xA9,{item:!0,xs:9,sx:{color:"#8f949c",display:"flex",justifyContent:"flex-start"},children:t})]})})},T=e=>{let{ruleText:t}=e;return(0,g.jsx)(i.Fragment,{children:(0,g.jsxs)(l.xA9,{container:!0,sx:{color:"#C83B51",display:"flex",justifyContent:"flex-start"},children:[(0,g.jsx)(l.xA9,{item:!0,xs:1,sx:{paddingRight:1},children:(0,g.jsx)(l.xWY,{width:"16px",height:"16px"})}),(0,g.jsx)(l.xA9,{item:!0,xs:9,sx:{color:"#C83B51",display:"flex",justifyContent:"flex-start",paddingLeft:1},children:t})]})})},y=e=>{let{ruleText:t}=e;return(0,g.jsx)(i.Fragment,{children:(0,g.jsxs)(l.xA9,{container:!0,sx:{display:"flex",justifyContent:"flex-start"},children:[(0,g.jsx)(l.xA9,{item:!0,xs:1,children:(0,g.jsx)(l.GQ2,{width:"12px",height:"12px",style:{color:"#8f949c"}})}),(0,g.jsx)(l.xA9,{item:!0,xs:9,sx:{color:"#8f949c",display:"flex",justifyContent:"flex-start"},style:{},children:t})]})})},C=e=>{let{errorList:t}=e;const n="Bucket names must be between 3 (min) and 63 (max) characters long.",s="Bucket names can consist only of lowercase letters, numbers, dots (.), and hyphens (-).",r="Bucket names must not contain two adjacent periods, or a period adjacent to a hyphen.",a="Bucket names must not be formatted as an IP address (for example, 192.168.5.4).",o="Bucket names must not start with the prefix xn--.",c="Bucket names must not end with the suffix -s3alias. This suffix is reserved for access point alias names.",d="Bucket names must be unique within a partition.",x=(0,u.d4)((e=>e.addBucket.name)),[h,m]=(0,i.useState)(!1),b=(0,u.d4)((e=>e.addBucket.loading)),[j,p,f,k,C,B,_]=t;return(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)(l.J2w,{id:"toggle-naming-rules",type:"button",open:h,label:"".concat(h?"Hide":"View"," Bucket Naming Rules"),onClick:()=>{m(!h)}}),h&&(0,g.jsxs)(l.xA9,{container:!0,sx:{fontSize:14,paddingTop:12},children:[(0,g.jsxs)(l.xA9,{item:!0,xs:6,children:[0===x.length?(0,g.jsx)(y,{ruleText:n}):j?(0,g.jsx)(v,{ruleText:n}):(0,g.jsx)(T,{ruleText:n}),0===x.length?(0,g.jsx)(y,{ruleText:s}):p?(0,g.jsx)(v,{ruleText:s}):(0,g.jsx)(T,{ruleText:s}),0===x.length?(0,g.jsx)(y,{ruleText:r}):f?(0,g.jsx)(v,{ruleText:r}):(0,g.jsx)(T,{ruleText:r}),0===x.length?(0,g.jsx)(y,{ruleText:a}):k?(0,g.jsx)(v,{ruleText:a}):(0,g.jsx)(T,{ruleText:a})]}),(0,g.jsxs)(l.xA9,{item:!0,xs:6,children:[0===x.length?(0,g.jsx)(y,{ruleText:o}):C?(0,g.jsx)(v,{ruleText:o}):(0,g.jsx)(T,{ruleText:o}),0===x.length?(0,g.jsx)(y,{ruleText:c}):B?(0,g.jsx)(v,{ruleText:c}):(0,g.jsx)(T,{ruleText:c}),0===x.length?(0,g.jsx)(y,{ruleText:d}):_?(0,g.jsx)(v,{ruleText:d}):(0,g.jsx)(T,{ruleText:d})]})]}),b&&(0,g.jsx)(l.xA9,{item:!0,xs:12,children:(0,g.jsx)(l.z21,{})})]})};var B=n(48793),_=n(56629),O=n(53518),S=n(92452),w=n(12848);const A=s.Ay.div((e=>{let{theme:t}=e;return{color:a()(t,"signalColors.danger","#C51B3F"),border:"1px solid ".concat(a()(t,"signalColors.danger","#C51B3F")),padding:8,borderRadius:3}})),E=()=>{const e=(0,d.jL)(),t=(0,o.Zp)(),n=new RegExp("^[a-z0-9][a-z0-9\\.\\-]{1,61}[a-z0-9]$"),s=new RegExp("^(\\d+\\.){3}\\d+$"),r=(0,u.d4)((e=>e.addBucket.name)),a=(0,u.d4)((e=>e.addBucket.isDirty)),[v,T]=(0,i.useState)([]),y=v.filter((e=>!e)).length>0,[E,F]=(0,i.useState)([]),V=(0,u.d4)((e=>e.addBucket.versioningEnabled)),I=(0,u.d4)((e=>e.addBucket.excludeFolders)),N=(0,u.d4)((e=>e.addBucket.excludedPrefixes)),P=(0,u.d4)((e=>e.addBucket.lockingEnabled)),U=(0,u.d4)((e=>e.addBucket.quotaEnabled)),R=(0,u.d4)((e=>e.addBucket.quotaSize)),q=(0,u.d4)((e=>e.addBucket.quotaUnit)),z=(0,u.d4)((e=>e.addBucket.retentionEnabled)),L=(0,u.d4)((e=>e.addBucket.retentionMode)),G=(0,u.d4)((e=>e.addBucket.retentionUnit)),M=(0,u.d4)((e=>e.addBucket.retentionValidity)),J=(0,u.d4)((e=>e.addBucket.loading)),K=(0,u.d4)((e=>e.addBucket.invalidFields)),Y=(0,u.d4)((e=>e.addBucket.lockingFieldDisabled)),D=(0,u.d4)(x.Rq),$=(0,u.d4)(x.nM),Q=(0,u.d4)((e=>e.addBucket.navigateTo)),W=(0,k._)("*",[f.OV.S3_PUT_BUCKET_VERSIONING,f.OV.S3_PUT_BUCKET_OBJECT_LOCK_CONFIGURATION,f.OV.S3_PUT_ACTIONS],!0),H=(0,k._)("*",[f.OV.S3_PUT_BUCKET_VERSIONING,f.OV.S3_PUT_ACTIONS]);(0,i.useEffect)((()=>{const e=[!(a&&(r.length<3||r.length>63)),n.test(r),!(r.includes(".-")||r.includes("-.")||r.includes("..")),!s.test(r),!r.startsWith("xn--"),!r.endsWith("-s3alias"),!E.includes(r)];T(e)}),[r,a]),(0,i.useEffect)((()=>{e((0,b.i)("")),e((0,b.Xb)(!1));_.F.buckets.listBuckets().then((t=>{if(t.data){var n=[];null!=t.data.buckets&&t.data.buckets.length>0&&t.data.buckets.forEach((e=>{n.push(e.name)})),F(n)}else t.error&&e((0,x.C9)((0,O.S)(t.error)))})).catch((t=>{e((0,x.C9)((0,O.S)(t)))}))}),[e]);return(0,i.useEffect)((()=>{if(""!==Q){const n="".concat(Q);e((0,b.E2)()),t(n)}}),[Q,t,e]),(0,i.useEffect)((()=>{e((0,x.ph)("add_bucket"))}),[]),(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)(B.A,{label:(0,g.jsx)(l.EGL,{label:"Buckets",onClick:()=>t("/buckets")}),actions:(0,g.jsx)(S.A,{})}),(0,g.jsx)(l.Mxu,{children:(0,g.jsx)(l.Hbc,{title:"Create Bucket",icon:(0,g.jsx)(l.brV,{}),helpBox:(0,g.jsx)(l.lVp,{iconComponent:(0,g.jsx)(l.brV,{}),title:"Buckets",help:(0,g.jsxs)(i.Fragment,{children:["FST uses buckets to organize objects. A bucket is similar to a folder or directory in a filesystem, where each bucket can hold an arbitrary number of objects.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("b",{children:"Versioning"})," allows to keep multiple versions of the same object under the same key.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("b",{children:"Object Locking"})," prevents objects from being deleted. Required to support retention and legal hold. Can only be enabled at bucket creation.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("b",{children:"Quota"})," limits the amount of data in the bucket.",W&&(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("b",{children:"Retention"})," imposes rules to prevent object deletion for a period of time. Versioning must be enabled in order to set bucket retention policies."]}),(0,g.jsx)("br",{}),(0,g.jsx)("br",{})]})}),children:(0,g.jsxs)("form",{noValidate:!0,autoComplete:"off",onSubmit:t=>{t.preventDefault(),e((0,j._)())},children:[(0,g.jsxs)(l.azJ,{children:[(0,g.jsx)(p,{hasErrors:y}),(0,g.jsx)(l.azJ,{sx:{margin:"10px 0"},children:(0,g.jsx)(C,{errorList:v})}),(0,g.jsx)(l._xt,{separator:!0,children:"Features"}),(0,g.jsxs)(l.azJ,{sx:{marginTop:10},children:[!D&&(0,g.jsxs)(i.Fragment,{children:[(0,g.jsxs)(A,{children:["These features are unavailable in a single-disk setup.",(0,g.jsx)("br",{}),"Please deploy a server in"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/linux/operations/install-deploy-manage/deploy-minio-multi-node-multi-drive.html?ref=con",target:"_blank",rel:"noopener",children:"Distributed Mode"})," ","to use these features."]}),(0,g.jsx)("br",{}),(0,g.jsx)("br",{})]}),$.enabled&&(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)("br",{}),(0,g.jsxs)(l.azJ,{withBorders:!0,sx:{display:"flex",alignItems:"center",padding:"10px","& > .min-icon ":{width:20,height:20,marginRight:10}},children:[(0,g.jsx)(l.mo0,{})," Versioning setting cannot be changed as cluster replication is enabled for this site."]}),(0,g.jsx)("br",{})]}),(0,g.jsx)(l.dOG,{value:"versioned",id:"versioned",name:"versioned",checked:V,onChange:t=>{e((0,b.tr)(t.target.checked))},label:"Versioning",disabled:!D||P||$.enabled||!H,tooltip:H?"":(0,f.vj)([f.OV.S3_PUT_BUCKET_VERSIONING,f.OV.S3_PUT_ACTIONS],"Versioning"),helpTip:(0,g.jsxs)(i.Fragment,{children:[P&&V&&(0,g.jsxs)("strong",{children:[" ","You must disable Object Locking before Versioning can be disabled ",(0,g.jsx)("br",{})]}),"FST supports keeping multiple"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/kubernetes/upstream/administration/object-management/object-versioning.html#minio-bucket-versioning",target:"blank",children:"versions"})," ","of an object in a single bucket.",(0,g.jsx)("br",{}),"Versioning is required to enable"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management.html#object-retention",target:"blank",children:"Object Locking"})," ","and"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#object-retention-modes",target:"blank",children:"Retention"}),"."]}),helpTipPlacement:"right"}),V&&D&&!P&&(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)(l.dOG,{id:"excludeFolders",label:"Exclude Folders",checked:I,onChange:t=>{e((0,b.uQ)(t.target.checked))},indicatorLabels:["Enabled","Disabled"],helpTip:(0,g.jsxs)(i.Fragment,{children:["You can choose to"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/windows/administration/object-management/object-versioning.html#exclude-folders-from-versioning",children:"exclude folders and prefixes"})," ","from versioning if Object Locking is not enabled.",(0,g.jsx)("br",{}),"FST requires versioning to support replication.",(0,g.jsx)("br",{}),"Objects in excluded prefixes do not replicate to any peer site or remote site."]}),helpTipPlacement:"right"}),(0,g.jsx)(w.A,{elements:N,label:"Excluded Prefixes",name:"excludedPrefixes",onChange:t=>{let n="";n=Array.isArray(t)?t.join(","):t,e((0,b.pw)(n))},withBorder:!0})]}),(0,g.jsx)(l.dOG,{value:"locking",id:"locking",name:"locking",disabled:Y||!D||!W,checked:P,onChange:t=>{e((0,b.SO)(t.target.checked)),t.target.checked&&!$.enabled&&e((0,b.tr)(!0))},label:"Object Locking",tooltip:W?"":(0,f.vj)([f.OV.S3_PUT_BUCKET_VERSIONING,f.OV.S3_PUT_BUCKET_OBJECT_LOCK_CONFIGURATION,f.OV.S3_PUT_ACTIONS],"Locking"),helpTip:(0,g.jsxs)(i.Fragment,{children:[z&&(0,g.jsxs)("strong",{children:[" ","You must disable Retention before Object Locking can be disabled ",(0,g.jsx)("br",{})]}),"You can only enable"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management.html#object-retention",target:"blank",children:"Object Locking"})," ","when first creating a bucket.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("a",{href:"https://min.io/docs/minio/windows/administration/object-management/object-versioning.html#exclude-folders-from-versioning",children:"Exclude folders and prefixes"})," ","options will not be available if this option is enabled."]}),helpTipPlacement:"right"}),(0,g.jsx)(l.dOG,{value:"bucket_quota",id:"bucket_quota",name:"bucket_quota",checked:U,onChange:t=>{e((0,b.N2)(t.target.checked))},label:"Quota",disabled:!D,helpTip:(0,g.jsxs)(i.Fragment,{children:["Setting a"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/linux/reference/minio-mc/mc-quota-set.html",target:"blank",children:"quota"})," ","assigns a hard limit to a bucket beyond which FST does not allow writes."]}),helpTipPlacement:"right"}),U&&D&&(0,g.jsx)(i.Fragment,{children:(0,g.jsx)(l.cl_,{type:"string",id:"quota_size",name:"quota_size",onChange:t=>{e((0,b.A1)(t.target.value))},label:"Capacity",value:R,required:!0,min:"1",overlayObject:(0,g.jsx)(h.A,{id:"quota_unit",onUnitChange:t=>{e((0,b.rS)(t))},unitSelected:q,unitsList:(0,c.l9)(["Ki"]),disabled:!1}),error:K.includes("quotaSize")?"Please enter a valid quota":""})}),V&&D&&W&&(0,g.jsx)(l.dOG,{value:"bucket_retention",id:"bucket_retention",name:"bucket_retention",checked:z,onChange:t=>{e((0,b.VB)(t.target.checked))},label:"Retention",helpTip:(0,g.jsxs)(i.Fragment,{children:["FST supports setting both"," ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#configure-bucket-default-object-retention",target:"blank",children:"bucket-default"})," ","and per-object retention rules.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{})," For per-object retention settings, defer to the documentation for the PUT operation used by your preferred SDK."]}),helpTipPlacement:"right"}),z&&D&&(0,g.jsxs)(i.Fragment,{children:[(0,g.jsx)(l.z6M,{currentValue:L,id:"retention_mode",name:"retention_mode",label:"Mode",onChange:t=>{e((0,b.Og)(t.target.value))},selectorOptions:[{value:"compliance",label:"Compliance"},{value:"governance",label:"Governance"}],helpTip:(0,g.jsxs)(i.Fragment,{children:[" ",(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#minio-object-locking-compliance",target:"blank",children:"Compliance"})," ","lock protects Objects from write operations by all users, including the FST root user.",(0,g.jsx)("br",{}),(0,g.jsx)("br",{}),(0,g.jsx)("a",{href:"https://min.io/docs/minio/macos/administration/object-management/object-retention.html#minio-object-locking-governance",target:"blank",children:"Governance"})," ","lock protects Objects from write operations by non-privileged users."]}),helpTipPlacement:"right"}),(0,g.jsx)(l.cl_,{type:"number",id:"retention_validity",name:"retention_validity",onChange:t=>{e((0,b.VY)(t.target.valueAsNumber))},label:"Validity",value:String(M),required:!0,overlayObject:(0,g.jsx)(h.A,{id:"retention_unit",onUnitChange:t=>{e((0,b.JW)(t))},unitSelected:G,unitsList:[{value:"days",label:"Days"},{value:"years",label:"Years"}],disabled:!1})})]})]})]}),(0,g.jsxs)(l.xA9,{item:!0,xs:12,sx:{display:"flex",justifyContent:"flex-end",alignItems:"center",gap:10,marginTop:15},children:[(0,g.jsx)(l.$nd,{id:"clear",type:"button",variant:"regular",className:"clearButton",onClick:()=>{e((0,b.E2)())},label:"Clear"}),(0,g.jsx)(m.A,{tooltip:K.length>0||!a||y?"You must apply a valid name to the bucket":"",children:(0,g.jsx)(l.$nd,{id:"create-bucket",type:"submit",variant:"callAction",color:"primary",disabled:J||K.length>0||!a||y,label:"Create Bucket"})})]}),J&&(0,g.jsx)(l.xA9,{item:!0,xs:12,children:(0,g.jsx)(l.z21,{})})]})})})]})}},12848:(e,t,n)=>{n.d(t,{A:()=>l});var i=n(65043),s=n(33097),r=n.n(s),a=n(89923),o=n(70579);const l=e=>{let{elements:t,name:n,label:s,tooltip:l="",commonPlaceholder:c="",onChange:d,withBorder:u=!1}=e;const[x,h]=(0,i.useState)([""]),m=(0,i.createRef)();(0,i.useEffect)((()=>{if(1===x.length&&""===x[0]&&t&&""!==t){const e=t.split(",");e.push(""),h(e)}}),[t,x]),(0,i.useEffect)((()=>{if(x.length>1){const e=m.current;e&&e.scrollIntoView(!1)}}),[x,m]);const b=(0,i.useCallback)((e=>{d(e)}),[d]),j=(0,i.useRef)(!0);(0,i.useEffect)((()=>{if(j.current)return void(j.current=!1);const e=x.filter((e=>""!==e.trim())).join(",");b(e)}),[x]);const g=e=>{e.persist();let t=[...x];const n=r()(e.target,"dataset.index","0");t[parseInt(n)]=e.target.value,h(t)},p=x.map(((e,t)=>(0,o.jsx)(a.cl_,{id:"".concat(n,"-").concat(t.toString()),label:"",name:"".concat(n,"-").concat(t.toString()),value:x[t],onChange:g,index:t,placeholder:c,overlayIcon:t===x.length-1?(0,o.jsx)(a.REV,{}):null,overlayAction:()=>{(e=>{if(""!==e[e.length-1].trim()){const t=[...e];t.push(""),h(t)}})(x)}},"csv-multi-".concat(n,"-").concat(t.toString()))));return(0,o.jsx)(i.Fragment,{children:(0,o.jsxs)(a.azJ,{sx:{display:"flex"},className:"inputItem",children:[(0,o.jsxs)(a.l1Y,{sx:{alignItems:"flex-start"},children:[(0,o.jsx)("span",{children:s}),""!==l&&(0,o.jsx)(a.azJ,{sx:{marginLeft:5,display:"flex",alignItems:"center","& .min-icon":{width:13}},children:(0,o.jsx)(a.m_M,{tooltip:l,placement:"top",children:(0,o.jsx)(a.azJ,{className:l,children:(0,o.jsx)(a.NTw,{})})})})]}),(0,o.jsxs)(a.azJ,{withBorders:u,sx:{width:"100%",overflowY:"auto",height:150,position:"relative"},children:[p,(0,o.jsx)("div",{ref:m})]})]})})}},23758:(e,t,n)=>{n.d(t,{A:()=>d});var i=n(65043),s=n(89923),r=n(94574),a=n(33097),o=n.n(a),l=n(70579);const c=r.Ay.button((e=>{let{theme:t}=e;return{border:"1px solid ".concat(o()(t,"borderColor","#E2E2E2")),borderRadius:3,color:o()(t,"secondaryText","#5B5C5C"),backgroundColor:o()(t,"boxBackground","#FBFAFA"),fontSize:12}})),d=e=>{let{id:t,unitSelected:n,unitsList:r,disabled:a=!1,onUnitChange:o}=e;const[d,u]=i.useState(null),x=Boolean(d),h=e=>{u(null),""!==e&&o&&o(e)};return(0,l.jsxs)(i.Fragment,{children:[(0,l.jsx)(c,{id:"".concat(t,"-button"),"aria-controls":"".concat(t,"-menu"),"aria-haspopup":"true","aria-expanded":x?"true":void 0,onClick:e=>{u(e.currentTarget)},disabled:a,type:"button",children:n}),(0,l.jsx)(s.Vey,{id:"upload-main-menu",options:r,selectedOption:"",onSelect:e=>h(e),hideTriggerAction:()=>{h("")},open:x,anchorEl:d,anchorOrigin:"end"})]})}}}]);
//# sourceMappingURL=9619.572ad00d.chunk.js.map
