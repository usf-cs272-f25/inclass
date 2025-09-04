package main

type Rectangle struct {
	Width, Height float32
}

func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{
		width,
		height,
	}
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float32 {
	return (2 * r.Height) + (2 * r.Width)
}
