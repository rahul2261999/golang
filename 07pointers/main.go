package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")

	num := 23

	ptr := &num

	fmt.Println("Memory address of num:", ptr)
	fmt.Println("Value at memory address:", *ptr)

	*ptr = *ptr + 1

	fmt.Println("Modified value at memory address:", num)

}
