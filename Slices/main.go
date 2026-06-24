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

	var nums = make([]int, 2, 5)
	// capacity -> maximum numbers of a elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	nums = append(nums, 1)
	nums = append(nums, 2)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

}
