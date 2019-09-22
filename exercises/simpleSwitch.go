package main

import "fmt"

var initialValue int64 = 1

func main() {
	for initialValue <= 100 {
		switch true {
		case (initialValue%3 == 0) && (initialValue%5 == 0):
			fmt.Println("FizzBuzz")
		case (initialValue%3 == 0):
			fmt.Println("Fizz")
		case (initialValue%5 == 0):
			fmt.Println("Buzz")
		default:
			fmt.Println(initialValue)
		}
		initialValue++
	}
}
