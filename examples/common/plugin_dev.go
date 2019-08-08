// +build dev

package common

import (
	"context"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	"github.com/hb-go/pkg/log"
	mregistry "github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"google.golang.org/grpc"
)

func init() {
	// mregistry.DefaultRegistry = consul.NewRegistry(consul.TCPCheck(time.Second * 5))
	mregistry.DefaultRegistry = consul.NewRegistry()
}

func clientInterceptors() []grpc.UnaryClientInterceptor {
	return nil
	interceptors := make([]grpc.UnaryClientInterceptor, 0, 1)

	// retry
	interceptors = append(interceptors, grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(time.Millisecond*100),
	))

	// tracing
	interceptors = append(interceptors, grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(nil)))

	interceptors = append(interceptors, func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Infof("client interceptor method: %v", method)
		// breaker hystrix
		return hystrix.Do(method, func() error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}, func(e error) error {
			log.Errorf("circuit breaker fallback with error: %v", e)
			return e
		})
	})

	return interceptors
}

func serverInterceptors() []grpc.UnaryServerInterceptor {
	interceptors := make([]grpc.UnaryServerInterceptor, 0, 1)

	// recovery
	interceptors = append(interceptors, grpc_recovery.UnaryServerInterceptor())

	// tracing
	interceptors = append(interceptors, grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(nil)))

	return interceptors
}