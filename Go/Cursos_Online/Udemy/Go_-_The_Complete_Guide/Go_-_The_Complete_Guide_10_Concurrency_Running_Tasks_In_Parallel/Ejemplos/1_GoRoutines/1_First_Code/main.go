package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
}

func main() {
	go greet("Nice to meet you") // We use the reserved keyword "go" to specify golang that the codes should be used in a routine and therefore in parallel
	go greet("How are you")
	go slowGreet("How... are ... you?")
	go greet("Hope you like the course")
}
