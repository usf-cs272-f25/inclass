package main

import "fmt"

// Go's "type parameters" (here T) provide a generic
// way to write logic, when the logic could be used
// with any data type. Similar to C++ template or Rust trait

func Sum[T int | float32](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3}
	sum := Sum(arr)
	fmt.Println(sum)

	arr2 := []float32{0.1, 0.2, 0.3}
	sum2 := Sum(arr2)
	fmt.Println(sum2)
}
