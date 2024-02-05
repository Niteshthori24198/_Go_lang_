package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")
	defer fmt.Println("Hello, Bachoo!")
	defer fmt.Println("Lets, Go!")

	loopChalo()
}

func loopChalo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
