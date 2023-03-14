package fz

import (
	"fmt"
	"sync"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzz := NewFizzBuzz(31)
	var wg sync.WaitGroup
	wg.Add(4)
	// thread A
	go func() {
		fizzBuzz.Fizz(func() {
			fmt.Println("fizz")
		})
		wg.Done()
	}()
	// thread B
	go func() {
		fizzBuzz.Buzz(func() {
			fmt.Println("buzz")
		})
		wg.Done()
	}()
	// thread C
	go func() {
		fizzBuzz.FizzBuzz(func() {
			fmt.Println("fizzbuzz")
		})
		wg.Done()
	}()
	// thread D
	go func() {
		fizzBuzz.Number(func(n int) {
			fmt.Println(n)
		})
		wg.Done()
	}()
	wg.Wait()
}
