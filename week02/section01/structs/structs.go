package main

import "fmt"

type Car struct {
	Make  string
	Model string
	Year  int
}

// Function to print car details
// Passing a copy of the struct doesn't matter because
// the data cache is fast enough for this example.
func printCar(cars []Car) {
	for _, c := range cars {
		fmt.Printf("Car Make: %s, Model: %s, Year: %d\n", c.Make, c.Model, c.Year)
	}
}

func findCar(cars []Car, model string) *Car {
	for i, c := range cars {
		if c.Model == model {
			car := Car{cars[i].Make, cars[i].Model, cars[i].Year}
			// Stack escapes to heap because of return
			// https://go.dev/blog/escape-analysis
			return &car
		}
	}
	return nil // Return nil if not found
}

func main() {
	myCars := []Car{
		{"Porsche", "959", 1984},
		{"Honda", "Civic", 2022},
	}
	printCar(myCars)
	findCar(myCars, "Civic")
}
