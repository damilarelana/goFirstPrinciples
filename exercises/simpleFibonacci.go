package main

import (
	"fmt"
	"time"
)

func main() {
	var fibIndex int64
	/*
		Assumptions:
		- Fibonacci sequence starts from index `0` when counting the sequence `0, 1, 1, 2, 3 ...`
		- going from `0` to `1` is one step i.e. `value at index 1` correlates to `number after 1 step`
	*/
	fmt.Println(".............. Fibonacci ...............\n")
	fmt.Println("Assume:")
	fmt.Println("  - sequence index starts from index `0` for sequence `0, 1, 1, 2, 3 ...`")
	fmt.Println("  - index `0` means we have not taken any steps")
	fmt.Println("  - going from `0` to `1` is one step")
	fmt.Println("  - thus `value at index 1` correlates to `number after 1 step`")
	fmt.Println("  - likewise `value at index 50` correlates to `number after 50 steps`\n")
	fmt.Println("......................................\n")
	fmt.Println("Enter index (i.e. step) of desired Fibonacci number:")
	fmt.Scanf("%v", &fibIndex)
	if fibIndex < 0 {
		fmt.Println("ERROR - Please enter a non-negative number: ")
		fmt.Scanf("%v", &fibIndex)
	}

	// start time counter
	fibStartTime := time.Now()
	computedValue := computeFibSeq(fibIndex)
	fibStopTime := time.Now()
	fmt.Printf("Fibonacci value (after %v steps): %v\n", fibIndex, computedValue)
	fmt.Printf("runtime: %v seconds \n", fibStopTime.Sub(fibStartTime).Seconds())
}

func computeFibSeq(fibIndex int64) int64 {
	if fibIndex == 0 || fibIndex == 1 { // this handles both Fibonacci index 0 and 1
		return fibIndex
	}
	return computeFibSeq(fibIndex-1) + computeFibSeq(fibIndex-2)
}
