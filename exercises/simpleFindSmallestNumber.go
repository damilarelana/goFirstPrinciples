package main

import "fmt"

func main() {
	initialArray := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("Given array: ", initialArray)

	/* call bubblesort to sort the initial array */
	var sortedArray = bubblesort(initialArray)

	/* pick the smallest number */
	fmt.Println("Bubble Sort gives the smallest number as:", sortedArray[0])
}

// bubblesort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - while still using a swap flag AND break point

func bubblesort(dynamicArray []int) []int { // argument is a slice here
	outerCount := 0
	inputArrayLength := len(dynamicArray)
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
		inputArrayLength-- // decrement array length before next iteration, since previous largest value does not need to be involved in next iterations
	}
	return dynamicArray
}
