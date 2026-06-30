package main

import "fmt"

// Order struct

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder)
	// If you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())

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
