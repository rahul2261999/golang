package main

import "fmt"

const LoginToken string = "your_login_token_here"

func main() {
	var username string = "rahul"

	fmt.Println(username)
	fmt.Printf("This is of type %T.\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("This is of type %T.\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("This is of type %T.\n", smallVal)

	var smallFloat float32 = 3.142333444532
	fmt.Println(smallFloat)
	fmt.Printf("This is of type %T.\n", smallFloat)

	var largeFloat float64 = 922.3372036854
	fmt.Println(largeFloat)
	fmt.Printf("This is of type %T.\n", largeFloat)

	fmt.Println(LoginToken)
	fmt.Printf("This is of type %T.\n", LoginToken)
}
