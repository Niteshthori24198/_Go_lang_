package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	a, _ := reader.ReadString('\n')

	b, err := strconv.ParseInt(strings.TrimSpace(a), 2, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}

}
