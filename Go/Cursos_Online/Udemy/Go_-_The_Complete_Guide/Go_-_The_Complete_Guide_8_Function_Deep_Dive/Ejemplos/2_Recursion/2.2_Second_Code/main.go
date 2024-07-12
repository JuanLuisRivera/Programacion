package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int { // Recursive approach factorial solution
	if number > 1 {
		return number * factorial(number-1) // We use the fact we can use functions as arguments
	} else {
		return 1
	}
}
