# Istio部署

```bash
curl -v -HHost:hbchen.com  -X GET 'http://192.168.39.147:31380/v1/example/call/hobo?services=\[\{%22name%22:%22ExampleService1%22,%22version%22:%22v1%22,%22services%22:\[\{%22name%22:%22ExampleService2%22,%22version%22:%22v1%22,%22services%22:\[\]\}\]\}\]'
curl -v -HHost:hbchen.com  -X GET 'http://192.168.39.147:31380/v1/example/call/hobo'

TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.1/security/tools/jwt/samples/demo.jwt -s)
curl -v -HHost:hbchen.com -H "Authorization: Bearer $TOKEN" -X GET 'http://192.168.39.147:31380/v1/example/call/hobo'
```