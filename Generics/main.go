package main

import "fmt"

// Comparable for all types

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//  LIFO
// type stack[T any] struct {
// 	elements []T
// }

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"Golang"},
	// }

	// fmt.Println(myStack)

	// nums := []int{1, 2, 3}
	// names := []string{"Golang", "Typescript"}
	// printStringSlice(names)
	vals := []bool{true, false, true}
	printSlice(vals, "John")
}
