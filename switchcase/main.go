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

	// In go the control does not automatically fall through to next block as it usually happens in other lang like C, Java, etc. if we want this then we need to add keyword  [ fallthrough ] to move control to next block automatically.

	switch a {

	case "Hello":
		fmt.Println("Hello, World!")
	case "Nitesh":
		fmt.Println("Hello, Nitesh!")
		fallthrough
	case "Kumar":
		fmt.Println("Hello, Kumar!")
	case "Loki":
		fmt.Println("Hello, Loki!")
	default:
		fmt.Println("Hello, Guest!")
	}

}
