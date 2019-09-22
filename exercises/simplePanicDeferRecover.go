package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer initiated")
		str := recover()
		fmt.Println("recover initiated")
		fmt.Println(str)
	}()
	// now at the panic function
	fmt.Println("at panic function")
	panic("PANIC")
}
