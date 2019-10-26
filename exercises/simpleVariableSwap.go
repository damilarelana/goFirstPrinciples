package main

import "fmt"

func swap(x *float64, y *float64) (float64, float64) {
	// `Method 1`
	// temp := *x
	// *x = *y
	// *y = temp

	// `Method 2`
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
	return *x, *y
}
func main() {
	var x, y float64
	fmt.Println("Enter two numbers [separated by a space] to be swapped:")

	// Reads 2 input variables x and y, that are separated by a space
	fmt.Scanf("%f %f", &x, &y)
	fmt.Println("initial x:", x)
	fmt.Println("initial y:", y)
	swappedX, swappedY := swap(&x, &y)
	fmt.Println("swapped x:", swappedX)
	fmt.Println("swapped y:", swappedY)
}
