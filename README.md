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

## 框架

- registry
    - istio
    - go-micro
        - k8s
        - consul
        - etcd
- tracing
    - opentracing
        - jaeger + ES / kafka
- metrics
    - prometheus
        - ES / kafka

### TODO

- [hb-chen/grpc-gateway](https://github.com/hb-chen/grpc-gateway)
    - [x] `gen-grpc-gateway`扩展，支持gRPC服务本地调用，在`service`中启动http server，已在[v1.10.0](https://github.com/grpc-ecosystem/grpc-gateway/releases/tag/v1.10.0)并入社区版本
    - [ ] `gen-istio-gateway`通过grpc-gateway API自动生成istio gateway的`.yaml`配置
    - [ ] `swagger-codegen`
- Istio部署
    - [x] k8s`.yaml`脚本
    - [x] `JWT`Gateway认证
    - [x] `RBAC`服务间访问控制
    - [x] 自定义[auth-adapter](/examples/adapter/auth)
- 服务
    - [x] gRPC ClientConn对象池
    - [ ] 并发
        - [ ] 并发控制`Wait`、`Cancel`
        - [ ] 超时控制`Timeout`
    - [ ] 流处理
    - [ ] 同步&异步调用
    - [ ] 配置中心
    - [ ] CI/CD

---
### Protoc

#### 安装

[envoyproxy/protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate#installation)

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

# 自定义代码生成
# 1.导出grpc.ServiceDesc
# 2.注册中心接口
go get -u github.com/hb-go/grpc-contrib/protoc-gen-hb-grpc
```

***~~使用fork grpc-gateway的protoc-gen-grpc-gateway~~***
```bash
cd $GOPATH/src/github.com/grpc-ecosystem
git clone github.com/hb-chen/grpc-gateway
make bin/protoc-gen-grpc-gateway
mv bin/protoc-gen-grpc-gateway $GOPATH/bin/protoc-gen-grpc-gateway
```

#### 代码生成

```bash
# go+grpc
# grpc-gateway
# swagger
# hb-grpc
protoc -I$GOPATH/src/ -I./ \
--go_out=plugins=grpc:. \
--grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/gateway.yaml:. \
--swagger_out=logtostderr=true,grpc_api_configuration=proto/gateway.yaml:. \
--validate_out=lang=go:. \
--hb-grpc_out=plugins=registry+desc:. \
proto/*.proto
```
