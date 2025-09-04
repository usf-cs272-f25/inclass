package main

/*
A go "interface" is an abstract base class

	which can be implemented in terms of multiple
	derived classes
*/
type Shape interface {
	Area() float32
	Perimeter() float32
}
