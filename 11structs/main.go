package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs")

	user := User{"Rahul", "rs@abc.com", true, 25}
	fmt.Println("User: ", user)
	fmt.Printf("My details are: %+v\n", user)
	fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", user.Name, user.Email, user.Status, user.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
