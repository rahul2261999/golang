package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the study of time")

	presentTime := time.Now()
	fmt.Println("Current date:", presentTime)
	fmt.Println("Format date:", presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)
	fmt.Println("Created date:", createdDate)
}
