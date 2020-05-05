package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/gorilla/mux"
)

// ErrMsgHandler defines the error message handler
func ErrMsgHandler(msg string, err error) {
	if err != nil {
		log.Println(msg, err.Error())
		os.Exit(1)
	}
}

// insertionSort()
func insertionSort(dynamicArray []int) ([]int, map[int][]int) {
	iSPlotDataMap, iSMapKey := initAnimationDataMap(dynamicArray) // initialize animation data gathering
	iSPlotDataMapPtr := &iSPlotDataMap                            // create pointer to the plotDataMap
	inputArrayLength := len(dynamicArray)
	if inputArrayLength == 1 {
		return dynamicArray, iSPlotDataMap
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
				// golang inherently creates a copy of the argument when passing it to a function
				//	 - but it is still being passed a copy of the same address [so ends up with same problem as python]
				tempArrayState := make([]int, inputArrayLength)
				copy(tempArrayState, dynamicArray)
				updateAnimationDataMap(tempArrayState, iSMapKey, iSPlotDataMapPtr, inputArrayLength)
				iSMapKey++ //  increase Map key before it is used in updateAnimationDataMap
			}
			outerCount++ // here we are increasing the sorted set boundaries [which weirdly also acts like the next `first element of the now shrinking unsorted set`]
		}
		return dynamicArray, iSPlotDataMap
	}
}

// elegantBubbleSort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - without using a swap flag AND break point
func elegantBubbleSort(dynamicArray []int) ([]int, map[int][]int) {
	eBSPlotDataMap, eBSMapKey := initAnimationDataMap(dynamicArray) // initialize animation data gathering
	eBSPlotDataMapPtr := &eBSPlotDataMap                            // create pointer to the plotDataMap
	outerCount := 0
	inputArrayLength := len(dynamicArray)
	for outerCount < len(dynamicArray) {
		innerCount := 0
		for innerCount < (inputArrayLength - 1) {
			if dynamicArray[innerCount] > dynamicArray[innerCount+1] {
				dynamicArray[innerCount], dynamicArray[innerCount+1] = dynamicArray[innerCount+1], dynamicArray[innerCount]
			}
			innerCount++
			tempArrayState := make([]int, inputArrayLength)
			copy(tempArrayState, dynamicArray)
			updateAnimationDataMap(tempArrayState, eBSMapKey, eBSPlotDataMapPtr, inputArrayLength)
			eBSMapKey++ // increase Map key before it is used in updateAnimationDataMap
		}
		outerCount++
		inputArrayLength-- // decrement array length before next iteration, since previous largest value does not need to be involved in next iterations
	}

	return dynamicArray, eBSPlotDataMap
}

// hybridBubbleSort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - while still using a swap flag AND break point
func hybridBubbleSort(dynamicArray []int) ([]int, map[int][]int) {
	hBSPlotDataMap, hBSMapKey := initAnimationDataMap(dynamicArray) // initialize animation data gathering
	hBSPlotDataMapPtr := &hBSPlotDataMap                            // create pointer to the plotDataMap
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
			tempArrayState := make([]int, inputArrayLength)
			copy(tempArrayState, dynamicArray)
			updateAnimationDataMap(tempArrayState, hBSMapKey, hBSPlotDataMapPtr, inputArrayLength)
			hBSMapKey++ // increase Map key before it is used in updateAnimationDataMap
		}
		/* exiting from loop when already sorted input and sorting completion */
		if !swapflag {
			break OuterForLoop
		}
		outerCount++
		inputArrayLength-- // decrement array length before next iteration, since previous largest value does not need to be involved in next iterations
	}

	return dynamicArray, hBSPlotDataMap
}

// selectionSort()
func selectionSort(dynamicArray []int) ([]int, map[int][]int) {
	sSPlotDataMap, sSMapKey := initAnimationDataMap(dynamicArray) // initialize animation data gathering
	sSPlotDataMapPtr := &sSPlotDataMap                            // create pointer to the plotDataMap
	inputArrayLength := len(dynamicArray)
	if inputArrayLength == 1 {
		return dynamicArray, sSPlotDataMap
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

			tempArrayState := make([]int, inputArrayLength)
			copy(tempArrayState, dynamicArray)
			updateAnimationDataMap(tempArrayState, sSMapKey, sSPlotDataMapPtr, inputArrayLength)
			sSMapKey++ //  increase Map key before it is used in updateAnimationDataMap

			outerCount++                              // increase outer counter i.e. expanding the sorted set
			minElement = outerCount                   // reset new temporary minimum index e.g. if initial was index `0`, it would now be `1`
			if outerCount == (inputArrayLength - 1) { // i.e. only one unsorted element remains, break outer loop
				break OuterForLoop
			}
			// note that we CANNOT use the optimization (loopRange -= 1) since we are shifting values/index around
			// as such the last value after every iteration can still need to be touched
			// this is one difference with BubbleSort() where the last index can be removed from dataset after every loop
		}
		return dynamicArray, sSPlotDataMap
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

// createRandomArray()
func createRandomArray(minValue int, maxValue int, stepValue int) (array []int) {
	totalIncrements := (maxValue - minValue) / stepValue // initialize the array
	array = make([]int, totalIncrements+1)
	for i, _ := range array { // populate the array elements
		if i == 0 {
			array[i] = minValue
		} else {
			array[i] = array[i-1] + stepValue
		}
	}
	return array // return the randomized array
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
func indexSplitter(inputListLowerIndex int, inputListUpperIndex int) int {
	nuMerator := inputListLowerIndex + inputListUpperIndex
	deNumerator := 2
	splitIndex := nuMerator / deNumerator
	return splitIndex
}

// initAnimationDataMap()
// - initializes the dictionary used to store arrayState for each sorting loop cycle
// - takes in an array to be sorted
// - returns an initialized `map` and initialized `map key` pair
// - map key is initialized as '0'
func initAnimationDataMap(inputArray []int) (map[int][]int, int) {
	key := 0                           // initialize the map's key to store
	plotDataMap := make(map[int][]int) // initialize the map itself
	return plotDataMap, key
}

/*
 augmentArray()
   - is mostly applicable to algorithms that break down the Array before sorting e.g. MergeSort()
   - it takes currentArray state
   - it gets length of currentArray i.e. currentArrayLength
   - it takes in length of original unsorted Array i.e. inputArrayLength
   - checks if there has been a decomposition i.e. if inputArrayLength != currentArrayLength
       + if false, it simply returns `currentArray` to the callback function
       + if true, it pads the currentArray with `0` at the end so as to make the `arrays` inside `stateDataArray` to be of same length i.e.
       + returns the padded currentArray
*/
func augmentArray(currentArray []int, inputArrayLength int) []int {
	currentArrayLength := len(currentArray)
	if currentArrayLength != inputArrayLength {
		padLength := inputArrayLength - currentArrayLength // number of extra zeros required to pad the unbalanced array
		padArray := make([]int, padLength)
		currentArray = append(currentArray, padArray...)
	}
	return currentArray
}

/* updateAnimationDataMap()
	- takes in current state of array being sorted
	- takes in current key i.e. assuming it was already incremeneted before being passed
	- takes in pointer to the current map state i.e. to ensure we alter that actual map and not a copy (due to Golang internals)
	- takes in inputArrayLength
	- call augmentArray() to check and augment the current stateData i.e. arrayWhileSorting
	- deference the plotDataMapPtr i.e. (*plotDataMapPtr), so as to allow for indexing [as pointers cannot be indexed]
	- saves the stateData (initial, current and future) into the map, with specific index value : key
	- does not return anything as it is working on the array and dictionary in-place
	- dictionary is being used so as to obtain:
 		+ a 2-D array once we extract the `array of array` from the dictionary
 		+ as required by the animation approach [https://brushingupscience.com/2016/06/21/matplotlib-animations-the-easy-way/]
 		+ `array of array` approach is also being used to improved performance:
*/
func updateAnimationDataMap(arrayWhileSorting []int, key int, plotDataMapPtr *map[int][]int, inputArrayLength int) {
	augmentedArray := augmentArray(arrayWhileSorting, inputArrayLength)
	(*plotDataMapPtr)[key] = augmentedArray // add augmentArray to map : Golang automatically accommodates new key
}

// genericBar
func genericBar(xAxisItems []int, algorithmName string, stateData []int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: algorithmName},
		charts.ToolboxOpts{Show: true},
	)
	bar.AddXAxis(xAxisItems).AddYAxis("Array State", stateData)
	return bar
}

// overlapBar
func overlapBar(baseBar *charts.Bar, xAxisItems []int, algorithmName string, plotDataArrays [][]int) *charts.Bar {
	baseBar.Overlap(genericBar(xAxisItems, algorithmName, plotDataArrays[1]))
	return baseBar
}

// createAnimation()
func createAnimation(plotDataArrays [][]int, arrayLength int, algorithmName string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			lastIndex := arrayLength - 1
			xAxisItems := createRandomArray(0, lastIndex, 1)
			page := charts.NewPage()
			err := os.Remove("bar.html")
			if err != nil {
				errMsg := fmt.Sprintf("Unable to remove previous 'bar.html' for plotDataArray")
				ErrMsgHandler(errMsg, err)
			}
			f, err := os.Create("bar.html")
			if err != nil {
				errMsg := fmt.Sprintf("Unable to create bar.html for plotDataArray")
				ErrMsgHandler(errMsg, err)
			}
			page.Add(
				overlapBar(genericBar(xAxisItems, algorithmName, plotDataArrays[0]), xAxisItems, algorithmName, plotDataArrays),
			)
			page.Render(w, f)
		})
}

// defaultMux defines the router Mux
func defaultMux(plotDataArrays [][]int, arrayLength int, algorithmName string) *mux.Router {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.Handle("/", createAnimation(plotDataArrays, arrayLength, algorithmName))
	return newRouter
}

// InitializeWebServer
func initRenderWebServer(port string, plotDataArrays [][]int, arrayLength int, algorithmName string) {
	mux := defaultMux(plotDataArrays, arrayLength, algorithmName) // create an instance of defaultMux()
	log.Println("animation rendering at http://127.0.0.1:" + port)
	// log.Fatal(errors.Wrap(http.ListenAndServe(":"+port, mux), "Failed to start webserver at port:"+port))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// getFuncName()
// - returns the name of a given function as a string
func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

/* parsePlotData()
- takes in the returned plotDataMap
- parses it through a for loop so as to generate a `array of array`
  + to extract the [index, stateData] at each dictionary key
*/
func parsePlotData(plotDataMap map[int][]int) [][]int {
	mapLength := len(plotDataMap)               // get number of items in the map i.e. how many stateData do we have
	arrayOfArrays := make([][]int, mapLength)   // create an empty array of Arrays to be populated
	for index, stateData := range plotDataMap { // iterated over the Map to extract stateData and insert into an arrayOfArrays
		arrayOfArrays[index] = stateData
	}
	return arrayOfArrays // returns the array of arrays, required by animation
}

// main
func main() {

	// initialArray := createRandomArray(0, 1247635, 48)
	initialArray := createRandomArray(0, 480, 48)
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
		printedSliceLength = arrayLength - 1
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	fmt.Println("Comparing Golang performance of 5 algorithms [ selectionSort + mergeSort + hybridBubbleSort + elegantBubbleSort + insertionSort]:")
	fmt.Println("  - using randomly generated data")
	fmt.Printf("  - of an array of integer values\n")
	fmt.Printf("  - with %d elements\n", arrayLength)
	fmt.Printf("  - first %d elements: %v\n", printedSliceLength+1, initialArray[:printedSliceLength+1])
	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// start time counter
	mSStartTime := time.Now()
	mergesorted := mergeSort(mSInputArray) // the array is passed as set of variadic arguments
	mSTimeNow := time.Now()
	fmt.Printf("Merge Sort gives first %d elements as: %v \n", printedSliceLength+1, mergesorted[:printedSliceLength+1])
	fmt.Printf("runtime: %v seconds \n", mSTimeNow.Sub(mSStartTime).Seconds())
	fmt.Printf("largest num: %d \n", mergesorted[arrayLength-1])
	fmt.Printf("smallest num: %d \n", mergesorted[0])
	fmt.Println("================================")

	// start time counter
	iSStartTime := time.Now()
	insertionsorted, iSPlotData := insertionSort(iSInputArray) // the array is passed as set of variadic arguments
	iSTimeNow := time.Now()
	fmt.Printf("\nInsertion Sort gives first %d elements as: %v \n", printedSliceLength+1, insertionsorted[:printedSliceLength+1])
	fmt.Printf("runtime: %v seconds \n", iSTimeNow.Sub(iSStartTime).Seconds())
	fmt.Printf("largest num: %d \n", insertionsorted[arrayLength-1])
	fmt.Printf("smallest num: %d \n", insertionsorted[0])
	fmt.Println("================================")
	// animation creation
	algorithmName := getFuncName(insertionSort)                                    // get the function name as a string. This gives "main.insertionSort"
	algorithmName = strings.TrimLeft(strings.TrimLeft(algorithmName, "main"), ".") // trim the string to remove extraneous stuff
	plotDataArrays := parsePlotData(iSPlotData)
	initRenderWebServer("8080", plotDataArrays, arrayLength, algorithmName)
	print("================================")

	// start time counter
	sSStartTime := time.Now()
	selectionSorted, _ := selectionSort(sSInputArray) // the array is passed as set of variadic arguments
	sSTimeNow := time.Now()
	fmt.Printf("\nSelection Sort gives first %d elements as: %v \n", printedSliceLength+1, selectionSorted[:printedSliceLength+1])
	fmt.Printf("runtime: %v seconds \n", sSTimeNow.Sub(sSStartTime).Seconds())
	fmt.Printf("largest num: %d \n", selectionSorted[arrayLength-1])
	fmt.Printf("smallest num: %d \n", selectionSorted[0])
	fmt.Println("================================")

	// start time counter
	hBSStartTime := time.Now()
	hybridBubblesorted, _ := hybridBubbleSort(hBSInputArray) // the array is passed as set of variadic arguments
	hBSTimeNow := time.Now()
	fmt.Printf("\nHybrid Bubble Sort gives first %d elements as: %v \n", printedSliceLength+1, hybridBubblesorted[:printedSliceLength+1])
	fmt.Printf("runtime: %v seconds \n", hBSTimeNow.Sub(hBSStartTime).Seconds())
	fmt.Printf("largest num: %d \n", hybridBubblesorted[arrayLength-1])
	fmt.Printf("smallest num: %d \n", hybridBubblesorted[0])
	fmt.Println("================================")

	// start time counter
	eBSStartTime := time.Now()
	elegantBubblesorted, _ := elegantBubbleSort(eBSInputArray) // the array is passed as set of variadic arguments
	eBSTimeNow := time.Now()
	fmt.Printf("\nElegant Bubble Sort gives first %d elements as: %v \n", printedSliceLength+1, elegantBubblesorted[:printedSliceLength+1])
	fmt.Printf("runtime: %v seconds \n", eBSTimeNow.Sub(eBSStartTime).Seconds())
	fmt.Printf("largest num: %d \n", elegantBubblesorted[arrayLength-1])
	fmt.Printf("smallest num: %d \n", elegantBubblesorted[0])
	fmt.Println("================================")

	// Check if both sorted list are equivalent
	checkSortedArrayEquivalence(selectionSorted, hybridBubblesorted, elegantBubblesorted, mergesorted, insertionsorted)
}
