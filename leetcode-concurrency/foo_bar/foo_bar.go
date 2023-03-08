package foobar

import "fmt"

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

func newSemaphore(concurrency int) Semaphore {
	return &semaphore{
		ch: make(chan struct{}, concurrency),
	}
}

type FooBar struct {
	N  int
	Cf Semaphore
	Cb Semaphore
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		N:  n,
		Cf: newSemaphore(1),
		Cb: newSemaphore(0),
	}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.N; i++ {
		fb.Cf.Acquire()
		fmt.Print("foo")
		fb.Cb.Release()
	}
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.N; i++ {
		fb.Cb.Acquire()
		fmt.Print("bar\n")
		fb.Cf.Release()
	}
}
