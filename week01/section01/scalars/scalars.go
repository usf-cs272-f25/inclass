package main

import "fmt"

func main() {
	i := 3

	for j := range i {
		fmt.Println("j: ", j)
		isZero := j == 0
		fmt.Println("isZero: ", isZero)
	}
}
