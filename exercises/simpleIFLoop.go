package main

import "fmt"

var initialValue int64 = 1

func main() {
	for initialValue <= 10 {
		if initialValue%2 == 0 {
			fmt.Println(initialValue, "is even")
		} else {
			fmt.Println(initialValue, "is odd")
		}
		initialValue++
	}
}
