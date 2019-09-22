package main

import "fmt"

func makeOddNumber() func() uint {
	initialValue := uint(0)
	return func() (returnValue uint) {
		returnValue = initialValue
		// handles edge case for when initialValue is 0 AND  0%2 == 0
		if initialValue == 0 {
			returnValue = initialValue
		} else if returnValue%2 == 0 { // handles edge of even numbers
			returnValue = initialValue + 1
			initialValue = returnValue
		}
		initialValue++
		return
	}
}

func main() {
	nextOdd := makeOddNumber()
	fmt.Println(nextOdd()) // 0
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
	fmt.Println(nextOdd()) // 7
	fmt.Println(nextOdd()) // 9
	fmt.Println(nextOdd()) // 11
}
