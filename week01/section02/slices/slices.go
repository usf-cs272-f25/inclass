package main

import "fmt"

func main() {
	// var sl2 []int
	// sl2 = append(sl2, 4)

	sl := []int{}
	sl = append(sl, 4)
	sl = append(sl, 5)

	for idx, val := range sl {
		fmt.Printf("idx: %d val: %d\n", idx, val)
	}
}
