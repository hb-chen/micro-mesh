package dispatcher

import (
	"log"

	"github.com/hashicorp/go-multierror"
)

const queueAllocSize = 64

type session struct {
	dispatcher *Dispatcher
	handlers   []DispatchHandler

	activeDispatches int
	completed        chan *dispatchState

	err error
}

func (s *session) clear() {
	s.dispatcher = nil
	s.activeDispatches = 0
	s.handlers = nil
	s.err = nil

	// Drain the channel
	exit := false
	for !exit {
		select {
		case <-s.completed:
			log.Printf("Leaked dispatch state discovered!")
			continue
		default:
			exit = true
		}
	}
}

func (s *session) ensureParallelism(minParallelism int) {
	// Resize the channel to accommodate the parallelism, if necessary.
	if cap(s.completed) < minParallelism {
		allocSize := ((minParallelism / queueAllocSize) + 1) * queueAllocSize
		s.completed = make(chan *dispatchState, allocSize)
	}
}

func (s *session) dispatch() error {
	s.ensureParallelism(len(s.handlers))

	for _, h := range s.handlers {
		ds := s.dispatcher.getDispatchState()
		ds.handler = h
		s.dispatchToHandler(ds)
	}

	s.waitForDispatched()
	return s.err
}

func (s *session) dispatchToHandler(ds *dispatchState) {
	s.activeDispatches++
	ds.session = s
	s.dispatcher.gp.ScheduleWork(ds.invokeHandler, nil)
}

func (s *session) waitForDispatched() {
	for s.activeDispatches > 0 {
		state := <-s.completed
		s.activeDispatches--

		if state.err != nil {
			s.err = multierror.Append(s.err, state.err)
		}

		s.dispatcher.putDispatchState(state)
	}
}
