package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {

	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

type myType interface {
	int | float64
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

type values interface {
	int | string | float64
}

type customMap[T comparable, V values] map[T]V

func main() {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

	val := printBigger(1, 2)
	val2 := printBigger(1.1, 2.2)

	fmt.Println(val, val2)

	m1 := make(customMap[int, string])
	m2 := make(customMap[int, float64])

	m1[1] = "hello"
	m2[1] = 3.14

	fmt.Println(m1, m2)

}

func printBigger[T myType](a T, b T) T {

	if a > b {
		return a
	}

	return b
}
