package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files in go lang")

	content := "Hey! how are you?"

	file, err := os.Create("sample.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()

	readFromFile("sample.txt")
}

func readFromFile(fileName string) {
	dataByete, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("File content:", string(dataByete))
}
