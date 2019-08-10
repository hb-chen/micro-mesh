package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-go/grpc-contrib/metadata"
	"github.com/hb-go/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/hb-go/micro-mesh/examples/common"
	"github.com/hb-go/micro-mesh/examples/service"
	pb "github.com/hb-go/micro-mesh/proto"
)

var (
	services  string
	serveAddr string
	cmdHelp   bool
)

func init() {
	flag.StringVar(&serveAddr, "serve_addr", ":9080", "serve address.")
	flag.StringVar(&services, "services", `[{"name":"ExampleService1","version":"latest","services":[{"name":"ExampleService2","version":"latest","services":[]}]}]`, "remote address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func init() {
	// logger
	if err := common.Logger("example-api"); err != nil {
		panic(err)
	}
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			common.ServerInterceptors()...,
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

	mux := runtime.NewServeMux(
		// istio trace header
		runtime.WithMetadata(metadata.GatewayMetadataAnnotator(
			common.GatewayMetadataOptions()...,
		)),
	)
	err := pb.RegisterExampleServiceHandlerFromEndpoint(
		ctx,
		mux,
		"",
		[]grpc.DialOption{
			grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return bcLis.Dial()
			}),
			grpc.WithDefaultCallOptions(),
			grpc.WithChainUnaryInterceptor(
				common.ClientInterceptors()...,
			),
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
