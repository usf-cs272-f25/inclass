package main

import "fmt"

// Steerable has the same Wheels attribute. Use type name to disambiguate
type Steerable struct {
	Wheels uint
}

// Drivable is sort of a base class for cars, trucks, bikes, etc.
type Drivable struct {
	Wheels uint
}

type Car struct {
	Drivable // Drivable is "embedded" in Car
	Steerable
	Make  string
	Model string
	Year  uint
}

// Factory function
func NewCar(mak, model string, year, wheels uint) *Car {
	c := new(Car)
	c.Make = mak
	c.Model = model
	c.Year = year
	c.Drivable.Wheels = wheels // we can use Wheels directly, as though it was a member of Car
	return c
}

func (c *Car) Print() {
	fmt.Printf("Car: make: %s, model: %s, year: %d, wheels: %d\n", c.Make, c.Model, c.Year, c.Drivable.Wheels)
}
