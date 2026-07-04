package main

import "fmt"

// Send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {

	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result

	fmt.Println(res)

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
