package main

import "fmt"

func main() {
	var initialNum float64
	fmt.Println("Enter fibonnaci sequence index:")
	fmt.Scanf("%f", &initialNum)
	if initialNum < 0 {
		fmt.Println("ERROR - Please enter a non-negative number: ")
		fmt.Scanf("%f", &initialNum)
	}
	computedValue := fibonnaci(initialNum)
	fmt.Println("Fibonacci value is:", computedValue)
}

func fibonnaci(initialNum float64) (computedNum float64) {
	zeroOneCase := []float64{0, 1} // this handles both index 0 and 1
	for _, v := range zeroOneCase {
		if initialNum == v {
			computedNum = initialNum
			return
		}
	}
	computedNum = fibonnaci(initialNum-1) + fibonnaci(initialNum-2)
	return
}
