package printzeroevenodd

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

type ZeroEvenOdd struct {
	N  int
	sz Semaphore
	se Semaphore
	so Semaphore
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	return &ZeroEvenOdd{
		N:  n,
		sz: newSemaphore(1),
		se: newSemaphore(0),
		so: newSemaphore(0),
	}
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.N; i++ {
		z.sz.Acquire()
		printNumber(0)
		if i&1 == 0 {
			z.se.Release()
		} else {
			z.so.Release()
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 2; i <= z.N; i += 2 {
		z.se.Acquire()
		printNumber(i)
		z.sz.Release()
	}

}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.N; i += 2 {
		z.so.Acquire()
		printNumber(i)
		z.sz.Release()
	}
}
