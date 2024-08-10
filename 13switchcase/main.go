package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch statement in golang")

	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Dice number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
	case 3:
		fmt.Println("You rolled a 3")
	case 4:
		fmt.Println("You rolled a 4")
		fallthrough
	case 5:
		fmt.Println("You rolled a 5")
	case 6:
		fmt.Println("You rolled a 6")
	default:
		fmt.Println("Invalid dice number")
	}

}
