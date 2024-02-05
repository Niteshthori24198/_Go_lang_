package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to Files is Go Lang :- ")

	// create a new file using os module with file path added init

	file, err := os.Create("./temp/hello.txt")

	checkNilError(err)

	// write some text in file for this we use io module

	reader := bufio.NewReader(os.Stdin)

	content, _ := reader.ReadString('\n')

	checkNilError(err)

	datalength, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Printf("We wrote %v bytes into the file \n", datalength)

	// reading data from file

	readFile("./temp/hello.txt")

	defer file.Close()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {

	// to read file we use ioutil module or os module

	dataByte , err := os.ReadFile(file)

	checkNilError(err)

	fmt.Println("Text data inside the file is : ", string(dataByte))
}
