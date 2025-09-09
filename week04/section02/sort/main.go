package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Height uint
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Height > p[j].Height
}

func (p People) Swap(i, j int) {
	// tmp := p[i]
	// p[i] = p[j]
	// p[j] = tmp

	p[i], p[j] = p[j], p[i]
}

func main() {
	people := People{
		{"Phil", 70},
		{"Steph", 78},
		{"DannyD", 60},
	}

	fmt.Printf("Unsorted: %v\n", people)
	sort.Sort(people)
	fmt.Printf("Sorted: %v\n", people)
}
