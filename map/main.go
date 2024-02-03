package main

import (
	"fmt"
)

func main() {

	language := make(map[string]string)

	language["en"] = "English"
	language["hi"] = "Hindi"
	language["es"] = "Spanish"

	fmt.Println(language)

	delete(language, "es")
	fmt.Println(language)

	for key,val := range language{
		fmt.Printf("the key is %v and value is %v\n",key,val)
	}
}
