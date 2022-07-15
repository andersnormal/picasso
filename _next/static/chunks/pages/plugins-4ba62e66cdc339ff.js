(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[57],{5001:function(e,n,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/plugins",function(){return t(3578)}])},7845:function(e,n,t){"use strict";var s=t(5893);n.Z={github:"https://github.com/andersnormal/picasso",docsRepositoryBase:"https://github.com/andersnormal/picasso/blob/main/docs/pages",titleSuffix:" \u2013 Picasso",nextLinks:!0,prevLinks:!0,search:!0,customSearch:null,darkMode:!0,footer:!0,footerText:"MIT ".concat((new Date).getFullYear()," \xa9 Sebastian Doell (@katallaxie)."),footerEditLink:"Edit this page on GitHub",logo:(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)("svg",{children:"..."}),(0,s.jsx)("span",{children:"Picasso"})]})}},3578:function(e,n,t){"use strict";t.r(n);var s=t(5893),o=t(7829),r=t.n(o),i=t(3805),a=t(7845),c=(t(5675),function(e){return(0,i.withSSG)(r()({filename:"plugins.mdx",route:"/plugins",meta:{},pageMap:[{name:"index",route:"/",frontMatter:{title:"Picasso - A versatile task runner"}},{name:"installation",route:"/installation",frontMatter:{title:"Installation"}},{name:"meta.json",meta:{index:{title:"Introduction",type:"page",hidden:!0}}},{name:"plugins",route:"/plugins"},{name:"reference",route:"/reference"},{name:"usage",route:"/usage"}]},a.Z))(e)});function p(e){var n=Object.assign({h1:"h1",p:"p",a:"a",code:"code",pre:"pre"},e.components);return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.h1,{children:"Plugins"}),"\n",(0,s.jsxs)(n.p,{children:["Picasso is highly extensible via its plugin system. It uses the ",(0,s.jsx)(n.a,{href:"https://github.com/hashicorp/go-plugin",children:"golang plugin system over RPC"})," to add any new feature imaginable and to keep the core system lightweight."]}),"\n",(0,s.jsxs)(n.p,{children:["The rpc system uses gRPC to establish the communication between the command line tool and the plugins. The ",(0,s.jsx)(n.a,{href:"https://raw.githubusercontent.com/andersnormal/picasso/main/pkg/proto/plugin.proto",children:"protocol"})," is expressed in Google's protocol buffers."]}),"\n",(0,s.jsxs)(n.p,{children:["A plugin is created by importing the ",(0,s.jsx)(n.code,{children:"github.com/andersnormal/picasso/pkg/plugin"})," and ",(0,s.jsx)(n.code,{children:"github.com/andersnormal/picasso/pkg/proto"})," package and implementing the ",(0,s.jsx)(n.code,{children:"proto.PluginServer"})," interface."]}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-go",children:'package main\n\nimport (\n\t"context"\n\t"fmt"\n\n\t"github.com/andersnormal/picasso/pkg/plugin"\n\t"github.com/andersnormal/picasso/pkg/proto"\n)\n\nfunc main() {\n\tplugin.Serve(&plugin.ServeOpts{\n\t\tGRPCPluginFunc: func() proto.PluginServer {\n\t\t\treturn &server{}\n\t\t},\n\t})\n}\n\ntype server struct {\n\tproto.UnimplementedPluginServer\n}\n\n// Start ...\nfunc (s *server) Execute(ctx context.Context, req *proto.Execute_Request) (*proto.Execute_Response, error) {\n\treturn &proto.Execute_Response{}, nil\n}\n\n// Stop ...\nfunc (s *server) Stop(ctx context.Context, req *proto.Stop_Request) (*proto.Stop_Response, error) {\n\treturn &proto.Stop_Response{}, nil\n}\n'})}),"\n",(0,s.jsxs)(n.p,{children:["Any set ",(0,s.jsx)(n.code,{children:"--timeout"})," is enforced by the CLI, thus plugins are stopped if the set time elapses."]})]})}n.default=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return(0,s.jsx)(c,Object.assign({},e,{children:(0,s.jsx)(p,e)}))}}},function(e){e.O(0,[58,774,888,179],(function(){return n=5001,e(e.s=n);var n}));var n=e.O();_N_E=n}]);