package main

import "fmt"

func main() {
	s := "hi 😀 é"
	fmt.Println("s: ", s)

	s2 := "hi 😀 é x"
	fmt.Println(s == s2)

	for idx, r := range s {
		fmt.Printf("%d: %q\n", idx, r)
	}
}
