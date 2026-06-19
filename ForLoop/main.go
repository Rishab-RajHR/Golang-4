package main

import "fmt"

// for -> only contruct in go for looping
func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// Infinite Loop
	// for {
	// 	 println("1")
	// }

	// Classic for loop
	// for i := 0; i <= 3; i++ {
	// 	// break
	// 	// continue

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Range

	for i := range 3 {
		fmt.Println(i)
	}

}
