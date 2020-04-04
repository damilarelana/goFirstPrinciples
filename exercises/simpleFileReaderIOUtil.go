package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// PathURL declares the type structure we'll parse the YAML or JSON or SQL data into
type PathURL struct {
	Path string `format:"path"`
	URL  string `format:"url"`
}

// JSONHandler parses the JSON file [in byte form]
func JSONHandler(jsonBytes []byte) (pathUrls []PathURL) {
	// parse the JSON file
	pathUrls, err := parseJSON(jsonBytes)
	if err != nil {
		panic(err.Error())
	}
	return pathUrls
}

// parseJSON uses the `json` package to parse the JSON bytes into the Type struct pathURL
//  * json.Unmarshal reads `all` the content into memory at once
func parseJSON(jB []byte) (pathUrls []PathURL, err error) {
	err = json.Unmarshal(jB, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

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

	fmt.Println("==== Print Struct Representation of Data ==== ")
	fmt.Println(JSONHandler(byteSlice))

}
