package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["go"] = "Go"
	languages["java"] = "Java"

	fmt.Println("Languages: ", languages)
	fmt.Println("Length of the map: ", len(languages))
	fmt.Println("Value for 'js': ", languages["js"])

}
