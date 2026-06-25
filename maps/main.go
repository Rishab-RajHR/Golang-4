package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict

func main() {

	// Creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "Golang"
	// m["area"] = "Backend"

	// Get an element
	// fmt.Println(m["name"], m["area"])

	// IMP: If key does not exists in the map then it returns zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// delete(m, "price")

	// clear(m)

	// fmt.Println(m)

	// Make the map function without make method

	// m := map[string]int{"price": 40, "phones": 3}

	// v, Ok := m["price"]
	// fmt.Println(v)
	// if Ok {
	// 	fmt.Println("All Ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m1, m2))

}
