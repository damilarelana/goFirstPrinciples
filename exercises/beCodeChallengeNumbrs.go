package main

/*
 * Complete the 'budgetShopping' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY bundleQuantities
 *  3. INTEGER_ARRAY bundleCosts
 */

// Assumptions:
// - bundleQuantities are integers
// - bundleCosts are integers
// - budget are integers

// Implementation Logic:
// - Use unsorted bundleQuantities and bundleCosts
// 	+ as sorting messes up the indexing [which is required]
// - assume that maximum or minimum is not known [for either cost or quantity]
// - 3 loops would be required
//	+ 1st loop [1st inner] to compute:
//		- bundleMultiples = budget / bundleCost i.e. how many bundles: using Golang's integer division [which returns just the quotient]
//		- possibleQuantity = `bundleMultiples * bundleQuantity`
//		- possibleCost = `bundleMultiples * bundleCost`
//  	- remainingBudget = `budget % bundleCost` or `budget - possibleCost`
//		- choiceIndex = i.e. current index
//		- compute possiblePurchase per index [using bundleCosts[index]]
//
//	+ 2nd loop [2nd inner] to [need to check if you are the last index, since `index + 1` does not exist
//			- test if remainingBudget is 0
//			- if remainingBudget != 0, then
//					+ create a slice of `budgetQuantities` and `budgetCosts` by temporarily shifting index forward by 1
//					+ use the remainingBudget, as if it's a new budget [but only for the new remaining Slice]
//					+ use the new `budgetQuantitiesSlice` and `budgetCostsSlice` to recompute the following:
//						- extraBundleMultiples = remainingBudget / bundleCost i.e. this would return a zero
//						- extraQuantity = `extraBundleMultiples * bundleQuantity`
//						- possibleCost = `bundleMultiples * bundleCost`
//  					- remainingBudget = `budget % bundleCost` or `budget - possibleCost`
//						- choiceIndex = i.e. current index
//						- compute possiblePurchase per index [using bundleCosts[index]]
//
//
//
//
//
//
//	+ one loop [outer] to repeat same computation across other all indices and then compare to determine
//		- compute maxQuantity per index [using bundleQuantities[index]
//	+ one index at the time i.e. [having access to the quantity + cost at that index]
//
//	+ to determine costValue where
//

func budgetShopping(n int32, bundleQuantities []int32, bundleCosts []int32) int32 {

	sortedQuantities := bubbleSort(bundleQuantities) // sorted in descending order
	sortedCosts := bubbleSort(bundleCosts)

	maxQuantity := sortedQuantities[0]
	minCost := sortedCosts[len(sortedQuantities)-1]

	choiceQuantityIndex := len(sortedQuantities) - 1 // assumes that the largest quantity possible would element in last index of bundleQuantity
	choiceCostIndex := 0                             // assumes that the lowest cost possible would element in last index of bundleCosts

	for qIndex, quantity := range bundleQuantities {
		for cIndex, cost := range bundleCosts {
		costValue:
			if value < n {

			}
		}
	}
}

// bubbleSort()
// - dynamically adjusts
// - by reducing the inputListLength, after each inner loop iteration
// - while still using a swap flag AND break point
// - sorts in descending order
func bubbleSort(dynamicArray []int) []int { // argument is a slice here
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
