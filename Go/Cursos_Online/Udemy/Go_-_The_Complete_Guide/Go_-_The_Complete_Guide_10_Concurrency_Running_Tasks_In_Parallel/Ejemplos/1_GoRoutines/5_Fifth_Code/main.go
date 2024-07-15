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
	close(doneChan) // We use the function "close" as we know the last function to produce its output is this one, but if we do not know for sure it should not be used as it terminates the channel and it can not be used to send data anymore
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you", done)
	go slowGreet("How... are ... you?", done)

	for doneChan := range done {
		fmt.Println(doneChan)
	}
}
