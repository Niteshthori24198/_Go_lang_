package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race Condition :- ")

	arr := []int{0}

	wg := &sync.WaitGroup{}
	mutx := &sync.RWMutex{}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {

		fmt.Println("One R")

		mutx.Lock()
		arr = append(arr, 1)
		mutx.Unlock()

		wg.Done()

	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {

		fmt.Println("Two R")

		mutx.Lock()
		arr = append(arr, 2)
		mutx.Unlock()

		wg.Done()

	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {

		fmt.Println("Three R")

		mutx.Lock()
		arr = append(arr, 3)
		mutx.Unlock()

		wg.Done()

	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {

		mutx.RLock()
		fmt.Println(arr)
		mutx.RUnlock()

		wg.Done()
	}(wg, mutx)

	wg.Wait()
	fmt.Println(arr)

}
