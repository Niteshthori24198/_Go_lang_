package main

import "fmt"

func main() {

	// var arr = [5]string{"nitesh", "kumar", "kushwaha", "go", "lang"}

	arr := [5]string{"nitesh", "kumar", "kushwaha", "go", "lang"}

	for _, val := range arr {
		fmt.Println(val)
	}

	var b [2][3]int;

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = i + j
		}
	}

	fmt.Println(b)

	var c [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)
}
