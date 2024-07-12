package main

import "fmt"

func main() {
	fact := factorial(6)
	fmt.Println(fact)
}

func factorial(number int) int { // Iterative approach factorial solution
	result := 1
	for i := 1; i < number; i++ {
		result = result * i
	}
	return result
}
