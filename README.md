Micro Mesh
=====

<a href="#">![micro-mesh](/doc/img/micro-mesh.jpg "micro-mesh")</a>

### TODO

- [hb-go/grpc-gateway](https://github.com/hb-go/grpc-gateway)
    - [x] `gen-grpc-gateway`扩展，支持gRPC服务本地调用，在service中启动http server
    - [ ] `gen-istio-gateway`通过grpc-gateway API自动生成istio gateway的`.yaml`配置
- Istio部署
    - [ ] k8s`.yaml`脚本
    - [ ] `JWT`Gateway认证
    - [ ] `RBAC`服务间访问控制
- 服务
    - [ ] 同步&异步调用
    - [ ] 配置中心

### Protoc

##### 安装

```bash
go get -u github.com/hb-go/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

##### 代码生成

```bash
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -Ithird_party/googleapis \
  --go_out=plugins=grpc:. \
  proto/*.proto

# grpc-gateway 
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -Ithird_party/googleapis \
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/gateway.yaml:. \
  proto/*.proto

# swagger  
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -Ithird_party/googleapis \
  --swagger_out=logtostderr=true,grpc_api_configuration=proto/gateway.yaml:. \
  proto/*.proto
```