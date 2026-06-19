package main

import "fmt"

func main() {

	// age := 19

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// age := 19

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is Teenager")
	// } else {
	// 	fmt.Println("Person is a Kid")
	// }

	// Logical Operator in if-else

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("Yes")
	}

	// We can declare a variable inside if
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is Teenager", age)

		// Go does not have Ternary Operator
	}
}
