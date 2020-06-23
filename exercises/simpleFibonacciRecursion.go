package main

import (
	"fmt"
	"time"
)

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	var fibIndex int64
	/*
			Assumptions:
		  - sequence index starts from index `0` for sequence `0, 1, 1, 2, 3 ...`"
		  - index `0` means we have not taken any steps"
		  - going from `0` to `1` is one step"
		  - thus `value at index 1` correlates to `number after 1 step`"
		  - likewise `value at index 50` correlates to `number after 50 steps`
	*/
	fmt.Println(".............. Fibonacci (recursion) ...............\n")
	fmt.Printf("Enter desired Fibonacci index: ")
	fmt.Scanf("%v", &fibIndex)
	if fibIndex < 0 {
		fmt.Println("ERROR - Please enter a non-negative number: ")
		fmt.Scanf("%v", &fibIndex)
	}

	// start time counter
	fibStartTime := time.Now()
	computedValue := computeFibSeq(fibIndex)
	fibStopTime := time.Now()
	fmt.Printf("\nFibonacci (recursion) value (after %v steps): %v\n", fibIndex, computedValue)
	fmt.Printf("runtime: %v seconds \n", fibStopTime.Sub(fibStartTime).Seconds())
}

func computeFibSeq(fibIndex int64) int64 {
	if fibIndex == 0 || fibIndex == 1 { // this handles both Fibonacci index 0 and 1
		return fibIndex
	}
	return computeFibSeq(fibIndex-1) + computeFibSeq(fibIndex-2)
}
