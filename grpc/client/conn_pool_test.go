package client

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	pb "github.com/hb-go/micro-mesh/proto"
	"google.golang.org/grpc"
)

type testService struct{}

func (s *testService) Call(ctx context.Context, in *pb.ReqMessage) (*pb.RspMessage, error) {
	rsp := &pb.RspMessage{
		Response: &pb.RspMessage_Response{Name: in.Name},
	}

	return rsp, nil
}

func testPool(t *testing.T, size int, ttl time.Duration) {
	// setup server
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &testService{})

	go s.Serve(l)
	defer s.Stop()

	// zero pool
	p := NewPool(size, ttl)

	for i := 0; i < 10; i++ {
		// get a conn
		cc, err := p.Get(l.Addr().String(), grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}

		rsp, err := pb.NewExampleServiceClient(cc.GetCC()).Call(context.TODO(), &pb.ReqMessage{Name: "Hobo"})
		if err != nil {
			t.Fatal(err)
		}

		if rsp.Response.Name != "Hobo" {
			t.Fatalf("Got unexpected response %v", rsp.Response.Name)
		}

		// release the conn
		p.Put(l.Addr().String(), cc, nil)

		p.Lock()
		if i := len(p.conns[l.Addr().String()]); i > size {
			p.Unlock()
			t.Fatalf("pool size %d is greater than expected %d", i, size)
		}
		p.Unlock()
	}
}

func TestGRPCPool_0(t *testing.T) {
	testPool(t, 0, time.Minute)
}

func TestGRPCPool_5(t *testing.T) {
	testPool(t, 5, time.Minute)
}

func benchPool(t *testing.B, addr string, p *Pool) {

	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				// get a conn
				cc, err := p.Get(addr, grpc.WithInsecure())
				if err != nil {
					t.Fatal(err)
				}

				rsp, err := pb.NewExampleServiceClient(cc.GetCC()).Call(context.TODO(), &pb.ReqMessage{Name: "Hobo"})
				if err != nil {
					t.Fatal(err)
				}

				if rsp.Response.Name != "Hobo" {
					t.Fatalf("Got unexpected response %v", rsp.Response.Name)
				}

				// release the conn
				p.Put(addr, cc, nil)

				p.Lock()
				if i := len(p.conns[addr]); i > p.size {
					p.Unlock()
					t.Fatalf("pool size %d is greater than expected %d", i, p.size)
				}
				p.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

// Benchmark
// go test ./grpc/client -test.bench=".*"
func Benchmark_GRPCPool_0(b *testing.B) {
	b.StopTimer()

	// setup server
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		b.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &testService{})

	go s.Serve(l)
	defer s.Stop()

	// pool
	p := NewPool(0, time.Minute)

	b.StartTimer()

	for i := 0; i < b.N; i++ { // use b.N for looping
		benchPool(b, l.Addr().String(), p)
	}
}

func Benchmark_GRPCPool_5(b *testing.B) {
	b.StopTimer()

	// setup server
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		b.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &testService{})

	go s.Serve(l)
	defer s.Stop()

	// pool
	p := NewPool(5, time.Minute)

	b.StartTimer()

	for i := 0; i < b.N; i++ { // use b.N for looping
		benchPool(b, l.Addr().String(), p)
	}
}
