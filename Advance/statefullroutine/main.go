package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	fmt.Println("Statefull go Routines !")

	var readOprCount uint64
	var writeOprCount uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// goroutine for managining stateful go routines

	go func() {

		var state = make(map[int]int)

		for {

			select {

			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}

	}()

	// create read request operations

	for i := 0; i < 100; i++ {

		go func() {

			for {

				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOprCount, 1)
				time.Sleep(time.Millisecond)
			}

		}()

	}

	// write routines

	for i := 0; i < 100; i++ {

		go func() {

			for {

				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}

				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOprCount, 1)
				time.Sleep(time.Millisecond)
			}

		}()

	}

	time.Sleep(time.Second * 1)

	fmt.Println("read opr count : ", atomic.LoadUint64(&readOprCount))

	fmt.Println("write opr count : ", atomic.LoadUint64(&writeOprCount))

}
