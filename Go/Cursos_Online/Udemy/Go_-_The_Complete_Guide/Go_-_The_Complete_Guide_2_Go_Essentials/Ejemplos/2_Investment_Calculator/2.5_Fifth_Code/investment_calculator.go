package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investAmout, years float64 = 1000, 10
	var expectedReturn float64 //We declare the variable without initializing it

	fmt.Println("Investment amount: ") //We add a line to specify the user the variable value and where to input
	fmt.Scan(&investAmout)             //It should be noted that the "Scan" function works most easily with one word values
	//So that using it with multiple word values is not recommended
	fmt.Println("Expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue := investAmout * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
