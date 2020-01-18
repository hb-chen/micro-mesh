package client

import (
	"github.com/hb-go/grpc-contrib/registry"
	"google.golang.org/grpc"
)

type Option func(options *Options)

type Options struct {
	Name            string
	RegistryOptions []registry.Option
	DialOptions     []grpc.DialOption
}

// 默认gRPC DialOption
var DefaultDialOpts = []grpc.DialOption{
	grpc.WithInsecure(),
	grpc.WithBalancerName("round_robin"),
	grpc.WithBlock(),
}

func newOptions(options ...Option) Options {
	opts := Options{
		RegistryOptions: make([]registry.Option, 0),
		DialOptions:     DefaultDialOpts,
	}

	for _, o := range options {
		o(&opts)
	}

	return opts
}

// 指定gRPC服务名称，替换ServiceDesc.ServiceName
func WithName(name string) Option {
	return func(options *Options) {
		options.Name = name
	}
}

// 注册中心选项
func WithRegistryOptions(option ...registry.Option) Option {
	return func(options *Options) {
		options.RegistryOptions = append(options.RegistryOptions, option...)
	}
}

// gRPC DialOption
func WithDialOptions(option ...grpc.DialOption) Option {
	return func(options *Options) {
		options.DialOptions = append(options.DialOptions, option...)
	}
}
