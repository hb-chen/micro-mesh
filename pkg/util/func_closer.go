package util

type FuncCloser struct {
	CloseFunc func() error
}

func (c *FuncCloser) Close() error {
	return c.CloseFunc()
}
