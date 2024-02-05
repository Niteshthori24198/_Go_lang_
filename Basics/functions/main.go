package main

import "fmt"

func main() {

	sum := add(10, 15)

	fmt.Println("Sum is : ", sum)

	arr := []int{1, 2, 3, 4, 5}

	msg, res := bigSum(arr...)

	fmt.Print(msg, res)

}

func add(a int, b int) int {
	return a + b
}

// variadic function
func bigSum(values ...int) (string, int) {

	total := 0

	for _, val := range values {
		total += val
	}

	return "value is :", total
}
