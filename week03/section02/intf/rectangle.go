package main

type Rectangle struct {
	Height, Width float32
}

func NewRectangle(height, width float32) *Rectangle {
	r := new(Rectangle)
	r.Height = height
	r.Width = width
	return r
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return (2 * r.Height) + (2 * r.Width)
}
