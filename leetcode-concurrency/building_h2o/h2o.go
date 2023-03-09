package buildingh2o

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	ch chan struct{}
}

// Acquire implements Semaphore
func (s *semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release implements Semaphore
func (s *semaphore) Release() {
	<-s.ch
}

func newSemaphore(n int) Semaphore {
	return &semaphore{
		ch: make(chan struct{}, n),
	}
}

type H2O struct {
	hsema Semaphore
	osema Semaphore
}

func NewH2O() *H2O {
	return &H2O{
		hsema: newSemaphore(2),
		osema: newSemaphore(0),
	}
}

func (h2o *H2O) Hydrogen(h func()) {
	h2o.hsema.Acquire()
	h()
	h2o.osema.Release()
}

func (h2o *H2O) Oxygen(o func()) {
	h2o.osema.Acquire()
	h2o.osema.Acquire()
	// h2o.osema.Acquire(2)
	o()
	// h2o.hsema.Release(2)
	h2o.hsema.Release()
	h2o.hsema.Release()
}
