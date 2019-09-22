package main

import "fmt"

/* declare the variables */
const xGlobalScope = "first"

var exoManOWar = "to love is good"
var name string

func anotherFunction() {
	yNew := xGlobalScope
	yNew += " " + exoManOWar
	fmt.Scanf("%s", &name)
	fmt.Println(name, "being the", yNew)
}

func main() {
	var y = "second"
	thirdString := "third"
	thirdString += y + thirdString
	fmt.Println(thirdString)
	anotherFunction()
}
