// +build k8s

package common

import (
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	mregistry "github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/kubernetes"
	"google.golang.org/grpc"
)

func init() {
	mregistry.DefaultRegistry = kubernetes.NewRegistry()
}

func gatewayMetadataOptions() []metadata.Option {
	opts := []metadata.Option{}

	return opts
}

func clientInterceptors() []grpc.UnaryClientInterceptor {
	interceptors := make([]grpc.UnaryClientInterceptor, 0, 1)

	// retry
	interceptors = append(interceptors, grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(time.Millisecond*100),
	))

	return interceptors
}

func serverInterceptors() []grpc.UnaryServerInterceptor {
	interceptors := make([]grpc.UnaryServerInterceptor, 0, 1)

	// recovery
	interceptors = append(interceptors, grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
		log.Errorf("grpc_recovery: %+v", p)
		return status.Errorf(codes.Internal, "%s", p)
	})))

	return interceptors
}
