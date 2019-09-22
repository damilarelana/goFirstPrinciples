package main

import (
	"fmt"
	"os"
)

func main() {

	// Prepare the file content to be read by openning it first, using os.Open()
	file, err := os.Open("../testData/test.txt") // ".." was used to go to parent directory (introToGo) before we then specified path "/testData/test.txt"
	if err != nil {
		// blank for now
		return
	}
	defer file.Close() // this closes the opened file before the last return

	// Get the fileinfo structure e.g. length of the file in byte need to read the content later
	fileInfo, err := file.Stat()
	if err != nil {
		// blank for now
		return
	}

	// finally read the content itself
	byteSlice := make([]byte, fileInfo.Size()) // create an empty byte slice using the size of the file in byte as reference
	_, err = file.Read(byteSlice)              // we used the Read() method on the already existing byte data in "file" INTO the slice byteSlice
	if err != nil {
		// blank for now
		return
	}

	// Finally print the content of the fileInfo
	str := string(byteSlice) // convert the byte data to string
	fmt.Println(str)
	fmt.Println(byteSlice)
}
