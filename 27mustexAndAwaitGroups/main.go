package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Await Groups in Go!")

	var wg = &sync.WaitGroup{}
	var mutex = &sync.RWMutex{}

	scores := []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		mutex.Lock()
		fmt.Println("One Run")
		scores = append(scores, 1)
		mutex.Unlock()

		defer wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		mutex.Lock()
		fmt.Println("Two Runs")
		scores = append(scores, 2)
		mutex.Unlock()

		defer wg.Done()

	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		mutex.Lock()
		fmt.Println("Three Runs")
		scores = append(scores, 3)
		mutex.Unlock()
		defer wg.Done()

	}(wg, mutex)

	wg.Wait()
	fmt.Println("Final Scores:", scores)
}
