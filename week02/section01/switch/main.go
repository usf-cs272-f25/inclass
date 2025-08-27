package main

import "fmt"

func foo(s string) {
	// No break needed for exclusive cases
	// Use "fallthrough" to get "OR"-style semantics
	// Switch is preferred to cascading if/else blocks
	switch s {
		case "foo":
			fmt.Println("foo")
		case "bar":
			fmt.Println("bar")
		default:
			// not needed for strings
			// would be required for integers to cover all cases
	}
}

func bar(i int) {
	switch {
		case i < 10:
			// something
		case i > 20:
			// something else
	}
}

func main() {
	foo("foo")
}
