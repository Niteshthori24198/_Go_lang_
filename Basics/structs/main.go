package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Salary int
}

// embeded structure

type Department struct {
	name string
	User
}

func main() {

	user := User{"Nitesh", 25, "nitesh@gmail.com", 100}

	fmt.Printf("user Name is %v , user Age is : %v\n ", user.Name, user.Age)
	fmt.Printf("user info is %+v: ", user)

	// another way to create struct

	person := struct {
		name string
	}{
		name: "abc",
	}

	fmt.Println(person)

	dept := Department{
		"IT",
		User{
			"nitesh",
			25,
			"nitesh@gmail",
			100,
		},
	}

	fmt.Println(dept.Email, dept.User.Name)

}
