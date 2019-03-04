自定义[Adapter](https://preliminary.istio.io/zh/docs/concepts/policies-and-telemetry/#%E9%85%8D%E7%BD%AE%E6%A8%A1%E5%9E%8B)
===

### Istio源码
```bash
mkdir -p $GOPATH/src/istio.io/
cd $GOPATH/src/istio.io/
git clone https://github.com/istio/istio
```

### 开发环境
```bash
# build mixer server client 
cd istio
make mixs
make mixc

# copy auth adapter example
cp {micro-mesh path}/examples/adapter/auth mixer/adapter/auth

cd mixer/adapter/auth
go generate ./...
go build ./...
```

### 本地测试
```bash
go run cmd/main.go 44225

$GOPATH/out/darwin_amd64/release/mixs server \
--configStoreURL=fs://$GOPATH/src/istio.io/istio/mixer/adapter/auth/testdata \
--log_output_level=attributes:debug

$GOPATH/out/darwin_amd64/release/mixc check -s destination.service="mm-example-api.default.svc.cluster.local" --string_attributes "request.host=localhost" --stringmap_attributes "request.headers=x-custom-token:abc"
$GOPATH/out/darwin_amd64/release/mixc check -s destination.service="mm-example-api.default.svc.cluster.local" --string_attributes "request.host=localhost" --stringmap_attributes "request.headers=x-custom-token:efg"

```

### 打包镜像
```bash
cd mixer/adapter/auth

CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -v -o bin/auth ./cmd/
    
docker build -t hbchen/micro-mesh-example-adapter-auth:v0.0.1 .
docker push hbchen/micro-mesh-example-adapter-auth:v0.0.1
```

### Istio环境部署
```bash
# attributes.yaml -> istio/mixer/testdata/config/attributes.yaml 
# template.yaml -> istio/mixer/template/authorization/template.yaml
kubectl apply -f examples/adapter/auth/testdata/attributes.yaml -f examples/adapter/auth/testdata/template.yaml

kubectl apply -f examples/adapter/auth/testdata/auth-adapter.yaml
```

```bash
kubectl apply -f examples/adapter/auth/operatorconfig/cluser-service.yaml
kubectl apply -f examples/adapter/auth/operatorconfig/operator-cfg.yaml
```