package main

import "fmt"

func main() {
	s := "hi 👋 👍x"
	for idx, r := range s {
		fmt.Printf("idx %d r %q\n", idx, r)
	}

	s2 := "hi 👋 👍"
	s2 += "x"
	fmt.Println(s == s2)
}
