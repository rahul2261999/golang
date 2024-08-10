package main

import (
	"fmt"
)

func main() {
	fmt.Println("method in go lang")

	user := User{"John Doe", "john.doe@example.com", true, 20}
	fmt.Printf("User: %+v\n", user)
	user.GetStatus()
	user.SetStatus(false)
	fmt.Println("updated status: ", user.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("status:", u.Status)
}

func (u *User) SetStatus(newStatus bool) {
	u.Status = newStatus
}
