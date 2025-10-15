package main

import (
	"fmt"
	"time"
)

// This is a long-running goroutine which ranges over
// a channel to get its inputs
func worker(ch chan int) {
	// the range loop stops when the channel gets closed
	// NOT when there's no more data in the channel
	for i := range ch {
		fmt.Println("in worker i = ", i)
	}
	fmt.Println("done worker")
}

func main() {
	// Create a channel which can hold 3 integers
	// channels are a thread-safe queue
	ch := make(chan int, 3)
	go worker(ch)

	for i := range 3 {
		ch <- i
	}

	time.Sleep(10 * time.Millisecond)
	close(ch)
	fmt.Println("done main")
}
