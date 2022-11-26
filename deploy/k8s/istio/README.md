# Istio部署

```bash
kubectl label namespace default istio-injection=enabled --overwrite
```

```shell
hey -z 60s -c 2 -q 10 -host istio.k8s.hbchen.com 'http://172.25.183.53:30559/v1/example/call/hobo'

curl -v -HHost:istio.k8s.hbchen.com -X GET 'http://172.25.183.53:30559/v1/example/call/hobo'
curl -v -HHost:istio.k8s.hbchen.com -X GET 'http://172.25.183.53:30559/v1/example/call/responsebody/hobo'
```

```bash
curl -v -HHost:istio.k8s.hbchen.com -X GET 'http://192.168.39.147:31380/v1/example/call/hobo'
curl -v -HHost:istio.k8s.hbchen.com -X GET 'http://192.168.39.147:31380/v1/example/call/responsebody/hobo'
# query参数指定调用链路
curl -v -HHost:istio.k8s.hbchen.com -X GET 'http://192.168.39.147:30001/v1/example/call/hobo?services=\[\{%22name%22:%22ExampleService1%22,%22version%22:%22latest%22,%22services%22:\[\{%22name%22:%22ExampleService2%22,%22version%22:%22latest%22,%22services%22:\[\]\}\]\}\]'

curl -v -HHost:istio.k8s.hbchen.com -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://192.168.39.147:31380/v1/example/call'
curl -v -HHost:hbchen.com -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://192.168.39.147:31380/v1/example/call/responsebody'
```

## JWT
```bash
TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)
curl -v -HHost:hbchen.com -H "Authorization: Bearer $TOKEN" -X GET 'http://192.168.39.147:31380/v1/example/call/hobo'
```