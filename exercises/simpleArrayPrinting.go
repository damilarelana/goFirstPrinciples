package main

import "fmt"

func main() {
	initialArray := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("Array printed as a variadic arguments: ", initialArray...)
	fmt.Println("Array printed as is: ", initialArray)
}
