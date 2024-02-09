package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			fmt.Println(i + j)
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

	for i, val := range arr {
		fmt.Println(i, val)
	}

	for ind, val := range "nitesh" {
		fmt.Println(ind, string(val))
	}

	x := 0

	for x < 5 {
		if x == 2 {
			goto msg
		}
		x++
	}

msg:
	fmt.Println("Hello Go")
	fmt.Println(x)


	package main

import "fmt"

func main() {

lundloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			break lundloop;
		}
	}
}


}
