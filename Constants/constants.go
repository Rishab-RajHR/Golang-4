package main

import "fmt"

const age = 30

var name string = "Golang"

func main() {
	// const name = "Golang"
	// const age = 30
	// :=

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}
