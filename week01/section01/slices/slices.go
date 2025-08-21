package main

import "fmt"

func main() {
	s := []int{}
	s = append(s, 4)
	s = append(s, 5)

	fmt.Println(s)

	for _, val := range s {
		fmt.Println(val)
	}

	s2 := []string{}
	s2 = append(s2, "foo")
	s2 = append(s2, "bar")
	fmt.Println(s2)
}
