package main

import "fmt"

func advancedAddition(args ...float64) (totalValue float64) {
	for _, v := range args {
		totalValue += v
	}
	return totalValue
}

func main() {
	//	testValue := []float64{1, 2, 3, 4, 5}
	//	fmt.Println("Total Values:", advancedAddition(testValue...))
	fmt.Println("Total Values:", advancedAddition([]float64{1, 2, 3, 4, 5}...))
}
