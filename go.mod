module github.com/hb-go/micro-mesh

go 1.13

exclude (
	istio.io/api v0.0.0-20190726010239-927332251e18
	istio.io/istio v0.0.0-20190726050400-510058f64f7b
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/hb-go/grpc-contrib v0.0.0-20200504121232-acfdf6a4e1e0
	github.com/hb-go/pkg v0.0.2-0.20190805134718-346b31e462e2
	github.com/opentracing/opentracing-go v1.1.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.26.0
)
