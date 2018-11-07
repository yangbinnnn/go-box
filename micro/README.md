## Micro Services Support
微服务支持，实践如何最小成本将已有HTTP 服务映射成GRPC 服务

- Layout
```
- registry/ 负责服务注册
- services/ 实现grpc 接口，代理调用`go-box/core`
- test/ 服务测试
- micro.go 微服务入口，负责注册服务，启动GRCP
```

- 安装依赖
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
```

- 生成`.pb.go`
```
protoc -I micro/services/math micro/services/math/math.proto --go_out=plugins=grpc:micro/services/math
```


## Links
- https://github.com/grpc/grpc-go
