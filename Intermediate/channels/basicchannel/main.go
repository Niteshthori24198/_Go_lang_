package main

import "fmt"
import "sync"

type Indicator struct {
	name string
	val  int
}

var hugna1 bool
var hugna2 bool

func main() {

	wg := &sync.WaitGroup{}
	wg1 := &sync.WaitGroup{}

	// channel is a data structure that allows us to send and receive data in go lang. channel helps us to communicate between different go routines.

	myCh := make(chan Indicator, 5) // default size is 1 and we can increase it by making buffered channel.

	myCh2 := make(chan Indicator, 5)

	// this will create a deadlock as the channel is not being used. In channel we are listening but not sending and vice versa.

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg1.Add(2)

	// Note :- In general we can send or receive values from channel. but we can also make then either send or receive only.

	go func(myCh chan Indicator, wg1 *sync.WaitGroup) {

		
		myCh <- Indicator{name: "R-one", val: 1}
		hugna1 = true

		for i := 1; !hugna2; i++ {}

		fmt.Println("1=>: ", <-myCh)
		fmt.Println("1=>: ", <-myCh)

		myCh <- Indicator{name: "R-one", val: 2}
		
		fmt.Println("1 : ", <-myCh)
		wg1.Done()
	}(myCh, wg1)

	go func(myCh chan Indicator, wg1 *sync.WaitGroup) {

		for !hugna1 {}

		fmt.Println("2: ", <-myCh)

		myCh <- Indicator{name: "R-two", val: 3}
		myCh <- Indicator{name: "R-two", val: 4}

		hugna2 = true
		wg1.Done()
	}(myCh, wg1)

	wg.Add(2)

	// send only channel

	go func(myCh2 chan<- Indicator, wg *sync.WaitGroup) {

		myCh2 <- Indicator{name: "R-three", val: 5}
		myCh2 <- Indicator{name: "R-three", val: 6}

		wg.Done()
	}(myCh2, wg)

	// receive onyl channel

	go func(myCh2 <-chan Indicator, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh2
		val2, isChannelOpen1 := <-myCh2

		fmt.Println(val, isChannelOpen)
		fmt.Println(val2, isChannelOpen1)

		wg.Done()
	}(myCh2, wg)

	wg1.Wait()
	wg.Wait()
}
