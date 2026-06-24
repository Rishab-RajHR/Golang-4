package main

import (
	"fmt"
)

// Slices are dynamic array
// most used construct in go
// + useful methods
func main() {

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// capacity -> maximum numbers of a elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// Another way to make slice

	// nums := []int{}

	// var nums = make([]int, 2, 5)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))

	// // Copy Function
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// Slice Operator
	// var nums = []int{1, 2, 4}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])
	// fmt.Println(nums[:1])

	// Slice Package
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 3}

	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
