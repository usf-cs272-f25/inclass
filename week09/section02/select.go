package main

import (
	"fmt"
)

func process(ch1 chan int, ch2 chan int, done chan bool) {
	remaining := 2
	for {
		// You can use select as a "traffic cop" to monitor
		// a group of channels, where the case of each one
		// executes when there's data on that channel to read
		select {
		case c1 := <-ch1:
			fmt.Println("c1: ", c1)
			remaining--
		case c2 := <-ch2:
			fmt.Println("c2: ", c2)
			remaining--
		}

		// No more work to do, signal the main goroutine
		// that we're done and it can terminate
		if remaining == 0 {
			done <- true
		}
	}
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	chDone := make(chan bool)

	// Spin up a goroutine to read ch1 and ch2
	go process(ch1, ch2, chDone)

	// Write data into channels
	ch1 <- 1
	ch2 <- 2

	// Block main goroutine until there's data on the done channel
	<-chDone
	fmt.Println("main exiting")
}
