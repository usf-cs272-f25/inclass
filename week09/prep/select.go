package main

import (
	"fmt"
)

func process(ch1 chan int, ch2 chan int, chDone chan bool) {
	remaining := 2
	for {
		select {
		case c1 := <-ch1:
			fmt.Println(c1)
			remaining--
		case c2 := <-ch2:
			fmt.Println(c2)
			remaining--
		}
		if remaining == 0 {
			chDone <- true
		}
	}
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	chDone := make(chan bool)

	go process(ch1, ch2, chDone)

	ch1 <- 1
	ch2 <- 2

	<-chDone
}
