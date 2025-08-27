package main

import (
	"fmt"
	"iter"
)

// Let's trace through the yield pattern step by step
func SimpleIterator() iter.Seq[int] {
	// This function returns an iterator function
	return func(yield func(int) bool) {
		// The 'yield' function is PROVIDED BY Go's range loop
		// We CALL yield with each value we want to produce
		
		fmt.Println("Iterator: Starting...")
		
		// yield(1) sends 1 to the range loop
		// It returns true if loop continues, false if loop breaks
		if !yield(1) {
			fmt.Println("Iterator: Consumer stopped at 1")
			return
		}
		fmt.Println("Iterator: Yielded 1, continuing...")
		
		if !yield(2) {
			fmt.Println("Iterator: Consumer stopped at 2")
			return
		}
		fmt.Println("Iterator: Yielded 2, continuing...")
		
		if !yield(3) {
			fmt.Println("Iterator: Consumer stopped at 3")
			return
		}
		fmt.Println("Iterator: Yielded 3, continuing...")
		
		fmt.Println("Iterator: Finished all values")
	}
}

// Custom implementation showing what Go does behind the scenes
func ManualIteration() {
	fmt.Println("\n=== Manual iteration (what Go does internally) ===")
	
	iterator := SimpleIterator()
	
	// Go creates this yield function for us in a range loop
	iterator(func(value int) bool {
		fmt.Printf("Consumer: Received %d\n", value)
		
		// Simulate breaking after value 2
		if value == 2 {
			fmt.Println("Consumer: Deciding to stop")
			return false // Tell iterator to stop
		}
		
		return true // Tell iterator to continue
	})
}

func main() {
	fmt.Println("=== Normal range loop ===")
	for v := range SimpleIterator() {
		fmt.Printf("Consumer: Got %d\n", v)
		if v == 2 {
			fmt.Println("Consumer: Breaking out")
			break
		}
	}
	
	fmt.Println("\n=== Range loop that consumes all ===")
	for v := range SimpleIterator() {
		fmt.Printf("Consumer: Got %d\n", v)
	}
	
	// Show manual iteration to understand the mechanism
	ManualIteration()
}

/*
WHO DOES WHAT:

1. ITERATOR (our code):
   - Receives a 'yield' function as parameter
   - Calls yield(value) for each value to produce
   - Checks yield's return value (true=continue, false=stop)
   - Returns early if yield returns false

2. GO RUNTIME (range loop):
   - Creates the yield function
   - Passes yield to our iterator
   - Executes loop body when yield is called
   - Returns true from yield to continue iteration
   - Returns false from yield when loop breaks

3. CONSUMER (the for loop body):
   - Receives values through the range variable
   - Can break to stop iteration early

FLOW:
1. 'for v := range Iterator()' starts
2. Go creates a yield function
3. Go calls Iterator(yield)
4. Iterator calls yield(value1)
5. yield runs loop body with value1
6. If loop continues, yield returns true
7. Iterator calls yield(value2)
8. If loop breaks, yield returns false
9. Iterator sees false and returns early
*/