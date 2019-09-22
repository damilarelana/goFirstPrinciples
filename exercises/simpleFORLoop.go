package main

import "fmt"

var initialValue int64 = 1

func main() {
	for initialValue <= 10 {
		fmt.Println(initialValue)
		initialValue++
	}
	fmt.Println("==== another loop ====")
	anotherLoop()
}

func anotherLoop() {
	for initialValue := initialValue - 1; initialValue >= 1; initialValue-- {
		fmt.Println(initialValue)
	}
}
