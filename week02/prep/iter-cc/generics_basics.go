package main

import (
	"fmt"
)

// ============ BASIC GENERIC FUNCTIONS ============

// Generic function with single type parameter
func PrintSlice[T any](s []T) {
	fmt.Printf("Slice of %T: ", s)
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// Generic function that returns a value
func First[T any](slice []T) (T, bool) {
	var zero T // zero value of type T
	if len(slice) == 0 {
		return zero, false
	}
	return slice[0], true
}

// Multiple type parameters
func MakePair[K any, V any](key K, value V) (K, V) {
	return key, value
}

// ============ TYPE CONSTRAINTS ============

// Constraint: only types that can be ordered
func Max[T ~int | ~float64 | ~string](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Using the comparable constraint (types that support == and !=)
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Custom constraint interface
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// ============ GENERIC TYPES ============

// Generic struct
type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func (b *Box[T]) Set(v T) {
	b.value = v
}

// Generic struct with multiple type parameters
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Generic slice type
type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(*s) == 0 {
		return zero, false
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, true
}

// ============ TYPE INFERENCE ============

func Double[T Number](x T) T {
	return x * 2
}

// ============ ADVANCED: METHOD TYPE PARAMETERS ============

type DataProcessor struct {
	name string
}

// Methods can't have type parameters in Go (as of 1.23)
// But the receiver type can be generic
type List[T any] struct {
	items []T
}

func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// Map transforms each element (Note: methods can't have type params, 
// so we make it a function instead)
func MapList[T any, U any](l *List[T], transform func(T) U) *List[U] {
	result := &List[U]{items: make([]U, 0, len(l.items))}
	for _, item := range l.items {
		result.Add(transform(item))
	}
	return result
}

func main() {
	fmt.Println("========== BASIC GENERICS ==========")
	
	// Type inference - Go figures out T is int
	PrintSlice([]int{1, 2, 3})
	
	// Explicit type specification
	PrintSlice[string]([]string{"hello", "world"})
	
	// Works with custom types too
	type Person struct{ Name string }
	PrintSlice([]Person{{Name: "Alice"}, {Name: "Bob"}})

	fmt.Println("\n========== RETURN VALUES ==========")
	
	nums := []int{42, 17, 99}
	if first, ok := First(nums); ok {
		fmt.Printf("First number: %d\n", first)
	}
	
	empty := []string{}
	if first, ok := First(empty); ok {
		fmt.Printf("First string: %s\n", first)
	} else {
		fmt.Println("Empty slice, no first element")
	}

	fmt.Println("\n========== MULTIPLE TYPE PARAMETERS ==========")
	
	key, val := MakePair("age", 30)
	fmt.Printf("Pair: %s = %d\n", key, val)
	
	k2, v2 := MakePair(42, "answer")
	fmt.Printf("Pair: %d = %s\n", k2, v2)

	fmt.Println("\n========== TYPE CONSTRAINTS ==========")
	
	fmt.Printf("Max of 10 and 20: %d\n", Max(10, 20))
	fmt.Printf("Max of 'apple' and 'banana': %s\n", Max("apple", "banana"))
	
	// Custom type with underlying int type
	type Age int
	fmt.Printf("Max age: %d\n", Max(Age(25), Age(30)))
	
	// Contains with comparable constraint
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Contains 'Bob': %v\n", Contains(names, "Bob"))
	fmt.Printf("Contains 'Dave': %v\n", Contains(names, "Dave"))

	fmt.Println("\n========== CUSTOM CONSTRAINTS ==========")
	
	ints := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of ints: %d\n", Sum(ints))
	
	floats := []float64{1.5, 2.5, 3.5}
	fmt.Printf("Sum of floats: %.1f\n", Sum(floats))

	fmt.Println("\n========== GENERIC TYPES ==========")
	
	// Box with int
	intBox := Box[int]{value: 42}
	fmt.Printf("Box contains: %d\n", intBox.Get())
	intBox.Set(100)
	fmt.Printf("Box now contains: %d\n", intBox.Get())
	
	// Box with string
	strBox := Box[string]{value: "hello"}
	fmt.Printf("Box contains: %s\n", strBox.Get())
	
	// Pair
	p := Pair[string, int]{Key: "score", Value: 95}
	fmt.Printf("Pair: %s = %d\n", p.Key, p.Value)

	fmt.Println("\n========== GENERIC STACK ==========")
	
	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	
	for {
		if val, ok := stack.Pop(); ok {
			fmt.Printf("Popped: %s\n", val)
		} else {
			break
		}
	}

	fmt.Println("\n========== TYPE INFERENCE ==========")
	
	// Go infers T is int
	fmt.Printf("Double 21: %d\n", Double(21))
	
	// Go infers T is float64
	fmt.Printf("Double 3.14: %.2f\n", Double(3.14))
	
	// Explicit type when needed
	var x int32 = 10
	fmt.Printf("Double int32: %d\n", Double[int32](x))

	fmt.Println("\n========== GENERIC LIST WITH MAP ==========")
	
	list := &List[int]{items: []int{1, 2, 3, 4, 5}}
	
	// Map to strings
	stringList := MapList(list, func(n int) string {
		return fmt.Sprintf("num_%d", n)
	})
	
	fmt.Printf("Original: %v\n", list.items)
	fmt.Printf("Mapped: %v\n", stringList.items)
	
	// Map to squares
	squaresList := MapList(list, func(n int) int {
		return n * n
	})
	fmt.Printf("Squares: %v\n", squaresList.items)
}

/*
KEY CONCEPTS:

1. TYPE PARAMETERS: [T any] where T is the parameter name
2. TYPE CONSTRAINTS: [T comparable], [T ~int | ~string], custom interfaces
3. TYPE INFERENCE: Go often figures out types automatically
4. GENERIC TYPES: Structs and type aliases can be generic
5. TILDE (~): Means "underlying type", so ~int matches int and type MyInt int
6. COMPARABLE: Built-in constraint for types supporting == and !=
7. ANY: Constraint allowing any type (same as interface{})

LIMITATIONS:
- Methods can't have their own type parameters (yet)
- Type parameters can't be used with operators except as constrained
- No partial type inference
*/