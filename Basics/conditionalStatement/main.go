package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	a, _ := reader.ReadString('\n')

	a = strings.TrimSpace(a)

	// we can assign values to variable in if condition also and then apply checking condition

	if b := a; b == "1" {
		fmt.Println("nitesh")
	} else {
		fmt.Println("kumar", b)
	}

}
