package main

import (
	"fmt"
	"reflect"
)

// ============ DEMONSTRATION OF KEY DIFFERENCES ============

// 1. RUNTIME TYPE INFORMATION (Go has it, Java doesn't)
func ShowRuntimeType[T any](value T) {
	fmt.Printf("Runtime type of %v is: %T\n", value, value)
	fmt.Printf("Reflection type: %v\n", reflect.TypeOf(value))
}

// 2. NO WILDCARDS IN GO
// In Java you could do: List<? extends Number>
// In Go, you must be explicit:

type Number interface {
	~int | ~float64
}

func SumNumbers[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// 3. MONOMORPHIZATION DEMONSTRATION
// Go generates separate code for each type

func Identity[T any](v T) T {
	// In Go, this generates different machine code for Identity[int] and Identity[string]
	// In Java, there would be one version that works with Object
	return v
}

// 4. NO BOXING IN GO
func ProcessInts(nums []int) int {
	// Direct int manipulation, no Integer wrapper objects
	sum := 0
	for _, n := range nums {
		sum += n // No boxing/unboxing
	}
	return sum
}

// 5. TYPE SETS (Go's unique approach)
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 6. INVARIANCE IN GO (no covariance/contravariance)
type Animal struct{ Name string }
type Dog struct {
	Animal
	Breed string
}

// Can't treat []Dog as []Animal in Go
func FeedAnimals(animals []Animal) {
	for _, a := range animals {
		fmt.Printf("Feeding %s\n", a.Name)
	}
}

// 7. METHOD CONSTRAINTS (Go can constrain on methods)
type Stringer interface {
	String() string
}

func PrintAll[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

// 8. ZERO VALUES WORK NATURALLY
func GetZeroValue[T any]() T {
	var zero T // Natural zero value
	return zero
}

// Custom type for demonstration
type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func main() {
	fmt.Println("=== RUNTIME TYPE INFORMATION ===")
	ShowRuntimeType(42)
	ShowRuntimeType("hello")
	ShowRuntimeType([]int{1, 2, 3})
	
	fmt.Println("\n=== NO WILDCARDS (EXPLICIT CONSTRAINTS) ===")
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Printf("Sum of ints: %d\n", SumNumbers(ints))
	fmt.Printf("Sum of floats: %.1f\n", SumNumbers(floats))
	
	fmt.Println("\n=== MONOMORPHIZATION ===")
	// These create different compiled functions
	intResult := Identity(42)
	strResult := Identity("hello")
	fmt.Printf("Identity(42) = %d\n", intResult)
	fmt.Printf("Identity(\"hello\") = %s\n", strResult)
	
	fmt.Println("\n=== NO BOXING ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum without boxing: %d\n", ProcessInts(nums))
	
	fmt.Println("\n=== TYPE SETS ===")
	fmt.Printf("Min(10, 20) = %d\n", Min(10, 20))
	fmt.Printf("Min(3.14, 2.71) = %.2f\n", Min(3.14, 2.71))
	fmt.Printf("Min(\"apple\", \"banana\") = %s\n", Min("apple", "banana"))
	
	fmt.Println("\n=== INVARIANCE ===")
	animals := []Animal{{Name: "Cat"}, {Name: "Bird"}}
	FeedAnimals(animals)
	
	// This won't compile (invariance):
	// dogs := []Dog{{Animal: Animal{Name: "Rover"}, Breed: "Golden"}}
	// FeedAnimals(dogs) // ERROR: cannot use []Dog as []Animal
	
	fmt.Println("\n=== METHOD CONSTRAINTS ===")
	people := []Person{{Name: "Alice"}, {Name: "Bob"}}
	PrintAll(people)
	
	fmt.Println("\n=== ZERO VALUES ===")
	zeroInt := GetZeroValue[int]()
	zeroString := GetZeroValue[string]()
	zeroPerson := GetZeroValue[Person]()
	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero string: '%s'\n", zeroString)
	fmt.Printf("Zero Person: %+v\n", zeroPerson)
}

/*
KEY DIFFERENCES FROM JAVA:

1. TYPE ERASURE vs MONOMORPHIZATION
   - Java: One version working with Object
   - Go: Separate compiled code per type

2. WILDCARDS vs EXPLICIT CONSTRAINTS
   - Java: List<? extends Number>
   - Go: func F[T Number](list []T)

3. BOXING vs DIRECT VALUES
   - Java: List<Integer> with boxing overhead
   - Go: []int with direct values

4. VARIANCE
   - Java: Covariance/contravariance with wildcards
   - Go: Always invariant

5. RUNTIME TYPE INFO
   - Java: Lost due to erasure
   - Go: Preserved and accessible

6. CONSTRAINTS
   - Java: Based on inheritance (extends/implements)
   - Go: Based on type sets and method sets
*/