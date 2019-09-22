package main

import "fmt"

func main() {
	var x = "Hello, I am Gopher!"
	fmt.Println(x)

	arr := []byte("test") // this converts the string "test" to it's binary data equivalent in the form of a slice
	str := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(arr)
	fmt.Println(str)
}
