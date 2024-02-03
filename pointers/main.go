package main

import "fmt"


func main(){

	num := 25

	var ptr *int = &num

	*ptr = *ptr * 2

	fmt.Println(num)

}