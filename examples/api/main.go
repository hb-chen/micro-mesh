package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-go/grpc-contrib/metadata"
	"github.com/hb-go/micro-mesh/examples/common"
	"github.com/hb-go/micro-mesh/examples/service"
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
	flag.StringVar(&services, "services", `[{"name":"ExampleService","version":"latest","services":[{"name":"ExampleService","version":"latest","services":[]}]}]`, "service config.")
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

	srv := service.Service{
		Services: services,
	}

	mux := runtime.NewServeMux(
		// istio trace header
		runtime.WithMetadata(metadata.GatewayMetadataAnnotator(
			common.GatewayMetadataOptions()...,
		)),
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc-gateway进程内调用
	// grpc-gateway version ≥ v1.11.1
	err := pb.RegisterExampleServiceHandlerServer(
		ctx,
		mux,
		&srv,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("http serve addr: %v", serveAddr)
	if err := http.ListenAndServe(serveAddr, mux); err != nil {
		log.Fatalf("http failed to serve: %v", err)
	}
}
