package main

import "fmt"

var measureInFeets float64

func main() {
	fmt.Println("How many feet: ")
	fmt.Scanf("%f", &measureInFeets)
	measureInMeters := 0.3048 * measureInFeets
	fmt.Println("this gives", measureInMeters, "meters")
}
