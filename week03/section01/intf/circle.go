package main

import "math"

type Circle struct {
	Radius float32
}

// Factory function for circles

func NewCircle(radius float32) *Circle {
	c := new(Circle)
	c.Radius = radius
	return c
}

// Implementation of Shape interface in terms of circle

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2.0 * math.Pi * c.Radius
}
