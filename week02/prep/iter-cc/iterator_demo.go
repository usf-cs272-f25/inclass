package main

import (
	"fmt"
	"iter"
)

// IntRange returns an iterator that yields integers from start to end (exclusive)
func IntRange(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Squares returns an iterator that yields squares of numbers up to n
func Squares(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i * i) {
				return
			}
		}
	}
}

// Fibonacci returns an iterator that yields the first n Fibonacci numbers
func Fibonacci(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	fmt.Println("Range from 1 to 10:")
	for v := range IntRange(1, 10) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\nSquares up to 5:")
	for v := range Squares(5) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\nFirst 10 Fibonacci numbers:")
	for v := range Fibonacci(10) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\nUsing break to stop early:")
	for v := range IntRange(1, 100) {
		if v > 5 {
			break
		}
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}