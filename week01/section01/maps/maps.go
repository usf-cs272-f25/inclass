package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 42
	v := m["foo"]
	fmt.Println(v)

	k := "bar"
	val, ok := m[k]
	if !ok {
		// val is not valid
		fmt.Println("not found: ", k)
	} else {
		// val is valid
		fmt.Println("val: ", val)
	}


	type Person struct {
		Name string
		Height int
	}

	p := Person{"Phil", 70}
	m2 := make(map[int]Person)
}
