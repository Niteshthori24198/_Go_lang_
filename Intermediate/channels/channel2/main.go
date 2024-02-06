package main

import "fmt"

func main() {

	ch := make(chan int, 5)
	
	go myFunc(ch)

	for val := range ch {
		fmt.Println(val)
	}

}

func myFunc(ch chan int) {

	for i := 0; i < 5; i++ {
		ch <- i + 1
	}

	close(ch)
}
