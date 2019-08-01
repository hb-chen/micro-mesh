module github.com/hb-go/micro-mesh

exclude (
	istio.io/api v0.0.0-20190726010239-927332251e18
	istio.io/istio v0.0.0-20190726050400-510058f64f7b
)

require (
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/grpc-gateway v1.9.2
	github.com/hb-go/grpc-contrib v0.0.0-20190731054334-689f0a531cef
	github.com/hb-go/pkg v0.0.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.22.1
)
