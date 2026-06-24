package main

import "fmt"

// Slices are dynamic array
// most used construct in go
// + useful methods
func main() {

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2)

	fmt.Println(nums == nil)

}
