package main

import (
	"fmt"
	"math/rand"
	"time"
)

// mergeSort()
// - works by:
//   + splitting the iteratively splitting the original arrays into subarrays
//   + until you have sublarrays that are not more than 1 element long i.e. inherently sorted
//   + iteratively compares and determines whether to `append sublist1` to `sublist2` OR `append sublist2 to sublist1`
// - this is NOT happening in place
//   + hence there is a space penalty for the algorithm
func mergeSort(dynamicArray []int) []int {
	temploopRange := len(dynamicArray) // handles reducing elements, so can't use global

	if temploopRange < 2 { // using "<2" instead of "==", handles when dynamicArray=[]
		return dynamicArray
	} else {
		// we could also use:
		//   - demarcationIndex = int(math.ceil(temploopRange/2))
		//   - works for even, odd and prime temploopRange values
		// ceil() this helps to ensure that we get the ceil(of the division) instead of the floor that `a//b` gives
		// demarcationIndex = int(math.ceil(temploopRange/2))`
		demarcationIndex := indexSplitter(0, temploopRange)
		tempArrayOne := dynamicArray[:demarcationIndex] // initialize tempListOne sub-array
		tempArrayTwo := dynamicArray[demarcationIndex:] // initialize tempListTwo sub-array
		tempArrayOne = mergeSort(tempArrayOne)          // recursive call to mergeSorter()
		tempArrayTwo = mergeSort(tempArrayTwo)          // recursive call to mergeSorter()
		return subArrayMerge(tempArrayOne, tempArrayTwo)
	}
}

// subArrayMerge()
// - is called by mergeSort()
// - to handle list merging operations
// - append(tempMergedArray, tempSubArrayTwo[indexSubArrayTwo]) requires
// 		+ tempSubArrayTwo[indexSubArrayTwo] to be an integer
//		+ otherwise unpack it as `[slice]...` to its constituent integers before appending
func subArrayMerge(tempSubArrayOne []int, tempSubArrayTwo []int) []int {

	tempMergedArray := make([]int, 0) // initialise empty Array to merge sub-Arrays into

	loopRangeSubArrayOne := len(tempSubArrayOne)
	loopRangeSubArrayTwo := len(tempSubArrayTwo)

	indexSubArrayOne := 0 // avoids using array.pop() to remove element
	indexSubArrayTwo := 0

	for loopRangeSubArrayOne > indexSubArrayOne && loopRangeSubArrayTwo > indexSubArrayTwo { // a=[1]->len(a)=1
		if tempSubArrayOne[indexSubArrayOne] > tempSubArrayTwo[indexSubArrayTwo] { // test smaller element
			tempMergedArray = append(tempMergedArray, tempSubArrayTwo[indexSubArrayTwo]) // add to end of tempMergeArray
			indexSubArrayTwo++
		} else {
			tempMergedArray = append(tempMergedArray, tempSubArrayOne[indexSubArrayOne]) // add to end of tempMergeArray
			indexSubArrayOne++
		}
	}

	for loopRangeSubArrayOne > indexSubArrayOne { // no elements to merge in SubArrayTwo
		tempMergedArray = append(tempMergedArray, tempSubArrayOne[indexSubArrayOne]) // remaining elements are appended
		indexSubArrayOne++
	}

	for loopRangeSubArrayTwo > indexSubArrayTwo { // no elements to merge in SubArrayOne
		tempMergedArray = append(tempMergedArray, tempSubArrayTwo[indexSubArrayTwo]) // remaining elements are appended
		indexSubArrayTwo++
	}

	return tempMergedArray
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

//
// indexSplitter()
//   - gives the floor of the division between inputListUpperIndex and inputListLowerIndex (i.e. a // b )
//   - alternative is to give the ceiling
//   - we are instead going with floor so as to more easily control the startIndex for rightSubList
//           + i.e. if splitIndex = 2 and hence the last index for the leftSubList
//           + then startIndex for rightSubList is splitIndex += 1
//   - def indexSplitter(inputListLowerIndex: int, inputListUpperIndex: int):
//           numerator = inputListLowerIndex + inputListUpperIndex
//           denumerator = 2
//           splitIndex, quotient = divmod(numerator, denumerator)
//           adjustedSplitIndex = value + bool(quotient)
//           return adjustedSplitIndex
func indexSplitter(inputListLowerIndex int, inputListUpperIndex int) int {
	nuMerator := inputListLowerIndex + inputListUpperIndex
	deNumerator := 2
	splitIndex := nuMerator / deNumerator
	return splitIndex
}

func main() {
	initialArray := createRandomList(0, 1247635, 96)
	arrayShuffler(initialArray) // shuffler the elements of the array
	arrayLength := len(initialArray)

	// sort out the printedSliceLength
	var printedSliceLength int
	if arrayLength > 20 {
		printedSliceLength = 15
	} else {
		printedSliceLength = arrayLength
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("A golang implementation of an Merge Sort algorithm :")
	fmt.Println("  - using randomly generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements\n", arrayLength)
	fmt.Printf("  - first %d elements: %v\n", printedSliceLength, initialArray[:printedSliceLength+1])
	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	/*
	 call bubblesort to sort the initial array
	 initialArray... is used to pass the slice initialArray as variadic parameters
	*/
	// start time counter
	startTime := time.Now()
	var sortedArray = mergeSort(initialArray) // the array is passed as set of variadic arguments
	timeNow := time.Now()
	fmt.Printf("\nMerge Sort gives first 15 elements as: %v \n", sortedArray[:15])
	fmt.Printf("runtime duration: %v seconds \n", timeNow.Sub(startTime).Seconds())
	fmt.Printf("largest number is : %d \n", sortedArray[arrayLength-1])
	fmt.Printf("smallest number is : %d \n\n", sortedArray[0])
}
