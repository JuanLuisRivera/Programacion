package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum := sumup(numbers)
	sum2 := sumup2(1, 2, 3)
	sum3 := sumup3(1, 2, 3)
	sum4 := sumup4(1, 2, 3)
	sum5 := sumup2(numbers...) //We use the "..." operator in conjuntion to a slice to extract all the elements of the slice

	fmt.Println(sum)
	fmt.Println(sum2)
	fmt.Println(sum3)
	fmt.Println(sum4)
	fmt.Println(sum5)
}

func sumup(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}

func sumup2(numbers ...int) int { // We define a "variadic" function with the operator "..." that merges all the arguments passed to it in a slice
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}

func sumup3(notUsed int, numbers ...int) int { // We define a variadic function but it uses the arguments diferent, as now it leaves the first argument out of the slice it will create
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}

func sumup4(notUsed1 int, notUsed2 int, numbers ...int) int { // We define a variadic function but will now exclude the first two arguments of the sum
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}

// Note: We should note that the operator "..." groups elements of the same type into a slice with:  "...datatype" and extracts the data from a slice with: "slice_name..."
// And this operator can only be used as an argument of a function
