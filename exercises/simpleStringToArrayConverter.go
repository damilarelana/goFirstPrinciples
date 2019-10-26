package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	str := "[2,15,23]"
	var ints []int
	byteData := []byte(str)
	fmt.Println("Strings printed as slice of type Byte", byteData)
	err := json.Unmarshal(byteData, &ints)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Array generated from string:", ints)
	fmt.Println("The type is: ", reflect.TypeOf(ints).String())
}
