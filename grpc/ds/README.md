# gRPC服务发现

使用google.golang.org/grpc/resolver实现服务发现

- Feature
    - [x] [go-micro/registry](https://github.com/micro/go-micro/tree/master/registry)
        - [x] 服务发现
        - [x] 版本选择
    - [ ] istio
    

## 使用

```go
// 注册Resolver Builder
r := grpcds.NewBuilder()
resolver.Register(r)

// 服务注册与注销
err := grpcds.Register("service_name", "v1", serveAddr)
grpcds.Deregister()

// 服务发现
conn, err := grpc.Dial(grpcds.Schema+":///"+"service_name", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
conn, err := grpc.Dial(grpcds.Schema+":///"+"service_name?version=v1", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
conn, err := grpc.Dial(grpcds.Schema+"://authority/"+"service_name?version=v1", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
```