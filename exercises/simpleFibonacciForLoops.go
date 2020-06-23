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
	fmt.Println(".............. Fibonacci (for-loop) ...............\n")
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
	fmt.Printf("\nFibonacci (for-loop) value (after %v steps): %v\n", fibIndex, computedValue)
	fmt.Printf("runtime: %v seconds \n", fibStopTime.Sub(fibStartTime).Seconds())
}

func computeFibSeq(fibIndex int64) int64 {
	/*
		create a slice to use for caching results
			- fibIndex+1 is the length of the slice i.e. helps to accommodate that value at index `2` means the length is `2+1` when counting from `0` index
			- fibIndex+1 is the capacity of slice i.e. used to improve code efficiency at the low level to avoid memory reallocation
	*/
	fibCache := make([]int64, fibIndex+1, fibIndex+2)

	// this handles both Fibonacci index 0 and 1
	if fibIndex == 0 || fibIndex == 1 {
		return fibIndex
	}

	// if fibIndex is >= 2, then start caching
	fibCache[0] = 0
	fibCache[1] = 1

	/*
	 The array assignments already imply an inherent caching
	 Since each array element is calculated in 1-pass - it also means that `re-use of cache` would be not possible
	*/
	for cIndex := int64(2); cIndex <= fibIndex; cIndex++ { // initialize for loop index for other fibonacci values
		fibCache[cIndex] = fibCache[cIndex-1] + fibCache[cIndex-2]
	}
	return fibCache[fibIndex]
}
