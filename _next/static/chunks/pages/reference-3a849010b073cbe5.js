(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[187],{5994:function(e,d,s){(window.__NEXT_P=window.__NEXT_P||[]).push(["/reference",function(){return s(1545)}])},7845:function(e,d,s){"use strict";var n=s(5893);d.Z={github:"https://github.com/andersnormal/picasso",docsRepositoryBase:"https://github.com/andersnormal/picasso/blob/main/docs/pages",titleSuffix:" \u2013 Picasso",nextLinks:!0,prevLinks:!0,search:!0,customSearch:null,darkMode:!0,footer:!0,footerText:"MIT ".concat((new Date).getFullYear()," \xa9 Sebastian Doell (@katallaxie)."),footerEditLink:"Edit this page on GitHub",logo:(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)("svg",{children:"..."}),(0,n.jsx)("span",{children:"Picasso"})]})}},1545:function(e,d,s){"use strict";s.r(d);var n=s(5893),r=s(7829),i=s.n(r),c=s(3805),t=s(7845),l=(s(5675),function(e){return(0,c.withSSG)(i()({filename:"reference.mdx",route:"/reference",meta:{},pageMap:[{name:"index",route:"/",frontMatter:{title:"Picasso - A versatile task runner"}},{name:"installation",route:"/installation",frontMatter:{title:"Installation"}},{name:"meta.json",meta:{index:{title:"Introduction",type:"page",hidden:!0}}},{name:"plugins",route:"/plugins"},{name:"reference",route:"/reference"},{name:"usage",route:"/usage"}]},t.Z))(e)});function h(e){var d=Object.assign({h1:"h1",h2:"h2",p:"p",pre:"pre",code:"code",table:"table",thead:"thead",tr:"tr",th:"th",tbody:"tbody",td:"td",h3:"h3",a:"a",blockquote:"blockquote"},e.components);return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(d.h1,{children:"API Reference"}),"\n",(0,n.jsx)(d.h2,{children:"CLI"}),"\n",(0,n.jsx)(d.p,{children:"Picasso has the following command line syntax:"}),"\n",(0,n.jsx)(d.pre,{children:(0,n.jsx)(d.code,{className:"language-bash",children:"picasso [--flags] [tasks...] [-- ARGS...]\n"})}),"\n",(0,n.jsxs)(d.table,{children:[(0,n.jsx)(d.thead,{children:(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.th,{children:"Short"}),(0,n.jsx)(d.th,{children:"Flag"}),(0,n.jsx)(d.th,{children:"Type"}),(0,n.jsx)(d.th,{children:"Default"}),(0,n.jsx)(d.th,{children:"Description"})]})}),(0,n.jsxs)(d.tbody,{children:[(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-c"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--config"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:".picasso.yml"})}),(0,n.jsxs)(d.td,{children:["Config file. Enabled by default. Set to ",(0,n.jsx)(d.code,{children:".picasso.yml"})," or change to the location of your config"]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-w"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--watch"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsxs)(d.td,{children:["Enables watch of the given tasks. This factors in the ",(0,n.jsx)(d.code,{children:"watch"})," config in your ",(0,n.jsx)(d.code,{children:".picasso.yml"})," file."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-f"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--force"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Forces the execution of operations."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-l"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--list"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsxs)(d.td,{children:["Lists the available tasks specified in the ",(0,n.jsx)(d.code,{children:".picasso.yml"})," file."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-v"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--verbose"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Enables verbose logging of runtime information."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-s"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--silent"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Does not log any runtime information."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-d"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--dry"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Does not apply destructive operations."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-p"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--plugin"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["Executes the provided plugin. Passes the CLI arguments via ",(0,n.jsx)(d.code,{children:"--vars"})," and after the ",(0,n.jsx)(d.code,{children:"--"})," to the execution of the plugin."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"-w"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--watch"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsxs)(d.td,{children:["Enables watch of the given tasks. This factors in the ",(0,n.jsx)(d.code,{children:"watch"})," config in your ",(0,n.jsx)(d.code,{children:".picasso.yml"})," file."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--dir"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"."})}),(0,n.jsx)(d.td,{children:"Sets the current working directory. Defaults to the current directory of execution."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--validate"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsxs)(d.td,{children:["Validates the specification file provided via ",(0,n.jsx)(d.code,{children:".picasso.yml"}),"."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--var"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"[]string"})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["Sets the a variable in the format of ",(0,n.jsx)(d.code,{children:"key=value"})]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--init"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsxs)(d.td,{children:["Creates a new ",(0,n.jsx)(d.code,{children:".picasso.yml"})," file at the provided location of ",(0,n.jsx)(d.code,{children:"--config"})," (default: ",(0,n.jsx)(d.code,{children:"./.picasso.yml"}),")"]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"--version"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Prints the current version."})]})]})]}),"\n",(0,n.jsx)(d.h2,{children:"Schema"}),"\n",(0,n.jsx)(d.h3,{children:"Example"}),"\n",(0,n.jsx)(d.pre,{children:(0,n.jsx)(d.code,{className:"language-yaml",children:'spec: 1\nversion: 1.0.0\nauthors:\n  - John Apple <john@example.com>\nhomepage: https://github.com/andersnormal/picasso\nrepository: https://andersnormal.github.io/picasso/\ntasks:\n  test:\n    disabled: true\n    desc: test\n    vars:\n      region: test\n    steps:\n      - \n        id: foo\n        if: {{if .OS == "linux"}}\n        cmd: |\n          echo "Hello World"\n          echo {{.CWD}}\n        timeout-in-seconds: 10\n        continue-on-error: false\n        uses: remote-exec\n        vars:\n          cwd: {{.CWD}}\n        with:\n          region: eu-west-1\n    watch:\n      paths:\n        - pkg/config\n    template:\n      - \n        file: ./examples/config.json.tpl\n        out: ./config.json\n        vars:\n          foo: bar\n  build:\n    default: true\n    deps:\n      - test\n    vars:\n      region: test\n    cmd:\n      - go build\n    watch:\n      paths:\n        - examples\n      ignore:\n        - .gitignore\n        - .picasso.yml\n'})}),"\n",(0,n.jsx)(d.h3,{children:"General"}),"\n",(0,n.jsxs)(d.table,{children:[(0,n.jsx)(d.thead,{children:(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.th,{children:"Attribute"}),(0,n.jsx)(d.th,{children:"Type"}),(0,n.jsx)(d.th,{children:"Default"}),(0,n.jsx)(d.th,{children:"Description"})]})}),(0,n.jsxs)(d.tbody,{children:[(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"spec"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"int"})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["Specification version to be used. The current version is ",(0,n.jsx)(d.code,{children:"1"}),"."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"version"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Version of the application."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"authors"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"[]string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"List of authors."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"homepage"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"URL to the homepage of the application."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"repository"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"URL to the repository of the application."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"vars"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Vars"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Global variables."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"env"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Env"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Global environment."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"tasks"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#task",children:(0,n.jsx)(d.code,{children:"Tasks"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"The task definitions."})]})]})]}),"\n",(0,n.jsx)(d.h3,{children:"Task"}),"\n",(0,n.jsxs)(d.table,{children:[(0,n.jsx)(d.thead,{children:(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.th,{children:"Attribute"}),(0,n.jsx)(d.th,{children:"Type"}),(0,n.jsx)(d.th,{children:"Default"}),(0,n.jsx)(d.th,{children:"Description"})]})}),(0,n.jsxs)(d.tbody,{children:[(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"name"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Name of the task."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"working-dir"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"cwd"})}),(0,n.jsx)(d.td,{children:"Current directort which the task should run in."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"disabled"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Disable the task in execution."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"depends_on"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"DependsOn"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"List of other task this task depends on in execution."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"vars"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Vars"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Variables for this task."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"env"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Env"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Task specific environment."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"watch"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#watch",children:(0,n.jsx)(d.code,{children:"Watch"})})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["Configuration for ",(0,n.jsx)(d.code,{children:"watch"})," flag of the task."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"template"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#template",children:(0,n.jsx)(d.code,{children:"Templates"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Templates to create for this task."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"steps"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#step",children:(0,n.jsx)(d.code,{children:"Steps"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Templates to create for this task."})]})]})]}),"\n",(0,n.jsx)(d.h3,{children:"Steps"}),"\n",(0,n.jsxs)(d.table,{children:[(0,n.jsx)(d.thead,{children:(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.th,{children:"Attribute"}),(0,n.jsx)(d.th,{children:"Type"}),(0,n.jsx)(d.th,{children:"Default"}),(0,n.jsx)(d.th,{children:"Description"})]})}),(0,n.jsxs)(d.tbody,{children:[(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"cmd"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Commands to run in the current working directory."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"working-dir"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"cwd"})}),(0,n.jsx)(d.td,{children:"Current directort which the task should run in."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"vars"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Vars"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Variables for this task."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"env"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Env"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Task specific environment."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"if"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#condition",children:(0,n.jsx)(d.code,{children:"If"})})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"true"})}),(0,n.jsx)(d.td,{children:"Condition to run this step."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"template"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#template",children:(0,n.jsx)(d.code,{children:"Templates"})})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"Templates to create for this task."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"uses"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"string"})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["A ",(0,n.jsx)(d.a,{href:"/plugins",children:"plugin"})," to be run in this step."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"with"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.a,{href:"#variable",children:(0,n.jsx)(d.code,{children:"Vars"})})}),(0,n.jsx)(d.td,{}),(0,n.jsxs)(d.td,{children:["Extra variables for the plugin in the ",(0,n.jsx)(d.code,{children:"uses"})," property."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"depends_on"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"DependsOn"})}),(0,n.jsx)(d.td,{}),(0,n.jsx)(d.td,{children:"List of other task this task depends on in execution."})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"timeout-in-seconds"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"int64"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"math.MaxInt64"})}),(0,n.jsxs)(d.td,{children:["The timeout for the execution of this step. This is borrowed from the ",(0,n.jsx)(d.code,{children:"context"})," timeout."]})]}),(0,n.jsxs)(d.tr,{children:[(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"continue-on-error"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"bool"})}),(0,n.jsx)(d.td,{children:(0,n.jsx)(d.code,{children:"false"})}),(0,n.jsx)(d.td,{children:"Enables to proceed with the next step even if the current step has failed."})]})]})]}),"\n",(0,n.jsxs)(d.blockquote,{children:["\n",(0,n.jsxs)(d.p,{children:["The ",(0,n.jsx)(d.code,{children:"working-dir"})," is set to the current directory."]}),"\n"]})]})}d.default=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return(0,n.jsx)(l,Object.assign({},e,{children:(0,n.jsx)(h,e)}))}}},function(e){e.O(0,[58,774,888,179],(function(){return d=5994,e(e.s=d);var d}));var d=e.O();_N_E=d}]);