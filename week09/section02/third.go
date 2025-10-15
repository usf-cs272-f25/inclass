package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	// Range over channels runs until the channel closes
	// NOT when there's nothing to read - that just sleeps
	for i := range ch {
		fmt.Println("in worker: i = ", i)
	}
}

func main() {
	// A channel is a thread-safe queue
	// Delivery is in-order (unlike scheduling goroutines)
	ch := make(chan int, 3)
	go worker(ch)
	for i := range 3 {
		// Insert i into the channel, to be read by worker
		ch <- i
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done main")
}
