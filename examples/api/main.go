package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	"github.com/hb-go/micro-mesh/examples/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/hb-go/micro-mesh/proto"
	"github.com/hb-go/pkg/log"
)

var (
	services  string
	serveAddr string
	cmdHelp   bool
)

func init() {
	flag.StringVar(&serveAddr, "serve_addr", ":9080", "serve address.")
	flag.StringVar(&services, "services", `[{"name":"ExampleService","version":"latest","services":[]}]`, "remote address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	srv := service.Service{
		Services: services,
	}
	pb.RegisterExampleServiceServer(s, &srv)

	bcLis := bufconn.Listen(1024 * 1024)
	go s.Serve(bcLis)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := pb.RegisterExampleServiceHandlerFromEndpoint(
		ctx,
		mux,
		"",
		[]grpc.DialOption{
			grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return bcLis.Dial()
			}),
			grpc.WithInsecure(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("http serve addr: %v", serveAddr)
	if err := http.ListenAndServe(serveAddr, mux); err != nil {
		log.Fatalf("http failed to serve: %v", err)
	}
}
