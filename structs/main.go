package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Salary int
}

func main() {

	user := User{"Nitesh", 25, "nitesh@gmail.com", 100}

	fmt.Printf("user Name is %v , user Age is : %v\n ", user.Name, user.Age)
	fmt.Printf("user info is %+v: ", user)

}
