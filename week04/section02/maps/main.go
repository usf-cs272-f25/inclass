package main

import "fmt"

func NewMap() map[string]any {
	m := make(map[string]any)
	m["foo"] = 1
	m["bar"] = "cool"
	return m
}

func main() {
	// Use make() to initialize internal data structures, e.g. hash buckets
	m1 := make(map[string]int)
	fmt.Printf("m1: %v\n", m1)

	m1["foo"] = 42
	// val, exists are only in scope in the if/else blocks
	if val, exists := m1["foo"]; !exists {
		fmt.Println("foo does not exist in m1")
	} else {
		fmt.Printf("foo: %d\n", val)
	}

	// It's equivalent to use a map literal, perhaps useful for test cases
	m2 := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println("m2: ", m2)

	// For maps, range returns the key and value
	for k, v := range m2 {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

	m3 := NewMap()
	fmt.Printf("m3: %v\n", m3)
}
