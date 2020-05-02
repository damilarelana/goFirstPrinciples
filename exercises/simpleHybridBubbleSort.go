package main

import (
	"fmt"
	"math/rand"
	"time"
)

// hybridBubbleSort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - while still using a swap flag AND break point
func hybridBubbleSort(dynamicArray []int) []int { // variadic arguments are used here
	outerCount := 0
	inputArrayLength := len(dynamicArray)
OuterForLoop:
	for outerCount < len(dynamicArray) {
		/* initial flag handles: sorted input, sorting completion, and bubbling loop */
		swapflag := false
		innerCount := 0
		for innerCount < (inputArrayLength - 1) {
			if dynamicArray[innerCount] > dynamicArray[innerCount+1] {
				dynamicArray[innerCount], dynamicArray[innerCount+1] = dynamicArray[innerCount+1], dynamicArray[innerCount]
				swapflag = true
			}
			innerCount++
		}
		/* exiting from loop when already sorted input and sorting completion */
		if !swapflag {
			break OuterForLoop
		}
		outerCount++
		inputArrayLength-- // decrement array length before next iteration, since previous largest value does not need to be involved in next iterations
	}

	return dynamicArray
}

// createRandomList()
func createRandomList(minValue int, maxValue int, stepValue int) (array []int) {

	// initialize the array
	totalIncrements := (maxValue - minValue) / stepValue
	array = make([]int, totalIncrements+1)

	// populate the array elements
	for i, _ := range array {
		if i == 0 {
			array[i] = minValue
		} else {
			array[i] = array[i-1] + stepValue
		}
	}

	// return the randomized array
	return array
}

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
	// initialArray := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	initialArray := createRandomList(0, 1247635, 96)
	arrayShuffler(initialArray) // shuffler the elements of the array
	arrayLength := len(initialArray)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("A golang implementation of a Hybrid Bubble Sort algorithm :")
	fmt.Println("  - using randomly generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements\n", arrayLength)
	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// start time counter
	startTime := time.Now()
	var sortedArray = hybridBubbleSort(initialArray) // the array is passed as set of variadic arguments
	timeNow := time.Now()
	fmt.Printf("\nHybrid Bubble Sort gives [first 15 elements as]: %v \n", sortedArray[:15])
	fmt.Printf("runtime duration: %v seconds \n", timeNow.Sub(startTime).Seconds())
	fmt.Printf("largest number is : %d \n", sortedArray[arrayLength-1])
	fmt.Printf("smallest number is : %d \n\n", sortedArray[0])
}
