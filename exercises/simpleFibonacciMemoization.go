package main

import (
	"fmt"
	"time"
)

func main() {
	var fibIndex int64
	/*
		Assumptions:
		  - sequence index starts from index `0` for sequence `0, 1, 1, 2, 3 ...`"
		  - index `0` means we have not taken any steps"
		  - going from `0` to `1` is one step"
		  - thus `value at index 1` correlates to `number after 1 step`"
			- likewise `value at index 50` correlates to `number after 50 steps`
		Uses memoization to speed up the number computation
	*/
	fmt.Println(".............. Fibonacci (via recursion + memoization) ...............\n")
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
	fmt.Printf("\nFibonacci (via recursion + memoization) value (after %v steps): %v\n", fibIndex, computedValue)
	fmt.Printf("runtime: %v seconds \n", fibStopTime.Sub(fibStartTime).Seconds())
}

func computeFibSeq(fibIndex int64) int64 {
	fibCacheMap := make(map[int64]int64)      // create a hashmap to use for caching results
	fibCacheMap[0] = 0                        // handles and caches Fibonacci number at index 0, in the initialized hashmap
	fibCacheMap[1] = 1                        // handles and caches Fibonacci number at index 1, in the initialized hashmap
	return cacheMapper(fibIndex, fibCacheMap) // returns the computed value after memoization is applied in the middleware
}

func cacheMapper(fibIndex int64, fibCacheMap map[int64]int64) int64 {
	/*
		check if result already exists in fibCacheMap
		- if it exists then ok is `true`
		- then proceed to return the cached value
	*/
	var cachedMapValue int64
	cachedMapValue, ok := fibCacheMap[fibIndex]
	if ok { // returns what is in cache
		return cachedMapValue
	}

	/*
		Since it does not already exists in fibCacheMap
		- calculate and cache it in fibCacheMap BY recursively calling cacheMapper
		- then return the cachedMapValue
	*/
	fibCacheMap[fibIndex] = cacheMapper(fibIndex-1, fibCacheMap) + cacheMapper(fibIndex-2, fibCacheMap)
	return fibCacheMap[fibIndex]
}
