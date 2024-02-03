package main

import "fmt"

func main() {

	c := myFun()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	d := myFun()

	fmt.Println(d())
	fmt.Println(d())

}

func myFun() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
