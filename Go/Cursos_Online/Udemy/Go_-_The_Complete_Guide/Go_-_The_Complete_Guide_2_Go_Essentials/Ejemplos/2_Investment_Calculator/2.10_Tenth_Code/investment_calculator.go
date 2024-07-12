package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 //We declare the const to a global state

func main() {
	var investAmout, years float64 = 1000, 10
	var expectedReturn float64

	outputText("Investment amount: ")
	fmt.Scan(&investAmout)
	fmt.Println("Expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculate(investAmout, expectedReturn, years) //We return and assign the return values of the custom function
	//To each one of the variables so that "fv" will be assigned to "futureValue" and "frv" will be assigned to "futureRealValue"

	formattedFV := fmt.Sprintf(`Future value: %.1f\n     
	asdaafas`, futureValue)
	formattedFRV := fmt.Sprintf("Future real value: %.0f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculate(investAmout, expectedReturn, years float64) (float64, float64) { //We define the number of variables that will be returned with the
	//Function using the "(float64, float64)" declaration

	fv := investAmout * math.Pow(1+expectedReturn/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv //We return both variable results
}
