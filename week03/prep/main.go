package main

import "fmt"

func PrintShape(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	c := Circle{Radius: 2.0}
	PrintShape(c)
}
