package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investAmout, years float64 = 1000, 10
	var expectedReturn float64

	fmt.Println("Investment amount: ")
	fmt.Scan(&investAmout)
	fmt.Println("Expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue := investAmout * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future values: ", futureValue)            //We print using a more readeable format using the "Println" function
	fmt.Printf("Future real value: %v", futureRealValue)   //We format the variable to appear at the place in the string we define
	fmt.Printf("Future real value: %0.f", futureRealValue) //We format to only output the integer part of the value
}
