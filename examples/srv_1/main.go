package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/hb-go/micro-mesh/examples/pkg/conv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/hb-go/micro-mesh/proto"
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

	nmd := metautils.ExtractIncoming(ctx)
	if tier, err := strconv.Atoi(nmd.Get("x-tier")); nmd.Get("x-tier") == "" || (err == nil && tier > 1) {
		opts := []grpc.DialOption{grpc.WithInsecure()}
		addr := conv.ServiceTargetParse(pb.Services_mm_example_srv_2.String(), remoteAddr)
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
		ServiceName: pb.Services_mm_example_srv_1.String(),
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

	lis, err := net.Listen("tcp", serveAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := service{}
	pb.RegisterExampleServiceServer(s, &srv)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc failed to serve: %v", err)
	}

}
