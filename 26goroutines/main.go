package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var signals []string
var waitGoup sync.WaitGroup

func main() {
	fmt.Println("Goroutines in Go!")

	urls := []string{
		"https://www.velotio.com",
		"https://www.openai.com",
		"https://www.velotio.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
	}

	for _, url := range urls {
		go getStatusCode(url)
		waitGoup.Add(1)
	}

	waitGoup.Wait()

	fmt.Println("Received signals from:", len(signals))

	fmt.Println("Received signals from:", signals)

	fmt.Println("All URLs checked!")
}

func getStatusCode(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, url)
		mutex.Unlock()

		fmt.Printf("StatusCode: %d -> %s\n", res.StatusCode, url)
	}

	defer waitGoup.Done()
}
