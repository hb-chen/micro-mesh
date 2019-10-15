module github.com/hb-go/micro-mesh

exclude (
	istio.io/api v0.0.0-20190726010239-927332251e18
	istio.io/istio v0.0.0-20190726050400-510058f64f7b
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/hb-go/grpc-contrib v0.0.0-20190809115636-90e6e3609fcd
	github.com/hb-go/pkg v0.0.2-0.20190805134718-346b31e462e2
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-plugins v1.2.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	go.uber.org/zap v1.10.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.22.1
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
)
