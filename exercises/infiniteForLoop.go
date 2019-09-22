package main

import "fmt"

func infiniteLoop() {
	var counter float64
	for {
		counter++
		fmt.Println(counter)
	}
}

func main() {
	infiniteLoop()
}
