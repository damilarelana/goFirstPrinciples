package main

import "fmt"

func square(x *float64) (returnedValue float64) {
	*x = *x * *x
	returnedValue = *x
	return returnedValue
}
func main() {
	var x float64
	fmt.Println("Enter a number to be squared")
	fmt.Scanf("%f", &x)
	squareValue := square(&x)
	fmt.Println("value of x:", squareValue)
}
