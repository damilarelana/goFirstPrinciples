package main

import (
	"container/list"
	"fmt"
)

// This is an implementation of the doubly-linked list
func main() {
	// create initial empty list newList from the struct List
	var newList list.List

	//Append new values to the list using method `PushBack()`
	newList.PushBack(map[string]string{"Name": "John Doe"})
	newList.PushBack(map[string]string{"Name": "Bruce Banner"})

	// Iterate and Print each item in the list
	for e := newList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(map[string]string))
	}
}
