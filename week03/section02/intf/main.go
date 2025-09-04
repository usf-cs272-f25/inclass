package main

import "fmt"

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func main() {
	c := NewCircle(2)
	PrintArea(c)

	r := NewRectangle(4, 5)
	PrintArea(r)
}
