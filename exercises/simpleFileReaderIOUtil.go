package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Simpler way to read a file without going through file opening, fileinfo, []byteslice and file.Read
	byteSlice, err := ioutil.ReadFile("../testData/test.txt")
	if err != nil {
		// blank for now
		return
	}

	// Finally print the content of the fileInfo
	str := string(byteSlice) // convert the byte data to string
	fmt.Println(str)
}
