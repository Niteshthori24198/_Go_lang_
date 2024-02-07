package main

import "fmt"
import "errors"

// custom error type

type myError struct {
	val int
	msg string
}

func (e *myError) Error() string {
	return e.msg
}

func main() {

	fmt.Println("Error in Golang !")

	for _, i := range []int{7, 10} {

		res, err := f1(i)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}

	for _, i := range []int{7, 10} {

		res, err := f2(i)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}

}

func f1(num int) (int, error) {

	if num == 10 {

		return -1, errors.New("Number Can't be 10")

	}

	return num * 2, nil
}

func f2(num int) (int, error) {

	if num == 10 {

		return -1, &myError{num, "Number Can't be 10"}

	}

	return num * 2, nil
}
