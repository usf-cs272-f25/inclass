package main

import "fmt"

// Drivable is (sort of) a base class for Car
type Drivable struct {
	Wheels uint
}

type Car struct {
	Drivable // Drivable is "embedded" in Car
	Make     string
	Model    string
	Year     uint
}

// NewCar is a "factory function"
func NewCar(mak, model string, year, wheels uint) *Car {
	c := new(Car)
	c.Make = mak
	c.Model = model
	c.Year = year
	c.Wheels = wheels
	return c
}

func (c *Car) Print() {
	fmt.Printf("Car: make: %s, model: %s, year: %d\n", c.Make, c.Model, c.Year)
}
