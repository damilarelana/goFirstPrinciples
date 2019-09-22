package main

import "fmt"

var arrayTotal float64

func main() {
	initialArray := [5]float64{98, 93, 77, 82, 83}
	for _, value := range initialArray {
		arrayTotal += value
	}
	fmt.Println(arrayTotal / float64(len(initialArray)))
}
