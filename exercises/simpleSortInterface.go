package main

import (
	"fmt"
	"sort"
)

// This is an implementation of Go's internal sort.Sort interface  that's focused on slices of ints and floats.
// It involves confluence between:
// 1. a struct i.e. define structure of what to be sorted
// 2. sort.inteface implemtations i.e. Len(), Less(), and Swap()
//
// Benefit is that:
// 1. it does require ground up implemtations of stuff like bubblesort, mergesort etc.
// 2. it is particularly applicable to maps [hence allowing sort via either key or value pair]

// Person struct defined as the data type we seek to sorted
type Person struct {
	Name string
	Age  int
}

// ByName is a new type [similar to and slices of Person], that we would use to implement the sort.inteface [so we can get access to the sort.Sort method set]
type ByName []Person

// ByAge is a new type [similar to and slices of Person], that we would use to implement the sort.inteface [so we can get access to the sort.Sort method set]
type ByAge []Person

// Define methods for ByName
func (byNameInstance ByName) Len() int {
	return len(byNameInstance) // get the total number of elements in that instance
}
func (byNameInstance ByName) Less(i, j int) bool {
	return byNameInstance[i].Name < byNameInstance[j].Name // check if the Name at element i is less than Name at element j
}
func (byNameInstance ByName) Swap(i, j int) {
	byNameInstance[i], byNameInstance[j] = byNameInstance[j], byNameInstance[i] // swap the whole key value pair {i,j}
}

// Define methods for ByAge
func (byNameInstance ByAge) Len() int {
	return len(byNameInstance) // get the total number of elements in that instance
}
func (byNameInstance ByAge) Less(i, j int) bool {
	return byNameInstance[i].Age < byNameInstance[j].Age // check if the Name at element i is less than Name at element j
}
func (byNameInstance ByAge) Swap(i, j int) {
	byNameInstance[i], byNameInstance[j] = byNameInstance[j], byNameInstance[i] // swap the whole key value pair {i,j}
}

func main() {
	// create a test fieldset of data
	testData := []Person{
		{"John Doe", 19},
		{"Bruce Banner", 30},
		{"Sherlock Holmes", 24},
		{"Adam Brashear", 45},
	}

	//Now call the sort.Sort method of Go to help sort the slice/map data
	fmt.Println(" ")
	fmt.Println("==== Original unsorted Test Data ====")
	fmt.Println(testData)
	fmt.Println(" ")
	fmt.Println("==== Test Data sorted by Name ====")
	sort.Sort(ByName(testData))
	fmt.Println(testData)
	fmt.Println(" ")
	fmt.Println("==== Test Data sorted by Age ====")
	sort.Sort(ByAge(testData))
	fmt.Println(testData)
}
