package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	log.Infof("registry %v",registry.DefaultRegistry)

	// 服务注册
	if len(serviceName) > 0 {
		sd := pb.ServiceDescExampleService()
		sd.ServiceName = service.ServicePrefix + serviceName
		if err := registry.Register(&sd, registry.WithVersion(serviceVersion), registry.WithAddr(lis.Addr().String())); err != nil {
			log.Fatal(err)
		}
	} else if err := pb.RegisterExampleService(registry.WithVersion(serviceVersion), registry.WithAddr(lis.Addr().String())); err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			common.ServerInterceptors()...
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
		registry.Deregister(nil)
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
