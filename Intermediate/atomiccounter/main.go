package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.UInt64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 100; c++ {

				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
