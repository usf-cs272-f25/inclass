package main

import "fmt"

// Using iota with a sequence of constants behaves like C enum
// The compiler increments the value of each element
const (
	white = iota
	black
	grey
)

type Color int

func main() {
	// Any primitive (built-in) type can be constant
	const foo string = "foo"
	fmt.Println(foo)

	fmt.Println(white)
	fmt.Println(black)
	fmt.Println(grey)

	var c Color = white
	fmt.Println(c)

	// Constants here are untyped integers, can't assign to a string
	// var s string = black
	// fmt.Println(s)
}
