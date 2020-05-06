package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hb-go/micro-mesh/examples/common"
	"github.com/hb-go/micro-mesh/examples/service"
	pb "github.com/hb-go/micro-mesh/proto"
)

var (
	serveAddr      string
	serviceName    string
	serviceVersion string
	cmdHelp        bool
)

func init() {
	flag.StringVar(&serveAddr, "serve_addr", "", "serve address.")
	flag.StringVar(&serviceName, "service_name", "ExampleService", "service name.")
	flag.StringVar(&serviceVersion, "service_version", "latest", "service version.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func init() {
	// logger
	if err := common.Logger("example-srv"); err != nil {
		panic(err)
	}
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

	log.Infof("registry %v", registry.DefaultRegistry)

	// 服务注册
	example := pb.RegistryServiceExampleService
	if len(serviceName) > 0 {
		example.Name = service.ServicePrefix + serviceName
	}
	example.Version = serviceVersion
	example.Nodes = []*registry.Node{
		&registry.Node{
			Id:      example.Name + "-" + uuid.New().String(),
			Address: lis.Addr().String(),
		},
	}
	registry.Register(&example)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			common.ServerInterceptors()...,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	srv := service.Service{}
	pb.RegisterExampleServiceServer(s, &srv)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT)

	exitCh := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		select {
		case <-exitCh:
			log.Infof("server exit")
		case sig := <-ch:
			log.Infof("receive signal: %v", sig.String())
			s.GracefulStop()
		}

		// 注销服务
		registry.Deregister(&example)

		wg.Done()
	}(wg)

	log.Infof("grpc serve addr: %v", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc failed to serve: %v", err)
	} else {
		log.Infof("grpc serve end")
	}

	close(exitCh)
	wg.Wait()
	log.Infof("end")
}
