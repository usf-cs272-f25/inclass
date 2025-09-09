package main

import "fmt"

type Frequency map[string]uint
type InvIdx map[string]Frequency

type VisitedMap map[string]struct{}

func NewInvIdx() InvIdx {
	return make(InvIdx)
}

func main() {
	// Declaring a map using make might be useful for inv idx
	m1 := make(map[string]int)

	// Test whether a key exists using ok
	if _, ok := m1["foo"]; !ok {
		// the short declared variables are only in scope inside the if
		fmt.Println("m1 does not contain 'foo'")
	}

	// Declaring a map using a literal might be useful for test cases
	m2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	// For maps, range returns key and value
	// If you use only one identifier, you get the value
	// If you write k, _ the compiler will suggest you simplify
	for k, v := range m2 {
		fmt.Println("key: ", k)
		fmt.Println("val: ", v)
	}

	var a any
	a = 1
	a = "foo"
	fmt.Println("a: ", a)

	// Maps can be a return value, like any other data type
	ii := NewInvIdx()
	fmt.Println("ii: ", ii)

	// These two are equivalent. The first one is a map literal, initialized to empty
	m3 := map[string]struct{}{}
	m4 := make(map[string]struct{})

	m5 := VisitedMap{}
	m6 := make(VisitedMap)
}
