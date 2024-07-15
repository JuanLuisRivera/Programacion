package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true // We specify the direction the data should flow with the operator "<-" meaning the data should run from the bool "true" to the variable "doneChan"
}

func main() {
	done := make(chan bool) // We create a channel using the "chan" function, it should be noted that it is necessary to use the "make" function as well to initialize the channel and it is necessary to specify the data type the channel will be sending
	go slowGreet("How... are ... you?", done)
	<-done // We use the data obtained from the function "slowGreet", this helps specifying that the execution will finish only after we use the done data
}
