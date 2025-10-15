package main

import (
	"fmt"
)

type Pair struct {
	i, j int
}

func Adder(ch chan Pair, chSum chan int, chDone chan bool) {
	for pair := range ch {
		chSum <- pair.i + pair.j
	}
	fmt.Println("adder exiting")
	chDone <- true
}

func main() {
	ch := make(chan Pair, 3)
	chSum := make(chan int, 3)
	chDone := make(chan bool, 2)
	go Adder(ch, chSum, chDone)

	for i := range 3 {
		ch <- Pair{i, i + 1}
	}

	go func() {
		for sum := range chSum {
			fmt.Println("sum: ", sum)
		}
		fmt.Println("sum loop exiting")
		chDone <- true
	}()

	close(ch)
	<-chDone
	close(chSum)

	// Block until there's data on chDone
	<-chDone

	fmt.Println("done main")
}
