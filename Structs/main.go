package main

import (
	"fmt"
	"time"
)

// Order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct", myOrder2)

	// fmt.Println("Order struct", myOrder)
}
