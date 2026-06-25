package main

import "fmt"

// numbered sequence of specific length
func main() {
	// Zeroed values
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[0] = 1

	// fmt.Println(nums[0])
	// fmt.Println(nums)

	// Array Length
	// fmt.Println(len(nums))

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	// To declare within the single line
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)

	// 2D Arrays
	nums2 := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums2)

	// - Fixed size , that is predictable
	// - Memory Optimization
	// - Constant time access

}
