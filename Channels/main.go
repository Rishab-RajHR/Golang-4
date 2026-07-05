package main

import (
	"fmt"
	"time"
)

// Send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult   // blocking
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("Processing...")
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done Sending")

	close(emailChan)
	<-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)

	// <-done // block

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "ping"   // blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}
