package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// generating random number using math module

	// num := rand.Intn(5) // on the base of range it will give number in the given range from 0 to 5

	// fmt.Println(num)

	// random number via crypto

	num, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(num)

}
