// +build dev

package common

import (
	"context"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/hb-go/grpc-contrib/metadata"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	"github.com/hb-go/pkg/log"
	mregistry "github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	metadataOptions = []metadata.Option{
		metadata.WithHeader("x-request-id"),
		metadata.WithHeader("uber-trace-id"),
	}
)

type validator interface {
	Validate() error
}

func init() {
	mregistry.DefaultRegistry = etcd.NewRegistry()
}

func gatewayMetadataOptions() []metadata.Option {
	return metadataOptions
}

func clientInterceptors() []grpc.UnaryClientInterceptor {
	interceptors := make([]grpc.UnaryClientInterceptor, 0, 1)

	// retry
	interceptors = append(interceptors, grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(time.Millisecond*100),
	))

	// metadata
	interceptors = append(interceptors, metadata.UnaryClientInterceptor(metadataOptions...))

	// tracing
	interceptors = append(interceptors, grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())))

	// Options
	// Request参数自动验证
	interceptors = append(interceptors, func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return status.Errorf(codes.InvalidArgument, err.Error())
			}
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	})

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
	interceptors = append(interceptors, grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())))

	// 参数自动验证
	interceptors = append(interceptors, grpc_validator.UnaryServerInterceptor())

	return interceptors
}
