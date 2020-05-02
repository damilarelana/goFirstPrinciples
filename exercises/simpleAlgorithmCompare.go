package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// insertionSort()
func insertionSort(dynamicArray []int) []int {
	inputArrayLength := len(dynamicArray)
	if inputArrayLength == 1 {
		return dynamicArray
	} else {
		// we assume that element at index `0` i.e. outerCount = 0, is already sorted, hence why the unsorted starts at outerCount = 1
		outerCount := 1 // initialising unsorted list index to the first one to be removed from unsorted [we ]
		for outerCount < inputArrayLength {
			innerCount := outerCount // re-initialising sorted list's max index to allow countdown
			for innerCount > 0 {     // handles inner loop i.e. the `sorted list loop`. greater than zero
				if dynamicArray[innerCount-1] > dynamicArray[innerCount] { // this already carters for assuming list[0] is sorted
					dynamicArray[innerCount-1], dynamicArray[innerCount] = dynamicArray[innerCount], dynamicArray[innerCount-1]
				}
				innerCount-- // this is different to bubbleSort i.e. where there is an increment. Here we are decreasing the unsorted set
			}
			outerCount++ // here we are increasing the sorted set boundaries [which weirdly also acts like the next `first element of the now shrinking unsorted set`]
		}
		return dynamicArray
	}
}

// elegantBubbleSort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - without using a swap flag AND break point
func elegantBubbleSort(dynamicArray []int) []int { // variadic arguments are used here
	outerCount := 0
	inputArrayLength := len(dynamicArray)
	for outerCount < len(dynamicArray) {
		innerCount := 0
		for innerCount < (inputArrayLength - 1) {
			if dynamicArray[innerCount] > dynamicArray[innerCount+1] {
				dynamicArray[innerCount], dynamicArray[innerCount+1] = dynamicArray[innerCount+1], dynamicArray[innerCount]
			}
			innerCount++
		}
		outerCount++
		inputArrayLength-- // decrement array length before next iteration, since previous largest value does not need to be involved in next iterations
	}

	return dynamicArray
}

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

// selectionSort()
func selectionSort(dynamicArray []int) []int {
	inputArrayLength := len(dynamicArray)
	if inputArrayLength == 1 {
		return dynamicArray
	} else {
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
}

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

// checkSortedArrayEquivalence()
// - leverages reflect.DeepEqual
// - to check if two ordered lists are equal
// - propagates the test across board
func checkSortedArrayEquivalence(selectionSorted []int, hybridBubblesorted []int, elegantBubblesorted []int, mergesorted []int, insertionsorted []int) {
	sShBS := reflect.DeepEqual(hybridBubblesorted, selectionSorted)
	hBSmS := reflect.DeepEqual(hybridBubblesorted, mergesorted)
	hBSiS := reflect.DeepEqual(hybridBubblesorted, insertionsorted)
	hBSeBS := reflect.DeepEqual(hybridBubblesorted, elegantBubblesorted)
	if sShBS && hBSmS && hBSiS && hBSeBS {
		print("\nAll algorithms give the same sorted arrays")
	} else {
		print("\nAll algorithms did NOT give the same sorted arrays")
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

	initialArray := createRandomList(0, 1247635, 48)
	arrayShuffler(initialArray) // shuffler the elements of the array
	arrayLength := len(initialArray)

	// distinct inputArray are not required because:
	// - golang copies the original array when passing it to the function
	// - we can repeat `initialArray` across board when passing it as an argument to the different sorting algorithms
	// - passing the actual original array is what pointers are for [but not applicable in this case]
	// - thus this below is not required
	//		+	hBSInputArray := make([]int, arrayLength)
	//		+ copy(hBSInputArray, initialArray)
	// - they are however used nonetheless just to over-compensate

	hBSInputArray := make([]int, arrayLength)
	copy(hBSInputArray, initialArray)

	eBSInputArray := make([]int, arrayLength)
	copy(eBSInputArray, initialArray)

	mSInputArray := make([]int, arrayLength)
	copy(mSInputArray, initialArray)

	iSInputArray := make([]int, arrayLength)
	copy(iSInputArray, initialArray)

	sSInputArray := make([]int, arrayLength)
	copy(sSInputArray, initialArray)

	// sort out the printedSliceLength
	var printedSliceLength int
	if arrayLength > 20 {
		printedSliceLength = 15
	} else {
		printedSliceLength = arrayLength
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("Comparing Golang performance of 5 algorithms [ selectionSort + mergeSort + hybridBubbleSort + elegantBubbleSort + insertionSort]:")
	fmt.Println("  - using randomly generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements\n", arrayLength)
	fmt.Printf("  - first %d elements: %v\n", printedSliceLength, initialArray[:printedSliceLength+1])
	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// start time counter
	mSStartTime := time.Now()
	mergesorted := mergeSort(mSInputArray) // the array is passed as set of variadic arguments
	mSTimeNow := time.Now()
	fmt.Printf("Merge Sort gives first %d elements as: %v \n", printedSliceLength, mergesorted[:printedSliceLength+1])
	fmt.Printf("runtime duration: %v seconds \n", mSTimeNow.Sub(mSStartTime).Seconds())
	fmt.Printf("largest number is : %d \n", mergesorted[arrayLength-1])
	fmt.Printf("smallest number is : %d \n", mergesorted[0])
	fmt.Println("================================")

	// start time counter
	iSStartTime := time.Now()
	insertionsorted := insertionSort(iSInputArray) // the array is passed as set of variadic arguments
	iSTimeNow := time.Now()
	fmt.Printf("\nInsertion Sort gives first %d elements as: %v \n", printedSliceLength, insertionsorted[:printedSliceLength+1])
	fmt.Printf("runtime duration: %v seconds \n", iSTimeNow.Sub(iSStartTime).Seconds())
	fmt.Printf("largest number is : %d \n", insertionsorted[arrayLength-1])
	fmt.Printf("smallest number is : %d \n", insertionsorted[0])
	fmt.Println("================================")

	// start time counter
	sSStartTime := time.Now()
	selectionSorted := selectionSort(sSInputArray) // the array is passed as set of variadic arguments
	sSTimeNow := time.Now()
	fmt.Printf("\nSelection Sort gives first %d elements as: %v \n", printedSliceLength, selectionSorted[:printedSliceLength+1])
	fmt.Printf("runtime duration: %v seconds \n", sSTimeNow.Sub(sSStartTime).Seconds())
	fmt.Printf("largest number is : %d \n", selectionSorted[arrayLength-1])
	fmt.Printf("smallest number is : %d \n", selectionSorted[0])
	fmt.Println("================================")

	// start time counter
	hBSStartTime := time.Now()
	hybridBubblesorted := hybridBubbleSort(hBSInputArray) // the array is passed as set of variadic arguments
	hBSTimeNow := time.Now()
	fmt.Printf("\nHybrid Bubble Sort gives first %d elements as: %v \n", printedSliceLength, hybridBubblesorted[:printedSliceLength+1])
	fmt.Printf("runtime duration: %v seconds \n", hBSTimeNow.Sub(hBSStartTime).Seconds())
	fmt.Printf("largest number is : %d \n", hybridBubblesorted[arrayLength-1])
	fmt.Printf("smallest number is : %d \n", hybridBubblesorted[0])
	fmt.Println("================================")

	// start time counter
	eBSStartTime := time.Now()
	elegantBubblesorted := elegantBubbleSort(eBSInputArray) // the array is passed as set of variadic arguments
	eBSTimeNow := time.Now()
	fmt.Printf("\nElegant Bubble Sort gives first %d elements as: %v \n", printedSliceLength, elegantBubblesorted[:printedSliceLength+1])
	fmt.Printf("runtime duration: %v seconds \n", eBSTimeNow.Sub(eBSStartTime).Seconds())
	fmt.Printf("largest number is : %d \n", elegantBubblesorted[arrayLength-1])
	fmt.Printf("smallest number is : %d \n", elegantBubblesorted[0])
	fmt.Println("================================")

	// Check if both sorted list are equivalent
	checkSortedArrayEquivalence(selectionSorted, hybridBubblesorted, elegantBubblesorted, mergesorted, insertionsorted)
}
