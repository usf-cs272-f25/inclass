package main

import (
	"fmt"
	"iter"
)

// iter.Seq[T] is a generic type alias defined in the iter package
// It's essentially: type Seq[T any] func(yield func(T) bool)

// Generic iterator that works with any type
func Repeat[T any](value T, times int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < times; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// Iterator for pairs of values (using iter.Seq2)
func Enumerate[T any](slice []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

// Custom generic iterator type (showing what iter.Seq looks like)
type MySeq[T any] func(yield func(T) bool)

// Using our custom type
func CountBy[T ~int](start, step T) MySeq[T] {
	return func(yield func(T) bool) {
		current := start
		for {
			if !yield(current) {
				return
			}
			current += step
		}
	}
}

// Pre-generics approach (what we had to do before Go 1.18)
// We'd need separate functions for each type:
func RepeatInt(value int, times int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < times; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

func RepeatString(value string, times int) func(func(string) bool) {
	return func(yield func(string) bool) {
		for i := 0; i < times; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

func main() {
	fmt.Println("=== Generic Repeat with int ===")
	for v := range Repeat(42, 3) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\n=== Generic Repeat with string ===")
	for v := range Repeat("hello", 3) {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	fmt.Println("\n=== Enumerate (iter.Seq2 with two values) ===")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range Enumerate(fruits) {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	fmt.Println("\n=== Custom generic iterator ===")
	count := 0
	for v := range CountBy(10, 5) {
		fmt.Printf("%d ", v)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println()

	fmt.Println("\n=== Type constraints with ~int ===")
	type MyInt int
	for v := range CountBy(MyInt(100), MyInt(10)) {
		fmt.Printf("%d ", v)
		if v > 130 {
			break
		}
	}
	fmt.Println()
}

/*
GENERICS IN ITERATORS:

1. iter.Seq[T] - Generic iterator for single values
   - T can be any type (int, string, struct, etc.)
   - Defined as: type Seq[T any] func(yield func(T) bool)

2. iter.Seq2[K, V] - Generic iterator for pairs
   - Used for key-value pairs or indexed iterations
   - Defined as: type Seq2[K, V any] func(yield func(K, V) bool)

3. Type parameters in square brackets [T any]:
   - T is the type parameter name
   - 'any' is the constraint (allows any type)
   - Can use other constraints like [T ~int] for underlying int types

4. Why generics are essential:
   - Without generics, we'd need separate iterator types for each data type
   - Generics allow one iterator implementation to work with any type
   - Type safety at compile time

5. The yield function is also generic:
   - In iter.Seq[T], yield has signature func(T) bool
   - The T matches the iterator's type parameter
*/