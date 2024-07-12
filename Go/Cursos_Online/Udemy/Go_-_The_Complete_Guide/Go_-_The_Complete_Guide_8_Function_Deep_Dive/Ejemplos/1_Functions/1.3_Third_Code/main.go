package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 1, 9}
	doubled := transformNumbers(&numbers, double) // We pass our function "double" as an argument
	tripled := transformNumbers(&numbers, triple) // We pass our function "triple" as an argument
	//This way of passing functions in go makes it easy to not "hardcode" functions, avoiding code duplication

	transformFn1 := getTranformerFunction(&numbers)
	transformFn2 := getTranformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println("Original numbers: ", numbers)
	fmt.Println("Original more numbers: ", moreNumbers)
	fmt.Println("Doubled: ", doubled)
	fmt.Println("Tripled: ", tripled)
	fmt.Println("Tranformed numbers: ", transformedNumbers)
	fmt.Println("More tranformed numbers: ", moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int { // We modify the function to take a function as an argument, defining the function "transform" a function that takes 1 input "(int)" and produces 1 output "int"
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val)) // The "transform" variable is just a dummy variable that refers to a function with one input and one output, both of type "int"
	}
	return dNumbers
}

func getTranformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double // We return the function double as a result if the first number of the slice is 1
	} else {
		return triple // We return the function triple as a result if the first numer of the slice is not 1
	}

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
