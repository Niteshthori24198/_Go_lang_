package main

import (
	"fmt"
	"time"
)

func main() {

	// we use go keyword to create a new goroutine which is a lightweight thread in go lang which can run concurrently with main thread in go lang and can share memory with main thread in go lang. but in order to see result we need to wait for main thread to finish its execution.

	go greet("Hello, Golang!")
	greet("Hello, World!")
}

func greet(msg string) {

	for i := 0; i < 5; i++ {

		// we will wait for sometime to see the result

		time.Sleep(2 * time.Second)
		fmt.Println(msg)
	}
}
