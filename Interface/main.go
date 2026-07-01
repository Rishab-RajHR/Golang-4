package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// Open close principle: The payment struct is open for extension but closed for modification. We can add new payment gateways without modifying the existing code.
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment using razorpay
	fmt.Println("Making Payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making Payment using paypal", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}
