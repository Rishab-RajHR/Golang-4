package main

import "fmt"

// Pointers stores the address of the variable

// By Value it gets copied
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In ChangeNum", num)
// }

// Pass Function by references
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)

	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)
}
