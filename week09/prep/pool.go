package main

import (
	"fmt"
	"time"
)

type Pair struct {
	i, j int
}

func AddPairs(chPairs chan Pair, chSum chan int) {
	for pair := range chPairs {
		chSum <- pair.i + pair.j
	}
	// for pair := range chPairs {
	// 	sum := pair.i + pair.j
	// 	chSum <- sum
	// }
	fmt.Println("AddPairs exiting")
}

func main() {

	const depth int = 30
	chIn := make(chan Pair, depth)
	chOut := make(chan int, depth)

	go AddPairs(chIn, chOut)
	for i := range depth {
		chIn <- Pair{i, i + 1}
	}

	for j := range depth {
		sum := <-chOut
		fmt.Printf("iter %d sum %d\n", j, sum)
	}

	close(chIn)
	time.Sleep(10 * time.Millisecond)
}
