package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, i int) {
	// Decrement the waitgroup's counter
	defer wg.Done()
	fmt.Println("in worker: i = ", i)
}

// Your process terminates when the main goroutine exits
// We don't wait for child goroutines to complete, unless
// we add synchronization
func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		// increment the waitgroup's counter
		wg.Add(1)
		go worker(&wg, i)
	}

	// Wait waits for the counter to decrement to zero
	// Main goroutine blocks until counter is zero
	wg.Wait()

	fmt.Println("done main")
}
