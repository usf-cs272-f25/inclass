package main

import "fmt"

func foo(s string) {
	// Switch is more concise than cascading if/else, so preferred
	switch s {
		case "foo":
			fmt.Println("foo")
		case "bar":
			fmt.Println("bar")
		default:
			fmt.Println("default")
	}
}

func bar(i int) {
	switch i {
		case 0:
			fmt.Println("0")
		default:
			fmt.Println("default int")
	}
}

func baz(int j) {
	switch {
		case j < 10:
			fmt.Println("< 10")
		case j > 100:
			fmt.Println("> 100")
	}
}
func main() {
	foo("foo")
}
