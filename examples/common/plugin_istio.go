// +build istio

package common

import (
	_ "github.com/hb-go/grpc-contrib/registry/istio"
	"google.golang.org/grpc"
)

func init() {
}

func clientInterceptors() []grpc.UnaryClientInterceptor {
	return nil
}

func serverInterceptors() []grpc.UnaryServerInterceptor {
	return nil
}
