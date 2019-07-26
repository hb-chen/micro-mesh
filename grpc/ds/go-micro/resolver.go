package etcdv3

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/micro/go-micro/registry"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/resolver"

	"github.com/hb-go/micro-mesh/pkg/log"
)

const Schema = "go-micro"
const watchLimit = 1.0
const watchBurst = 3

// implementation of grpc.resolve.Builder
type microBuilder struct {
	registry registry.Registry

	mu        sync.RWMutex
	resolvers map[string]*service
}

type service struct {
	name   string
	target resolver.Target

	builder *microBuilder

	mu       sync.RWMutex
	watching bool
	nodes    map[string][]resolver.Address // TODO 多版本

	conns     sync.Map
	connIndex int64
}

type microResolver struct {
	service *service
	version string

	index int64
	cc    resolver.ClientConn
}

// NewBuilder return resolver builder
func NewBuilder() resolver.Builder {
	return &microBuilder{
		resolvers: make(map[string]*service),
	}
}

// Scheme
func (b *microBuilder) Scheme() string {
	return Schema
}

// Build to resolver.Resolver
// target:schema://[authority]/{serviceName}[?version=v1]
// target使用query参数做version筛选
func (b *microBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	log.Infof("build with target: %v", target)
	var serviceName, serviceVersion string
	if c := strings.Count(target.Endpoint, "?"); c > 0 {
		if u, err := url.Parse(target.Endpoint); err == nil {
			serviceName = u.Path
			query := u.Query()
			if v := query.Get("version"); len(v) > 0 {
				serviceVersion = query.Get("version")
			}
		}
	} else {
		serviceName = target.Endpoint
	}

	log.Infof("service name: %v, version: %v", serviceName, serviceVersion)

	b.mu.Lock()
	s, ok := b.resolvers[serviceName]
	if ok {
		b.mu.Unlock()

		// 使用当前service nodes
		s.mu.Lock()
		var ccNodes []resolver.Address
		if serviceVersion == "" {
			ccNodes = make([]resolver.Address, 0, len(s.nodes))
			for _, v := range s.nodes {
				ccNodes = append(ccNodes, v...)
			}
		} else if nodes, ok := s.nodes[serviceVersion]; ok {
			ccNodes = nodes
		}

		// TODO 检查watching状态?
		if !s.watching {
			err := s.watch()
			if err != nil {
				s.mu.Unlock()
				return nil, err
			}
			s.watching = true
		}
		s.mu.Unlock()

		cc.UpdateState(resolver.State{Addresses: ccNodes})
	} else {
		s = &service{
			name:    serviceName,
			target:  target,
			builder: b,
			nodes:   make(map[string][]resolver.Address),
		}
		b.resolvers[s.name] = s

		s.mu.Lock()
		b.mu.Unlock()

		// 从registry获取services
		services, err := registry.GetService(s.name)
		if err != nil {
			s.mu.Unlock()
			return nil, err
		}

		count := 0
		for _, svc := range services {
			nodes := make([]resolver.Address, 0, len(svc.Nodes))
			for _, n := range svc.Nodes {
				addr := resolver.Address{
					Addr: n.Address,
				}
				nodes = append(nodes, addr)
			}
			s.nodes[svc.Version] = nodes
			count++
		}

		var ccNodes []resolver.Address
		if serviceVersion == "" {
			ccNodes = make([]resolver.Address, 0, count)
			for _, v := range s.nodes {
				ccNodes = append(ccNodes, v...)
			}
		} else if nodes, ok := s.nodes[serviceVersion]; ok {
			ccNodes = nodes
		}

		err = s.watch()
		if err != nil {
			s.mu.Unlock()
			return nil, err
		}

		s.watching = true
		s.mu.Unlock()

		cc.UpdateState(resolver.State{Addresses: ccNodes})
	}

	index := atomic.AddInt64(&s.connIndex, 1)
	r := &microResolver{
		service: s,
		version: serviceVersion,
		cc:      cc,
		index:   index,
	}

	s.conns.Store(index, r)
	return r, nil
}

// ResolveNow
func (r *microResolver) ResolveNow(rn resolver.ResolveNowOption) {
}

// Close
func (r *microResolver) Close() {
	r.service.conns.Delete(r.index)
}

func (s *service) watch() error {
	watcher, err := registry.Watch(registry.WatchService(s.name))
	if err != nil {
		return err
	}

	go func(watcher registry.Watcher) {
		limiter := rate.NewLimiter(rate.Limit(watchLimit), watchBurst)
		for {
			limiter.Wait(context.Background())
			if result, err := watcher.Next(); err == nil {
				if err := s.update(result); err == nil {
					var allNodes []resolver.Address
					s.conns.Range(func(key, value interface{}) bool {

						if r, ok := value.(*microResolver); ok {
							if r.version == "" {
								if allNodes == nil {
									allNodes = make([]resolver.Address, 0, len(s.nodes))
									for _, v := range s.nodes {
										allNodes = append(allNodes, v...)
									}
								}

								r.cc.UpdateState(resolver.State{Addresses: allNodes})
							} else if nodes, ok := s.nodes[r.version]; ok {

								r.cc.UpdateState(resolver.State{Addresses: nodes})
							}
						} else {
							log.Warnf("grpc.ds.go-micro: microResolver conv error")
						}

						return true
					})
				} else {
					log.Warnf("grpc.ds.go-micro: %v", err)
				}
			} else {
				log.Warnf("grpc.ds.go-micro: resolver watch error: %v", err)
				if err == registry.ErrWatcherStopped {
					return
				}
			}
		}
	}(watcher)

	return nil
}

func (s *service) update(res *registry.Result) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch res.Action {
	case "create", "update":
		nodes := make([]resolver.Address, 0, len(res.Service.Nodes))
		for _, n := range res.Service.Nodes {
			node := resolver.Address{
				Addr: n.Address,
			}
			nodes = append(nodes, node)
		}

		// append old nodes to new service
		if curNodes, ok := s.nodes[res.Service.Version]; ok {
			for _, node := range curNodes {
				var seen bool
				for _, n := range nodes {
					if node.Addr == n.Addr {
						seen = true
						break
					}
				}

				if !seen {
					nodes = append(nodes, node)
				}
			}
		}

		s.nodes[res.Service.Version] = nodes
		return nil

	case "delete":
		if curNodes, ok := s.nodes[res.Service.Version]; !ok {
			return nil
		} else {
			var nodes []resolver.Address

			// filter cur nodes to remove the dead one
			for _, cur := range curNodes {
				var seen bool
				for _, del := range res.Service.Nodes {
					if del.Address == cur.Addr {
						seen = true
						break
					}
				}
				if !seen {
					nodes = append(nodes, cur)
				}
			}

			s.nodes[res.Service.Version] = nodes
			return nil
		}
	default:
		return errors.New("un supported result action")
	}
}
