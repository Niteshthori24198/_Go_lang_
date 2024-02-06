package main

import "fmt"
import "sync"

func main() {

	wg := &sync.WaitGroup{}

	// channel is a data structure that allows us to send and receive data in go lang. channel helps us to communicate between different go routines.

	myCh := make(chan int, 5) // default size is 1 and we can increase it by making buffered channel.

	myCh2 := make(chan string, 5)

	// this will create a deadlock as the channel is not being used. In channel we are listening but not sending and vice versa.

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(4)

	// Note :- In general we can send or receive values from channel. but we can also make then either send or receive only.

	go func(myCh chan int, wg *sync.WaitGroup) {

		myCh <- 15

		fmt.Println("1: ", <-myCh)
		fmt.Println("1: ", <-myCh)

		myCh <- 20

		fmt.Println("1 : ", <-myCh)

		wg.Done()
	}(myCh, wg)

	go func(myCh chan int, wg *sync.WaitGroup) {

		fmt.Println("2: ", <-myCh)

		myCh <- 5
		myCh <- 10

		wg.Done()
	}(myCh, wg)

	// send only channel

	go func(myCh2 chan<- string, wg *sync.WaitGroup) {

		myCh2 <- "hello"
		myCh2 <- "World"

		wg.Done()
	}(myCh2, wg)

	// receive onyl channel

	go func(myCh2 <-chan string, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh2
		val2, isChannelOpen1 := <-myCh2

		fmt.Println(val, isChannelOpen)
		fmt.Println(val2, isChannelOpen1)

		wg.Done()
	}(myCh2, wg)

	wg.Wait()
}
