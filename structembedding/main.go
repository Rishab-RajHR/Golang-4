package main

import (
	"fmt"
	"time"
)

// Order struct

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {
	// newCustomer := customer{
	// 	name:  "John",
	// 	phone: "123456789",
	// }

	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "John",
			phone: "123456789",
		},
	}

	newOrder.customer.name = "Robin"

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}
