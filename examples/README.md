# gRPC微服务示例

## 服务发现
默认使用Etcd做服务发现

## Local测试

> 需要[Etcd](https://etcd.io/)服务

### 运行
```bash
# api
go run -tags "dev" main.go -serve_addr :9080 -services '[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]'
 
# srv
go run -tags "dev" main.go -service_name ExampleService1 -service_version latest
go run -tags "dev" main.go -service_name ExampleService2 -service_version latest

```

#### `srv`服务调用说明
用`JSON`结构的`services`控制多层服务调用，指定`name`、`version`以及`services`子服务。

```json
[
  {
    "name" : "ExampleService-1",
    "version" : "latest",
    "services" : [
      {
        "name" : "ExampleService-3",
        "version" : "latest",
        "services" : [

        ]
      },
      {
        "name" : "ExampleService-4",
        "version" : "latest",
        "services" : [

        ]
      }
    ]
  },
  {
    "name" : "ExampleService-2",
    "version" : "latest",
    "services" : [

    ]
  }
]
```

### curl
        
```bash
curl -v -X GET 'http://localhost:9080/v1/example/call/hobo'
curl -v -X GET 'http://localhost:9080/v1/example/call/responsebody/hobo'
# query参数指定调用链路
curl -v -X GET 'http://localhost:9080/v1/example/call/hobo?services=\[\{%22name%22:%22ExampleService1%22,%22version%22:%22latest%22,%22services%22:\[\{%22name%22:%22ExampleService2%22,%22version%22:%22latest%22,%22services%22:\[\]\}\]\}\]'

curl -v -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://localhost:9080/v1/example/call'
curl -v -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://localhost:9080/v1/example/call/responsebody'
```

### postman
- Import`micro-mesh.postman_collection.json` `micro-mesh.postman_environment.json`
- `HOST`、`PORT`、`AUTH_TOKEN`、`X_CUSTOM_TOKEN`环境变量修改
    - AUTH_TOKEN:`curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s`

## 部署

### 打包
```bash
# k8s
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -tags "k8s" ./main.go

docker build -t hbchen/micro-mesh-example-api:v0.0.4_k8s .
docker build -t hbchen/micro-mesh-example-srv:v0.0.4_k8s .

# istio
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -tags "istio" ./main.go

docker build -t hbchen/micro-mesh-example-api:v0.0.4_istio .
docker build -t hbchen/micro-mesh-example-srv:v0.0.4_istio .
```
### go-micro/k8s
[deploy/k8s/go-micro](/deploy/k8s/go-micro)

### Istio
[deploy/k8s/istio](/deploy/k8s/istio)

## Header说明
- `"Authorization: Bearer $TOKEN"`使用[JWT](/deploy/k8s/istio/jwt/gateway-jwt.yaml)需要
- `"x-custom-token: abc"`使用[自定义Auth adapter](/examples/adapter/auth)需要

```bash
# JWT token
$ TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)

$ curl -H "Authorization: Bearer $TOKEN" ……
```

