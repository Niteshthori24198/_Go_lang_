package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func main() {

	// We implement a comparison function for string lengths. cmp.Compare is helpful for this.

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length.

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	peoples := []Person{
		{"Nitesh", 25},
		{"Kumar", 26},
		{"Kushwaha", 21},
	}

	// We can use cmp.Compare to sort people by age.

	slices.SortFunc(peoples, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})


	fmt.Println(peoples)
}
