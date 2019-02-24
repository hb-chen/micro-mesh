package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"

	pb "github.com/hb-go/micro-mesh/proto"
)

var (
	grpcAddr string
	httpAddr string
	cmdHelp  bool
)

func init() {
	flag.StringVar(&grpcAddr, "grpc_addr", ":9080", "grpc server address.")
	flag.StringVar(&httpAddr, "http_addr", ":9080", "http server address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

type service struct{}

func (s *service) Call(ctx context.Context, in *pb.ReqMessage) (*pb.RspMessage, error) {
	log.Printf("received: %v", in.Name)

	// TODO cc pool
	addr := pb.Services_mm_example_srv_1.String() + grpcAddr
	opts := []grpc.DialOption{grpc.WithInsecure()}
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	rsp, err := pb.NewExampleServiceClient(cc).Call(ctx, in, )
	rsp.Chain = append(rsp.Chain, "api example")

	return rsp, nil
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	s := grpc.NewServer()
	srv := service{}
	pb.RegisterExampleServiceServer(s, &srv)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := pb.RegisterExampleServiceHandlerServer(ctx, mux, &srv)
	if err != nil {
		log.Panic(err)
	}

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("http failed to serve: %v", err)
	}
}
