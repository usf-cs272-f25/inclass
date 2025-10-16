package main

import (
	"fmt"
	"time"
)

type Pair struct {
	i, j int
}

// This is a long-running goroutine which ranges over
// a channel to get its inputs
func Adder(chIn chan Pair, chOut chan int) {
	// the range loop stops when the channel gets closed
	// NOT when there's no more data in the channel
	for pair := range chIn {
		chOut <- pair.i + pair.j
	}
	fmt.Println("done worker")
}

// chOut becomes input to the Printer goroutine, forming a pipeline
func Printer(ch chan int) {
	for i := range ch {
		fmt.Println("i: ", i)
	}
}

func main() {
	// Create a channel which can hold 3 integers
	// channels are a thread-safe queue
	chPair := make(chan Pair, 3)
	chSum := make(chan int, 3)
	go Adder(chPair, chSum)
	go Printer(chSum)

	for p := range 3 {
		chPair <- Pair{p, p + 1}
	}

	time.Sleep(10 * time.Millisecond)
	close(chPair)
	time.Sleep(10 * time.Millisecond)
	close(chSum)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("done main")
}
