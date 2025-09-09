package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Make string
	Year uint
}

type Cars []Car

func (c Cars) Less(i, j int) bool {
	return c[i].Year < c[j].Year
}

func (c Cars) Swap(i, j int) {
	// tmp := c[i]
	// c[i] = c[j]
	// c[j] = tmp

	c[i], c[j] = c[j], c[i]
}

func (c Cars) Len() int {
	return len(c)
}

func main() {
	garage := []Car{
		{"Honda", 1989},
		{"Porsche", 2013},
		{"Fiat", 1967},
	}

	fmt.Println("Unsorted: ", garage)
	sort.Sort(Cars(garage))
	fmt.Println("Sorted: ", garage)
}
