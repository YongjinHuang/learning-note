package printzeroevenodd

import (
	"fmt"
	"sync"
	"testing"
)

func TestFive(t *testing.T) {
	printNumber := func(n int) {
		fmt.Print(n)
	}
	for i := 2; i <= 9; i++ {
		fmt.Println("i", i)
		zeroEvenOdd := NewZeroEvenOdd(i)
		var wg sync.WaitGroup
		wg.Add(3)
		// Thread A
		go func() {
			zeroEvenOdd.Zero(printNumber)
			wg.Done()
		}()
		// Thread B
		go func() {
			zeroEvenOdd.Even(printNumber)
			wg.Done()
		}()
		// Thread C
		go func() {
			zeroEvenOdd.Odd(printNumber)
			wg.Done()
		}()
		wg.Wait()
	}

}
