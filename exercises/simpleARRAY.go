package main

import "fmt"

var arrayTotal float64

func main() {

	var initialArray [5]float64
	initialArray[0] = 98
	initialArray[1] = 93
	initialArray[2] = 77
	initialArray[3] = 82
	initialArray[4] = 83

	for initialValue := 0; initialValue < len(initialArray); initialValue++ {
		arrayTotal += initialArray[initialValue]
	}
	fmt.Println(arrayTotal / float64(len(initialArray)))
}
