package main

import "fmt"

func PrintArea(s Shape) {
	a := s.Area()
	fmt.Println("Area: ", a)
}

func main() {
	c := NewCircle(2.0)
	PrintArea(c)

	r := NewRectangle(3.0, 4.0)
	PrintArea(r)
}
