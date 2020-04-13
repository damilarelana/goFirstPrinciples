package main

import (
	"fmt"
	"math/rand"
	"time"
)

// arrayShuffler()
func arrayShuffler(array []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // required to help ensure that rand.Shuffle generates random pseudonumbers for shuffling
	for len(array) > 0 {                                 // starting at the end of the slice seems to work well
		arrayLength := len(array)
		randomIndex := r.Intn(arrayLength)                                                  // choose random index to shuffle with end of the slice at index `arrayLength - 1`
		array[arrayLength-1], array[randomIndex] = array[randomIndex], array[arrayLength-1] // shuffle
		array = array[:arrayLength-1]                                                       // reduce the slice elements yet to be randomised i.e. it immediately excludes the previous array[arrayLength-1]
		// note that `for len(array) > 0` would eventually terminate since we keep creating smaller and smaller slices
	}
}

func main() {
	array := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	arrayLength := len(array)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("A golang implementation of an Array Shuffler :")
	fmt.Println(" - using random generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements", arrayLength)
	fmt.Println("\n")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	/*
	 call bubblesort to sort the initial array
	 initialArray... is used to pass the slice initialArray as variadic parameters
	*/
	// start time counter
	fmt.Printf("\nOriginal Array: %v \n", array)
	startTime := time.Now()
	arrayShuffler(array) // array is shuffled
	timeNow := time.Now()
	fmt.Printf("Shuffled Array: %v \n", array)
	fmt.Printf("shuffling runtime duration: %v \n", timeNow.Sub(startTime))
}
