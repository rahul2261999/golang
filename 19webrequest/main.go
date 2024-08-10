package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://courses.learncodeonline.in/learn"

func main() {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is: %T\n", res)

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(bodyBytes)
	fmt.Println(content)
}
