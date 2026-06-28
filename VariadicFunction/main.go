package main

import "fmt"

// In Go interface means any data type

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	// In Variadic Function we can pass any number of parameters
	// fmt.Println(1, 2, 3, 5, 68, "Hello")

	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
