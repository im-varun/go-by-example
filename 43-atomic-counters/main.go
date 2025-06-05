package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicIncrements() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)

		go func() {
			for range 1000 {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter =", ops.Load())
}

func nonAtomicIncrements() {
	var ops int

	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)

		go func() {
			for range 1000 {
				ops++
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter =", ops)
}

func main() {
	fmt.Println("after atomic increments:")
	atomicIncrements()

	fmt.Println("after non-atomic increments:")
	nonAtomicIncrements()
}
