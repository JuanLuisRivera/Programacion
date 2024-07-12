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

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue) //We use the "Sprintf" function to format the string as we please
	formattedFRV := fmt.Sprintf("Future real value: %.0f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}
