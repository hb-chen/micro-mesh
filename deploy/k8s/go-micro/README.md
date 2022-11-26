# go-micro/k8s注册中心

```bash
kubectl label namespace default istio-injection=disable --overwrite
```

```shell
# Service
curl -v -X GET 'http://172.25.183.53:30001/v1/example/call/hobo'
curl -v -X GET 'http://172.25.183.53:30001/v1/example/call/responsebody/hobo'

hey -z 60s -c 2 -q 10 -host micro.k8s.hbchen.com 'http://172.25.183.53:30127/v1/example/call/hobo'

# Ingress
curl -v -HHost:micro.k8s.hbchen.com -X GET 'http://172.25.183.53:30127/v1/example/call/hobo'
curl -v -HHost:micro.k8s.hbchen.com -X GET 'http://172.25.183.53:30127/v1/example/call/responsebody/hobo'
```

```bash
curl -v -X GET 'http://192.168.39.147:30001/v1/example/call/hobo'
curl -v -X GET 'http://192.168.39.147:30001/v1/example/call/responsebody/hobo'
# query参数指定调用链路
curl -v -X GET 'http://192.168.39.147:30001/v1/example/call/hobo?services=\[\{%22name%22:%22ExampleService1%22,%22version%22:%22latest%22,%22services%22:\[\{%22name%22:%22ExampleService2%22,%22version%22:%22latest%22,%22services%22:\[\]\}\]\}\]'

curl -v -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://192.168.39.147:30001/v1/example/call'
curl -v -X POST -d '{"name":"hobo","services":[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]}' 'http://192.168.39.147:30001/v1/example/call/responsebody'
```