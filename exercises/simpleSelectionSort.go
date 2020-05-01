package main

import (
	"fmt"
	"math/rand"
	"time"
)

// selectionSort()
func selectionSort(dynamicArray []int) []int {
	inputArrayLength := len(dynamicArray)
	if inputArrayLength == 1 {
		return dynamicArray
	}
	outerCount := 0
	minElement := outerCount // assume first index "0" is temporary minimum (changes with each pass)
OuterForLoop:
	for outerCount < inputArrayLength {
		innerCount := outerCount + 1        // make (or reset) innerCount to current "outerCount + 1"
		for innerCount < inputArrayLength { // scanning by looping over all remaining items to test new minimum
			if dynamicArray[innerCount] < dynamicArray[minElement] { // if any of the items if less than current minimum
				minElement = innerCount // swaps out the index of the old with the new i.e. create new temporary minimum for remaining unsorted set
			}
			innerCount++
		} // increase inner counter i.e. reducing unsorted list of items
		dynamicArray[outerCount], dynamicArray[minElement] = dynamicArray[minElement], dynamicArray[outerCount] // confirm new minimum by swapping [temporary outerCount index with new minimum's index]
		outerCount++                                                                                            // increase outer counter i.e. expanding the sorted set
		minElement = outerCount                                                                                 // reset new temporary minimum index e.g. if initial was index `0`, it would now be `1`
		if outerCount == (inputArrayLength - 1) {                                                               // i.e. only one unsorted element remains, break outer loop
			break OuterForLoop
		}
		// note that we CANNOT use the optimization (loopRange -= 1) since we are shifting values/index around
		// as such the last value after every iteration can still need to be touched
		// this is one difference with BubbleSort() where the last index can be removed from dataset after every loop
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
	fmt.Println("A golang implementation of Selection Sort algorithm :")
	fmt.Println("  - using randomly generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements\n", arrayLength)
	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	/*
	 call bubblesort to sort the initial array
	 initialArray... is used to pass the slice initialArray as variadic parameters
	*/
	// start time counter
	startTime := time.Now()
	var sortedArray = selectionSort(initialArray) // the array is passed as set of variadic arguments
	timeNow := time.Now()
	fmt.Printf("\nSelection Sort gives [first 15 elements as]: %v \n", sortedArray[:15])
	fmt.Printf("runtime duration: %v seconds \n", timeNow.Sub(startTime).Seconds())
	fmt.Printf("largest number is : %d \n", sortedArray[arrayLength-1])
	fmt.Printf("smallest number is : %d \n\n", sortedArray[0])
}
