package foobar

import (
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
	fooBar := NewFooBar(200)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fooBar.Foo()
		wg.Done()
	}()
	go func() {
		fooBar.Bar()
		wg.Done()
	}()
	wg.Wait()
}
