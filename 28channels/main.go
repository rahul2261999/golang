package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go!")

	myChannels := make(chan int)
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)

		defer waitGroup.Done()
	}(myChannels, waitGroup)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 1
		ch <- 2
		close(ch)
		defer waitGroup.Done()

	}(myChannels, waitGroup)

	waitGroup.Wait()

}
