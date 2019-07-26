package grpcds

import (
	"log"
	"net"
	"strings"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/addr"
)

var deregisterCh = make(chan struct{})

func init() {
	// TODO Registry配置及初始化
	registry.DefaultRegistry = consul.NewRegistry()
}

// Register
func Register(name, version, address string) error {
	var err error
	var host, port string

	if cnt := strings.Count(address, ":"); cnt >= 1 {
		// ipv6 address in format [host]:port or ipv4 host:port
		host, port, err = net.SplitHostPort(address)
		if err != nil {
			return err
		}
	} else {
		host = address
	}

	addr, err := addr.Extract(host)
	if err != nil {
		return err
	}

	// register service
	node := &registry.Node{
		Id:      name + "" + address,
		Address: net.JoinHostPort(addr, port),
	}
	service := &registry.Service{
		Name:    name,
		Version: version,
		Nodes:   []*registry.Node{node},
	}

	log.Printf("register service: %v", service)

	err = registry.Register(service)
	if err != nil {
		return err
	}

	// wait deregister then delete
	go func() {
		<-deregisterCh
		registry.Deregister(service)
		deregisterCh <- struct{}{}
	}()

	return nil
}

// UnRegister delete registered service from etcd
func Deregister() {
	deregisterCh <- struct{}{}
	<-deregisterCh
}
