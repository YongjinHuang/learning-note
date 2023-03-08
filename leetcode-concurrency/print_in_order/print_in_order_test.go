package printinorder

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
	pf := func() {
		fmt.Printf("first")
	}
	ps := func() {
		fmt.Printf("second")
	}
	pt := func() {
		fmt.Printf("third\n")
	}
	foo := NewFoo()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			foo.First(pf)
			wg.Done()
		}()
		go func() {
			foo.Second(ps)
			wg.Done()
		}()
		go func() {
			foo.Third(pt)
			wg.Done()
		}()
		wg.Wait()
	}
}
