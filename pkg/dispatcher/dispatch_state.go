package dispatcher

type dispatchState struct {
	session *session
	handler DispatchHandler
	err     error
}

func (ds *dispatchState) clear() {
	ds.session = nil
	ds.err = nil
}

func (ds *dispatchState) invokeHandler(p interface{}) {
	ds.err = ds.handler(p)
	ds.session.completed <- ds
}
