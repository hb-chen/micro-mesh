// +build istio

package common

import (
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hb-go/grpc-contrib/metadata"
	_ "github.com/hb-go/grpc-contrib/registry/istio"
	"google.golang.org/grpc"
)

func init() {
}

func clientInterceptors() []grpc.UnaryClientInterceptor {
	interceptors := make([]grpc.UnaryClientInterceptor, 0, 1)

	// retry
	interceptors = append(interceptors, grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(time.Millisecond*100),
	))

	// metadata
	interceptors = append(interceptors, metadata.UnaryClientInterceptor(
		metadata.WithHeader("x-b3-traceid"),
		metadata.WithHeader("x-b3-spanid"),
		metadata.WithHeader("x-b3-parentspanid"),
		metadata.WithHeader("x-b3-sampled"),
		metadata.WithHeader("x-b3-flags"),
	))

	return interceptors
}

func serverInterceptors() []grpc.UnaryServerInterceptor {
	interceptors := make([]grpc.UnaryServerInterceptor, 0, 1)

	// recovery
	interceptors = append(interceptors, grpc_recovery.UnaryServerInterceptor())

	return interceptors
}
