package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Open current Directory to prep it to be read
	dirSlice, err := os.Open(".") // "." represents the current directory
	if err != nil {
		err := errors.New("unable to open the current Directory") // added a customised error message
		fmt.Println(err)
		return
	}

	// Add a new file to the current directory
	newFile, err := os.Create("../testData/newFileCreatedByCode.text")
	if err != nil {
		err := errors.New("unable to open the file `newFileCreatedByCode.text` in the current Directory") // added a customised error message
		fmt.Println(err)
		return
	}
	defer newFile.Close()                                                                       //defer the closing of the new File till later
	newFile.WriteString("This is a just random gibberish added here to make things more suave") //add some text content to the new file

	// Extract the content of the directory using os.Readdir
	extractedFileInfos, err := dirSlice.Readdir(-1) // this reads all the directory content and saves into slice
	if err != nil {
		err := errors.New("unable to read/save directory content into a slice") //added a customised error message
		fmt.Println(err)
		return
	}

	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")

	for _, v := range extractedFileInfos {
		fmt.Println(v.Name()) // Name for each file in the directly
	}

	fmt.Println(" ")
	fmt.Println("==== ==== ====")
	fmt.Println(" ")

	// Walkthrough and print of another Directory
	filepath.Walk("../testData/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
