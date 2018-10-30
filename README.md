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
2. 重命名工程，比如`mv go-box mydemo`，并修改`config.json` 中`name` 为`mydemo`
3. 切换到工程目录下，比如`cd mydemo`
4. 重建module，`./control remodule`，这一步会将`*.go` 中所有`import go-box/xxx` 替换成`import mydemo/xxx`
5. 编译工程，`./control build`
6. 运行工程，`./dist/mydemo`
7. 编译web，`cd web && yarn install && npm run build`
7. 访问站点，默认`http://127.0.0.1:8000`，如果页面正常显示，说明工程构建成功
8. 打包发布，`./control pack`

## POWERD BY

- [echo](https://echo.labstack.com/) 
- [vue](https://vuejs.org/) 
- [docker](https://www.docker.com/get-started)

## TODO
- [ ] Dockerfile
- [ ] 数据库: mongodb，redis
- [ ] 文件存储: 本地磁盘，对象存储
- [ ] 接口权限验证
- [ ] 定时任务
