package main

import "fmt"

// Iterating over data structures
func main() {
	// nums := []int{6, 7, 8}

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	// m := map[string]string{"fname": "John", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// For getting the key only

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// Unicode code point rune
	// Starting byte of rune
	// 255 -> 1 byte, more than 255 -> 2 byte
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
