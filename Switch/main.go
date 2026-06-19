package main

import "fmt"

func main() {
	// Simple Watch
	// i := 6

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("fifth")
	// default:
	// 	fmt.Println("Invalid Data")
	// }

	// Multiple Condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's workday")
	// }

	// Type Switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its an string")
		case bool:
			fmt.Println("Its an boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI(45.8)
}
