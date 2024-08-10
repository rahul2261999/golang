package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for index := 0; index < len(days); index++ {
		fmt.Printf("Day %d: %s\n", index+1, days[index])
	}

	for index := range days {
		fmt.Println(index+1, ":", days[index])
	}

	// For each loop
	for _, value := range days {
		fmt.Println(value)
	}

	// While loop
	rogueValue := 2

	for rogueValue < 4 {

		if rogueValue == 3 {
			break
		}

		fmt.Println("Rogue value:", rogueValue)
		rogueValue++
	}
}
