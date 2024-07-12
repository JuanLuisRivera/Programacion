package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)

	fmt.Println(numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val)) // We pass the result of the function "double" as a parameter to other function
	}

	return dNumbers
}

func double(number int) int { // We define a function that doubles a value
	return number * 2
}
