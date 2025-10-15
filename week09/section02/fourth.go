package main

import (
	"fmt"
)

func worker(ch chan int, chDone chan bool) {
	// Range over channels runs until the channel closes
	// NOT when there's nothing to read - that just sleeps
	for i := range ch {
		fmt.Println("in worker: i = ", i)
		// Bank balance updated here...
	}
	chDone <- true
}

func main() {
	// A channel is a thread-safe queue
	// Delivery is in-order (unlike scheduling goroutines)
	ch := make(chan int, 3)
	chDone := make(chan bool)

	go worker(ch, chDone)
	for i := range 3 {
		// Insert i into the channel, to be read by worker
		ch <- i
	}
	close(ch)

	// Block until there's data on chDone
	<-chDone

	fmt.Println("done main")
}
