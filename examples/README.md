### 打包
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' ./main.go
docker build -t hbchen/micro-mesh-example-api:v0.0.x .
docker build -t hbchen/micro-mesh-example-srv-1:v0.0.x .
docker build -t hbchen/micro-mesh-example-srv-2:v0.0.x .
```

### 验证

**Header说明**
- `"Authorization: Bearer $TOKEN"`使用[JWT](/deploy/k8s/gateway-jwt.yaml)需要
- `"x-custom-token: abc"`使用[自定义Auth adapter](/examples/adapter/auth)需要
- `"Grpc-Metadata-x-tier: 2"`为了方便测试，使用`x-tier`控制服务调用的层级，
    - =0或空:`api`
    - =1:`api`->`srv-1`
    - ≥2:`api`->`srv-1`->`srv-2`
    
###### curl
        
```bash
# JWT token
$ TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)

$ curl -H "Authorization: Bearer $TOKEN" -H "Grpc-Metadata-x-tier: 2" -X GET http://35.193.180.18/v1/example/call/Hobo
{"response":{"name":"Hobo"},"chain":["service_2","service_1","api example"]}                                          
 
$ curl -H "Authorization: Bearer $TOKEN" -H "Grpc-Metadata-x-tier: 2" -X GET http://35.193.180.18/v1/example/call/responsebody/Hobo
{"name":"Hobo"}

$ curl -H "Authorization: Bearer $TOKEN" -H "Grpc-Metadata-x-tier: 2" -X POST -d '{"name":"Hobo"}' http://35.193.180.18/v1/example/call
{"response":{"name":"Hobo"},"chain":["service_2","service_1","api example"]}

$ curl -H "Authorization: Bearer $TOKEN" -H "Grpc-Metadata-x-tier: 2" -X POST -d '{"name":"Hobo"}' http://35.193.180.18/v1/example/call/responsebody
{"name":"Hobo"}
```

###### postman
- Import`micro-mesh.postman_collection.json` `micro-mesh.postman_environment.json`
- `HOST`、`PORT`、`AUTH_TOKEN`、`X_CUSTOM_TOKEN`、`X_TIER`环境变量修改
    - AUTH_TOKEN:`curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s`

### Auth Adapter
- JWT
    - 服务端验证token有效性
    - 应对密码修改、终端数量限制等场景
- ACL
    - 服务端获取用户角色，做API访问控制
    - 用户角色及接口授权策略实时生效
