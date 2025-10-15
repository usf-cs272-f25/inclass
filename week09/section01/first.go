package main

import "fmt"

// The order that goroutines run is not guaranteed
func worker(i int) {
	fmt.Println("in worker ", i)
}

// When the main goroutine terminates, you process terminates,
// no matter how many child goroutines are running, or whether
// they've run at all
func main() {
	for i := range 3 {
		go worker(i)
	}
	fmt.Println("done main")
}
