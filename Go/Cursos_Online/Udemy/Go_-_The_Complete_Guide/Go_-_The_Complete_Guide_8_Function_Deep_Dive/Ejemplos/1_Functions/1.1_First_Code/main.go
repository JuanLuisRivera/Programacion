package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)

	fmt.Println(numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int { // We define a function that takes a pointer to a slice and returns another slice with the values doubled
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, val*2)
	}

	return dNumbers
}
