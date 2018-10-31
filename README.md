# go-box

An small bootstrap box for golang HTTP project

## layout

```
main.go            - 工程入口
config.json        - 配置文件
api/               - HTTP API
core/              - 业务逻辑
common/            - 通用全局代码
db/                - 数据库
file/              - 文件存储
test/              - 测试目录
web/               - 前端资源

Dockerfile         - Docker 构建文件
control            - 帮助脚本
```

## 使用

0. 切换到GOPATH 目录下，并执行`export GO111MODULE=on` 开启`go mod` 功能
1. 克隆工程 `git clone https://github.com/yangbinnnn/go-box.git`
2. 重命名工程，比如`mv go-box mydemo`
3. 切换到工程目录下，比如`cd mydemo`，修改`config.json` 中`name` 为`mydemo` 以及数据库相关配置
4. 重建module，`./control remodule`，这一步会将`*.go` 中所有`import go-box/xxx` 替换成`import mydemo/xxx`
5. 编译工程，`./control build`
6. 运行工程，`./dist/mydemo`
7. 编译web，`cd web && yarn install && npm run build`
7. 访问站点，默认`http://127.0.0.1:8000`，如果页面正常显示，说明工程构建成功
8. 打包发布，`./control pack`

## 开发指南

- 配置文件
> `common/config.go` 中定义工程所有配置项，并在`config.json` 中配置值，使用JSON 方便解析

- 数据结构定义
> `common/model.go` 中统一定义数据结构

- 业务逻辑
> `core` 中实现工程的业务逻辑，并统一`core/core.go - InitCore` 中初始化业务逻辑相关依赖如数据库表实例等，这样可以方便在`core` 包中使用，适当拆分业务逻辑分散到`core/*.go` 多个文件中，如用户相关的逻辑使用`core/user.go`

- API接口定义
> `api/api.go` 中注册路由，路由的具体函数需要按业务逻辑分散到`api/*.go` 多个文件中，如用户相关的接口使用`api/user.go`

- 工程入口
> `main.go` 为整个工程的入口，主要负责初始化各个模块，建议不要在其中实现复杂的业务逻辑


## 默认路由

- `/` 前端首页
- `/static` 前端静态资源
- `/ws` websocket api
- `/api/xxx` 业务 api


## 其他

- websocket 推送，直接使用`common.WSBroadcast(msg)` 即可，可以推送给所有连接的客户端


## POWERD BY

- [echo](https://echo.labstack.com/) 
- [vue](https://vuejs.org/) 
- [docker](https://www.docker.com/get-started)

## TODO

- [ ] Dockerfile
- [x] 数据库: mongodb
- [ ] 数据库: redis
- [x] Websocket
- [ ] 文件存储: 本地磁盘，对象存储
- [ ] 接口权限验证
- [ ] 定时任务
