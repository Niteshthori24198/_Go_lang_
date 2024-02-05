package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	
	fmt.Println("Reading input : ")

	reader := bufio.NewReader(os.Stdin)

	// comma , err syntax is used to read input

	text , _ := reader.ReadString('\n')

	fmt.Println(text)

}