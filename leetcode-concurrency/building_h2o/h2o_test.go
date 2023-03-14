package buildingh2o

import (
	"fmt"
	"sync"
	"testing"
)

func TestH2o(t *testing.T) {
	water := "OOHHHHOOHHHHOOHHHHOOHHHHOOHHHHOOHHHHOOHHHHOOHHHH"
	n := len(water)
	var wg sync.WaitGroup
	h2o := NewH2O()
	wg.Add(n)
	for i := 0; i < n; i++ {
		c := water[i]
		if c == 'O' {
			go func() {
				h2o.Oxygen(func() {
					fmt.Print(string('O'))
				})
				wg.Done()
			}()
		} else {
			go func() {
				h2o.Hydrogen(func() {
					fmt.Print(string('H'))
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()
}

func TestH2o2(t *testing.T) {
	water := "HHHOHOHHOHHO"
	n := len(water)
	var wg sync.WaitGroup
	h2o := NewH2O()
	wg.Add(n)
	for i := 0; i < n; i++ {
		c := water[i]
		if c == 'O' {
			go func() {
				h2o.Oxygen(func() {
					fmt.Print(string('O'))
				})
				wg.Done()
			}()
		} else {
			go func() {
				h2o.Hydrogen(func() {
					fmt.Print(string('H'))
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
