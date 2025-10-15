package main

import "fmt"

// The order that goroutines run is not guaranteed
func worker(ch chan bool, i int) {
	fmt.Println("in worker ", i)
	ch <- true
}

// When the main goroutine terminates, you process terminates,
// no matter how many child goroutines are running, or whether
// they've run at all
func main() {
	// Create a channel which can hold one bool
	// channels are a thread-safe queue
	ch := make(chan bool, 3)
	var i int
	for i = range 3 {
		go worker(ch, i)
	}

	// Read the channel, which blocks until there is somthing to read
	<-ch
	<-ch
	<-ch

	fmt.Println("done main")
}
