package client

import (
	"io"
	"time"

	"github.com/hb-go/grpc-contrib/client"
	"github.com/hb-go/grpc-contrib/registry"
	"google.golang.org/grpc"

	"github.com/hb-go/micro-mesh/pkg/util"
)

var (
	pool *client.Pool
)

func init() {
	// TODO client pool管理
	pool = client.NewPool(100, time.Second*30)
}

func Client(desc *grpc.ServiceDesc, options ...Option) (*grpc.ClientConn, io.Closer, error) {
	opts := newOptions(options...)
	if len(opts.Name) > 0 {
		desc.ServiceName = opts.Name
	}

	addr := registry.NewTarget(desc, opts.RegistryOptions...)

	conn, err := pool.Get(addr, opts.DialOptions...)
	if err != nil {
		return nil, nil, err
	}

	c := &util.FuncCloser{
		CloseFunc: func() error {
			pool.Put(addr, conn, err)
			return nil
		},
	}

	return conn.GetCC(), c, nil
}
