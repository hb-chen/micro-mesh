package service

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/hb-go/grpc-contrib/client"
	"github.com/hb-go/grpc-contrib/registry"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	pb "github.com/hb-go/micro-mesh/proto"
	"github.com/hb-go/pkg/dispatcher"
	"github.com/hb-go/pkg/log"
	gopool "github.com/hb-go/pkg/pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var (
	pool *client.Pool
)

const ServicePrefix = "com.hbchen."

func init() {
	pool = client.NewPool(100, time.Second*30)
}

type Service struct {
	Services string
}

func rangeService(services []*pb.Service) {
	for _, s := range services {
		log.Debugf("service: %v, sub services count:%d", s.Name, len(s.Services))
		if len(s.Services) > 0 {
			rangeService(s.Services)
		}
	}
}

func (s *Service) ApiCall(ctx context.Context, in *pb.ApiRequest) (*pb.Response, error) {
	log.Debugf("ApiCall() received: %v", in.Name)

	servicesJson := s.Services
	if len(in.Services) > 0 {
		servicesJson = in.Services
	}

	services := make([]*pb.Service, 0, 1)
	if err := json.Unmarshal([]byte(servicesJson), &services); err != nil {
		log.Infof("services parse error: %v", err)
	} else {
		rangeService(services)
	}

	req := &pb.Request{
		Name:     in.Name,
		Services: services,
	}
	return s.handler(ctx, req)
}

func (s *Service) Call(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Debugf("Call() received: %v", in.Name)
	return s.handler(ctx, in)
}

func (s *Service) handler(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	mu := sync.Mutex{}
	rsp := &pb.Response{
		Msg:   "Hello " + in.Name,
		Chain: make([]*pb.Response_Chain, 0, 1),
	}

	chain := &pb.Response_Chain{
		Service: &pb.Service{
			Name:     in.Name,
			Version:  in.Version,
			Services: in.Services,
		},
		Chain: make([]*pb.Response_Chain, 0, len(in.Services)),
	}
	nmd := metautils.ExtractIncoming(ctx)
	if incoming, err := json.Marshal(nmd); err == nil {
		chain.Ctx = string(incoming)
	}

	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock()}

	gp := gopool.NewGoroutinePool(len(in.Services), false)
	gp.AddWorkers(2) // worker num=3
	dp := dispatcher.NewDispatcher(gp)

	wg := sync.WaitGroup{}
	h := func(req *pb.Request) error {
		defer wg.Done()

		sd := pb.ServiceDescExampleService()
		sd.ServiceName = ServicePrefix + req.Name
		addr := registry.NewTarget(sd, registry.WithVersion(req.Version))
		log.Debugf("addr: %v", addr)

		cc1, err := pool.Get(addr, opts...)
		if err != nil {
			return err
		}

		rsp1, err := pb.NewExampleServiceClient(cc1.GetCC()).Call(ctx, req)
		log.Debugf("service: %v dispatch done, sub services: %v", req.Name, req.Services)

		mu.Lock()
		chain.Chain = append(chain.Chain, rsp1.Chain...)
		mu.Unlock()

		pool.Put(addr, cc1, err)
		return nil
	}

	handlers := make([]dispatcher.DispatchHandler, 0, len(in.Services))
	for _, srv := range in.Services {
		req := &pb.Request{
			Name:     srv.Name,
			Version:  srv.Version,
			Services: srv.Services,
		}
		handlers = append(handlers, func(i interface{}) error {
			return h(req)
		})
	}
	wg.Add(len(handlers))
	dp.Dispatch(handlers...)
	wg.Wait()

	rsp.Chain = append(rsp.Chain, chain)

	return rsp, nil
}
