package main

import (
	"fmt"
	"sort"
	"time"
)

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

func main() {
	generationStartTime := time.Now()
	array := createRandomList(0, 1247635, 96)
	generationTimeNow := time.Now()
	arrayLength := len(array)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("A golang implementation of an Array Generator :")
	fmt.Println(" - using random generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements", arrayLength)
	fmt.Println("\n")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// Sorting
	sortStartTime := time.Now()
	sort.Ints(array) // this sorts in place, without returning anything
	sortTimeNow := time.Now()

	// Printing
	fmt.Printf("generation runtime duration: %v \n", generationTimeNow.Sub(generationStartTime))
	fmt.Printf("golang's sort package gives [first 15 elements as]: %v \n", array[:15])
	fmt.Printf("sort runtime duration: %v \n", sortTimeNow.Sub(sortStartTime))
	fmt.Printf("largest number is : %d \n", array[arrayLength-1])
	fmt.Printf("smallest number is : %d \n\n", array[0])
}
