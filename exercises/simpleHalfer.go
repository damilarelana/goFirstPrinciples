package main

import "fmt"

func halfer(testValue int64) (divisionValue int64, isItEven bool, remainderValue int64) {
	if testValue%2 == 0 {
		divisionValue = testValue / 2
		var remainderValue int64
		return divisionValue, true, remainderValue
	}
	remainderValue = testValue % 2
	divisionValue = testValue / 2
	return divisionValue, false, remainderValue
}

func main() {
	var testValue int64
	fmt.Println("Enter an Integer:")
	fmt.Scanf("%d", &testValue)
	divisionValue, isItEven, remainderValue := halfer(testValue)
	fmt.Println("====")
	fmt.Println("division output:", divisionValue)
	fmt.Println("is it an integer:", isItEven)
	fmt.Println("remainder:", remainderValue)
}
