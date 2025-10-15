package main

import (
	"fmt"
	"sync"
)

// The order that goroutines run is not guaranteed
func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("in worker ", i)
}

// When the main goroutine terminates, you process terminates,
// no matter how many child goroutines are running, or whether
// they've run at all
func main() {
	// Create a waitgroup (thread-safe counter)
	var wg sync.WaitGroup

	var i int
	for i = range 3 {
		// Increment the counter
		wg.Add(1)
		go worker(&wg, i)
	}

	// wait for all the goroutines to finish
	// when the counter reaches zero
	wg.Wait()
	fmt.Println("done main")
}
