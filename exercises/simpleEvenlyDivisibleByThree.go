package main

import "fmt"

var initialValue int64 = 1

func main() {
	for initialValue <= 100 {
		if initialValue%3 == 0 {
			fmt.Println(initialValue, "is evenly divisble by 3")
		}
		initialValue++
	}
}
