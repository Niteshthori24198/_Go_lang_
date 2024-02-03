package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var arr = []int{10, 15, 12, 14, 13}
	var b = []int{11, 17, 16}

	arr = append(arr, b...)

	fmt.Println(arr)
	fmt.Println(arr[1:3])
	fmt.Println(arr[:3])

	sort.Ints(arr)
	fmt.Println(arr)

	var list = make([]string, 4)

	list[0] = "nitesh"
	list[1] = "kumar"
	list[2] = "go"
	list[3]= "lang"

	list = append(list, "kushwaha")

	fmt.Println(list, list[2], len(list))

	reader := bufio.NewReader(os.Stdin)

	val, _ := reader.ReadString('\n')

	val = strings.TrimSpace(val)

	index, _ := strconv.ParseInt(val, 10, 64)

	list = append(list[:index], list[index+1:]...)

	fmt.Println("New list ", list)

}
