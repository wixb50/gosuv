# 概述
GoSuv是GO语言重写的类supervisor的一个带web界面的进程管理程序，单文件运行，简单易用，界面美感十足且对用户友好。



## 特点

- [x] web 控制与配置页面
  - [x] 启动, 停止, 日志跟踪, 重载配置
  - [x] 实时查看日志
  - [x] 添加服务
  - [x] 编辑服务
  - [x] 删除服务
  - [x] 可视化查看服务对应的cpu与内存
- [x] HTTP 基本安全认证
- [x] Github webhook
- 

## 使用
* 启动服务

    ```
    gosuv start-server
    ```

查看服务状态

    $ gosuv status
    PROGRAM NAME            STATUS
    test                    running
    test_again              stopped
    
    $ gosuv stop test
    $ gosuv start test

默认端口 11113  本机测试请使用[http://localhost:11313](http://localhost:11313)

![例子](https://github.com/wixb50/gosuv/blob/master/docs/des.gif)

## 配置
默认配置文件都放在 `$HOME/.gosuv/`
    
* 项目文件名 ：     programs.yml
* 服务器配置文件名：    config.yml

验证信息配置

```yml
server:
    servername: locllhost
    httpauth:
      enabled: false
      username: uu
      password: pp
      addr: :11313
client:
  server_url: http://localhost:11313
```

## 默认日志文件位置
`$HOME/.gosuv/log/`

http  RESTFul 接口

获取或更新

`<GET|PUT> /api/programs/:name`

添加新的服务

`POST /api/programs`

删除服务

`DELETE /api/programs/:name`

## 二次开发

```
$ go get # 下载依赖
$ go generate  # 打包静态资源生成代码
$ goreleaser release --skip-publish  # 打包发布
```

## 贡献人
- [Docking](http://miaomia.com)  
- [codeskyblue](https://github.com/codeskyblue/gosuv)
