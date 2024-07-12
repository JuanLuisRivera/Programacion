package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investAmout, years float64 = 1000, 10
	var expectedReturn float64

	outputText("Investment amount: ")
	fmt.Scan(&investAmout)
	fmt.Println("Expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculate(investAmout, expectedReturn, years)

	formattedFV := fmt.Sprintf(`Future value: %.1f\n     
	asdaafas`, futureValue)
	formattedFRV := fmt.Sprintf("Future real value: %.0f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculate(investAmout, expectedReturn, years float64) (fv float64, frv float64) { //We declare the variables "fv" and "frv"
	//And state that we will be returning specifically them

	fv = investAmout * math.Pow(1+expectedReturn/100, years) //We assign the values to the declared variables
	frv = fv / math.Pow(1+inflationRate/100, years)

	return //Go automatically returns the variables declared in the function definition
}
