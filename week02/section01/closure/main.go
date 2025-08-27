package main

import "fmt"

func newSeq() func() int {
	// newSeq is a "closure" because it "closes over" the variable i
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

	t := newSeq()
	t()
	t()
}
