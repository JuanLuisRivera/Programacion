package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	dones := make([]chan bool, 4) // We create a slice for creating different channels for each routine

	dones[0] = make(chan bool)
	go greet("Nice to meet you", dones[0])

	dones[1] = make(chan bool)
	go slowGreet("How... are ... you?", dones[1])
	for _, done := range dones {
		<-done // We print all the data from the different channels
	}
}
