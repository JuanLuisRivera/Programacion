package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2) // We create a function with our factory function
	triple := createTransformer(3) // We create a function that has the same logic as double but with only different argument

	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 }) // We define an "Anonymus" function as it has no name and so it can only be used here
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}

func createTransformer(factor int) func(int) int { //We use the function as a factory function as we can create different functions with different arguments
	return func(number int) int { // We use an anonymus function
		return number * factor // We use the "factor" variable as an input, so we can create different functions with the same logic
	} // This way of allowing the variable "factor" free to use is due to the fact that the anonymus function locked the value to only be used once everytime
	// It it called, this property is called "closure"
}
