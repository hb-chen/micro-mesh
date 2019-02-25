### 打包
```
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' ./main.go
$ docker build -t hbchen/micro-mesh-example-api:v0.0.x .
$ docker build -t hbchen/micro-mesh-example-srv-1:v0.0.x .
$ docker build -t hbchen/micro-mesh-example-srv-2:v0.0.x .
```

### 验证
###### curl
```bash
# JWT token
$ TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)

$ curl -H "Authorization: Bearer $TOKEN" -X GET http://35.193.180.18/v1/example/call/Hobo
{"response":{"name":"Hobo"},"chain":["service_2","service_1","api example"]}                                          
 
$ curl -H "Authorization: Bearer $TOKEN" -X GET http://35.193.180.18/v1/example/call/responsebody/Hobo
{"name":"Hobo"}

$ curl -H "Authorization: Bearer $TOKEN" -X POST -d '{"name":"Hobo"}' http://35.193.180.18/v1/example/call
{"response":{"name":"Hobo"},"chain":["service_2","service_1","api example"]}

$ curl -H "Authorization: Bearer $TOKEN" -X POST -d '{"name":"Hobo"}' http://35.193.180.18/v1/example/call/responsebody
{"name":"Hobo"}
```

###### postman
- Import`micro-mesh.postman_collection.json` `micro-mesh.postman_environment.json`
- `HOST`、`AUTH_TOKEN`环境变量修改
    - `curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s`
