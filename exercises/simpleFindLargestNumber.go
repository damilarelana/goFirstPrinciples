package main

import "fmt"

func main() {
	initialArray := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("Given array: ", initialArray)

	/*
	 call bubblesort to sort the initial array
	 initialArray... is used to pass the slice initialArray as variadic parameters
	*/
	var sortedArray = bubblesort(initialArray...) // the array is passed as set of variadic arguments

	/* pick the largest number */
	fmt.Println("BubbleSort gives largest num as:", sortedArray[len(sortedArray)-1])
}

// variadic parameters being used for BubbleSort
func bubblesort(dynamicArray ...int) []int { // variadic arguments are used here
	outerCount := 0
OuterForLoop:
	for outerCount < len(dynamicArray) {
		/* initial flag handles: sorted input, sorting completion, and bubbling loop */
		swapflag := false
		innerCount := 0
		for innerCount < (len(dynamicArray) - 1) {
			if dynamicArray[innerCount] > dynamicArray[innerCount+1] {
				temp := dynamicArray[innerCount+1]
				dynamicArray[innerCount+1] = dynamicArray[innerCount]
				dynamicArray[innerCount] = temp
				swapflag = true
			}
			innerCount++
		}
		/* exiting from loop when already sorted input and sorting completion */
		if !swapflag {
			break OuterForLoop
		}
		outerCount++
	}
	return dynamicArray
}
