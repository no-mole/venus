# README

`@umijs/max` 模板项目，更多功能参考 [Umi Max 简介](https://umijs.org/docs/max/introduce)


### 安装依赖
```
    pnpm
```

### 安装pnpm
```
curl -fsSL https://get.pnpm.io/install.sh | sh -
$ pnpm -v
7.3.0
```

### 启动项目
执行 `pnpm dev` 命令


### 部署发布
执行 `pnpm build` 命令
```
> umi build
event - compiled successfully in 1179 ms (567 modules)
event - build index.html
```

产物默认会生成到 `./dist` 目录下，
```
./dist
├── index.html
├── umi.css
└── umi.js
```
完成构建后，就可以把 `dist` 目录部署到服务器上了。


### 项目脑图
`https://www.processon.com/v/63b7b52d56f6ad4f613db8df`

### 原型地址
`https://modao.cc/app/VV9yNxFwrnwfgqGy2Makuu#screen=slc7a4h4bpsm9ll`

### 参考UI
`https://lanhuapp.com/dashboard/#/item?fid=all&tid=971e8926-908d-41b5-98c5-2bd53e5d919e`


#### git常用的type类别
build	编译相关的修改，例如发布版本、对项目构建或者依赖的改动
chore	其他修改, 比如改变构建流程、或者增加依赖库、工具等
ci	    持续集成修改
docs	文档修改
feat	新特性、新功能
fix	    修改bug
perf	优化相关，比如提升性能、体验
refactor	代码重构
revert	回滚到上一个版本
style	代码格式修改, 注意不是 css 修改
test	测试用例修改
