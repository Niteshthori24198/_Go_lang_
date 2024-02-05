package main

import "fmt"

func main() {
	fmt.Println("Variables in Go !")

	var name string = "nitesh"
	var isValid bool = true

	fmt.Println(name)
	fmt.Println(isValid)
	fmt.Printf("variable type : %T\n",name);

	// default value
	var text float32
	fmt.Println(text)
	
	// implicit type

	var age = 25
	fmt.Println(age)

	// no var keyword [ we can use this only inside method only. outside it will give error. ]

	isEmployed := true
	fmt.Println(isEmployed)
}