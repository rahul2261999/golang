package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	loginCount := 2
	if loginCount >= 5 {
		fmt.Println("Welcome to the system")
	} else {
		fmt.Println("Access denied")
	}

}
