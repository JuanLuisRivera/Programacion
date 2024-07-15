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
	done := make(chan bool)
	go greet("Nice to meet you", done) // We will use the same channel with different routines without any issues as the channel and the routine uses the same data type "bool"
	go slowGreet("How... are ... you?", done)
	<-done
	<-done // We use two result calls to display all the results from our functions
}

// Note: The routines will fall to a race condition as the first function to finish execution is the one that prints the result and finishes the program, so if we use one channel with all functions only one result will be displayed
// We could then include as many (<-Channel) calls as functions we have if we want them all to finish execution
