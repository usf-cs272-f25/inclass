package main

import "fmt"

func newSeq() func() int {
	i := 0

	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}

func main() {
	s := newSeq()

	s()
	s()
}
