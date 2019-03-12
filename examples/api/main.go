package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-go/micro-mesh/examples/pkg/conv"
	pb "github.com/hb-go/micro-mesh/proto"
	"google.golang.org/grpc"
)

var (
	serveAddr string
	remoteAddr string
	cmdHelp  bool
)

func init() {
	flag.StringVar(&serveAddr, "serve_addr", ":9080", "serve address.")
	flag.StringVar(&remoteAddr, "remote_addr", ":9080", "remote address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

type service struct{}

func (s *service) Call(ctx context.Context, in *pb.ReqMessage) (*pb.RspMessage, error) {
	log.Printf("received: %v", in.Name)

	rsp := &pb.RspMessage{
		Response: &pb.RspMessage_Response{Name: in.Name},
	}

	// 注意这里是ExtractOutgoing()
	nmd := metautils.ExtractOutgoing(ctx)
	if tier, err := strconv.Atoi(nmd.Get("x-tier")); nmd.Get("x-tier") == "" || (err == nil && tier > 0) {
		// TODO cc pool
		addr := conv.ServiceTargetParse(pb.Services_mm_example_srv_1.String(), remoteAddr)
		opts := []grpc.DialOption{grpc.WithInsecure()}
		cc, err := grpc.Dial(addr, opts...)
		if err != nil {
			return nil, err
		}

		rsp, err = pb.NewExampleServiceClient(cc).Call(ctx, in)
		if err != nil {
			return nil, err
		}
	}

	chain := &pb.RspMessage_Chain{
		ServiceName: pb.Services_mm_example_api.String(),
	}
	if incoming, err := json.Marshal(nmd); err == nil {
		chain.Ctx = string(incoming)
	}

	rsp.Chain = append(rsp.Chain, chain)

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

	if err := http.ListenAndServe(serveAddr, mux); err != nil {
		log.Fatalf("http failed to serve: %v", err)
	}
}
