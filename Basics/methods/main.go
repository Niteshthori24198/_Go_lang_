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
	fmt.Printf("user info is %+v\n ", user)

	user.GetUserInfo()

	msg, email := user.GetEmail()

	fmt.Println(msg, email)

	user.updateEmail("thori@gmail.com")

	fmt.Print("Original user : ", user)

}

func (u User) GetUserInfo() {

	fmt.Printf("User info : \n Name is %v , Age is : %v , Email is : %v , Salary is : %v\n ", u.Name, u.Age, u.Email, u.Salary)
}

func (u User) GetEmail() (string, string) {

	return "user email is : ", u.Email
}

func (u User) updateEmail(newEmail string) {

	u.Email = newEmail

	fmt.Println("Updated user : ", u)
}
