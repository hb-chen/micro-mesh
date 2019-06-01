Micro Mesh
=====

<a href="#">![micro-mesh](/doc/img/micro-mesh.jpg "micro-mesh")</a>

**环境**

- OSX
- GKE **1.12.5-gke.10**
- Go **1.11.1**
- Istio **1.1.0**
- protoc **libprotoc 3.6.1**

**目录**

- [TODO](#TODO)
- [Protoc](#Protoc)

### TODO

- [hb-go/grpc-gateway](https://github.com/hb-go/grpc-gateway)
    - [x] `gen-grpc-gateway`扩展，支持gRPC服务本地调用，在service中启动http server
    - [ ] `gen-istio-gateway`通过grpc-gateway API自动生成istio gateway的`.yaml`配置
- Istio部署
    - [x] k8s`.yaml`脚本
    - [x] `JWT`Gateway认证
    - [x] `RBAC`服务间访问控制
    - [x] 自定义[auth-adapter](/examples/adapter/auth)
- 服务
    - [ ] 同步&异步调用
    - [ ] 配置中心
    - [ ] CI/CD

---
### Protoc

##### 安装

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

# 使用fork grpc-gateway的protoc-gen-grpc-gateway
cd $GOPATH/src/github.com/grpc-ecosystem
git clone github.com/hb-go/grpc-gateway
make bin/protoc-gen-grpc-gateway
mv bin/protoc-gen-grpc-gateway $GOPATH/bin/protoc-gen-grpc-gateway
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
