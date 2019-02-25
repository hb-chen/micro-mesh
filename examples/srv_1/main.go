package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/hb-go/micro-mesh/proto"
)

var (
	grpcAddr string
	cmdHelp  bool
)

func init() {
	flag.StringVar(&grpcAddr, "grpc_addr", ":9080", "grpc server address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

type service struct{}

func (s *service) Call(ctx context.Context, in *pb.ReqMessage) (*pb.RspMessage, error) {
	log.Printf("received: %v", in.Name)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	addr := strings.Replace(pb.Services_mm_example_srv_2.String(), "_", "-", -1) + grpcAddr
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	rsp, err := pb.NewExampleServiceClient(cc).Call(ctx, in, )
	rsp.Chain = append(rsp.Chain, "service_1")

	return rsp, nil
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	lis, err := net.Listen("tcp", grpcAddr)
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
