package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investAmout, years float64 = 1000, 10
	var expectedReturn float64

	outputText("Investment amount: ") //We use the defined function to print the text
	fmt.Scan(&investAmout)
	fmt.Println("Expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue := investAmout * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf(`Future value: %.1f\n     
	asdaafas`, futureValue)
	formattedFRV := fmt.Sprintf("Future real value: %.0f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) { //We define a custom function to print the variable "text" of type "string"
	fmt.Print(text)
}
