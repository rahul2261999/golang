package main

import "fmt"

func main() {
	fmt.Println("Func in golang")
	greet("John Doe")

}

func greet(name string) {
	fmt.Println("Hello,", name)
}
