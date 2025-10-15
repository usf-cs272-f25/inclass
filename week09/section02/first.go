package main

import (
	"fmt"
	"time"
)

func worker2(i int) {
	fmt.Println("in worker2: i = ", i)
}

func worker(i int) {
	go worker2(i)
	fmt.Println("in worker: i = ", i)
}

// Your process terminates when the main goroutine exits
// We don't wait for child goroutines to complete, unless
// we add synchronization
func main() {
	for i := range 3 {
		go worker(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done main")
}
