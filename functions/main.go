package main

import "fmt"

func main() {

	sum := add(10, 15)

	fmt.Println("Sum is : ", sum)

	msg, res := bigSum(1, 2, 3, 4, 5)

	fmt.Print(msg, res)

}

func add(a int, b int) int {
	return a + b
}

func bigSum(values ...int) (string, int) {

	total := 0

	for _, val := range values {
		total += val
	}

	return "Total sum is : ", total
}
