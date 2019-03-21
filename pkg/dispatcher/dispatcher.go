package dispatcher

import (
	"sync"

	"github.com/hb-go/micro-mesh/pkg/pool"
)

type DispatchHandler func(interface{}) error

type Dispatcher struct {
	// pool of sessions
	sessionPool sync.Pool

	// pool of dispatch states
	statePool sync.Pool

	// pool of goroutines
	gp *pool.GoroutinePool
}

func NewDispatcher(handlerGP *pool.GoroutinePool) *Dispatcher {
	d := &Dispatcher{
		gp: handlerGP,
	}

	d.sessionPool.New = func() interface{} { return &session{} }
	d.statePool.New = func() interface{} { return &dispatchState{} }
	return d
}

func (d *Dispatcher) Dispatch(handlers ...DispatchHandler) error {
	s := d.getSession()

	s.handlers = handlers
	err := s.dispatch()

	d.putSession(s)
	return err
}

func (d *Dispatcher) getSession() *session {
	s := d.sessionPool.Get().(*session)
	s.dispatcher = d
	return s
}

func (d *Dispatcher) putSession(s *session) {
	s.clear()
	d.sessionPool.Put(s)
}

func (d *Dispatcher) getDispatchState() *dispatchState {
	ds := d.statePool.Get().(*dispatchState)

	return ds
}

func (d *Dispatcher) putDispatchState(ds *dispatchState) {
	ds.clear()
	d.statePool.Put(ds)
}
