package main

import "fmt"

type Person struct {
	Name   string
	Height int
}

func main() {
	p := Person{}
	p.Name = "Phil"
	p.Height = 70

	p2 := Person{"Steph", 78}

	fmt.Println(p)
	fmt.Println(p2)
}
