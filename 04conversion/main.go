package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Conversion")
	fmt.Println("Please rate our application from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating our app", input)

	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 5.")
		return
	}

	fmt.Println("Adding 1 to your rating...", numInput+1)

}
