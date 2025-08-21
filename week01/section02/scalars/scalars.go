package main

import "fmt"

func main() {
	i := 3
	for j := range i {
		isZero := j == 0
		fmt.Println("isZero: ", isZero)
		fmt.Println("j: ", j)
	}
}
