package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("defer in go lang")
	defer fmt.Println("three")
}
