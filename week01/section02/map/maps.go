package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 1

	k := "bar"
	val, ok := m[k]
	if ok {
		// val is valid
		fmt.Println("val: ", val)
	} else {
		// val is not valid
		fmt.Println(val)
		fmt.Printf("key %s is not in the map\n", k)
	}
}
