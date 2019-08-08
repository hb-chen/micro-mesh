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

func clientInterceptors() []grpc.UnaryClientInterceptor {
	return nil
}

func serverInterceptors() []grpc.UnaryServerInterceptor {
	return nil
}
