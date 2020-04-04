package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {

	// Simpler way to read a file without going through file opening, fileinfo, []byteslice and file.Read
	byteSlice, err := ioutil.ReadFile("../testData/pathsData.json")

	fmt.Println("==== Print DataType of ByteSlice ==== ")
	fmt.Println(reflect.TypeOf(byteSlice))

	fmt.Println("==== Print Data Inside ByteSlice ==== ")
	fmt.Println(byteSlice)
	if err != nil {
		// blank for now
		return
	}

	// Finally print the content of the fileInfo
	str := string(byteSlice) // convert the byte data to string
	fmt.Println("==== Print ByteSlice (with data converted to string) ==== ")
	fmt.Println(str)
}
