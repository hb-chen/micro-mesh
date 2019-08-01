# gRPC微服务示例

## 服务发现
默认使用consul做服务发现

## 服务

### 运行
```bash
# api
$ go run main.go -serve_addr :9080 -services '[{"name":"ExampleService","version":"latest","services":[]}]'
 
# srv
go run main.go -service_name ExampleService service_version latest
```

### `srv`服务调用
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
$ curl -X GET 'http://localhost:9080/v1/example/call/hobo?services=\[\{%22name%22:%22ExampleService%22,%22version%22:%22latest%22,%22services%22:\[\]\}\]'
$ curl -X GET 'http://localhost:9080/v1/example/call/responsebody/hobo?services=\[\{%22name%22:%22ExampleService%22,%22version%22:%22latest%22,%22services%22:\[\]\}\]'

$ curl -X POST -d '{"name":"hobo","services":[{"name":"ExampleService","version":"latest","services":[]},{"name":"ExampleService","version":"latest","services":[]}]}' 'http://localhost:9080/v1/example/call'
$ curl -X POST -d '{"name":"hobo","services":[{"name":"ExampleService","version":"latest","services":[]},{"name":"ExampleService","version":"latest","services":[]}]}' 'http://localhost:9080/v1/example/call/responsebody'
```

### postman
- Import`micro-mesh.postman_collection.json` `micro-mesh.postman_environment.json`
- `HOST`、`PORT`、`AUTH_TOKEN`、`X_CUSTOM_TOKEN`环境变量修改
    - AUTH_TOKEN:`curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s`

## Istio

### 验证

**Header说明**
- `"Authorization: Bearer $TOKEN"`使用[JWT](/deploy/k8s/jwt/gateway-jwt.yaml)需要
- `"x-custom-token: abc"`使用[自定义Auth adapter](/examples/adapter/auth)需要

```bash
# JWT token
$ TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)

$ curl -H "Authorization: Bearer $TOKEN" ……
```