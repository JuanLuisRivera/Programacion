package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64
	revenue = getInput("Enter revenue: ")
	expenses = getInput("Enter expenses: ")
	taxRate = getInput("Enter tax rate: ")

	ebt, profit, ratio := performCalculation(revenue, expenses, taxRate)

	writeResultsToFile(ebt, profit, ratio)

	fmt.Println("EBT: ")
	fmt.Println(ebt)
	fmt.Println("Profit: ")
	fmt.Println(profit)
	fmt.Println("Ratio: ")
	fmt.Println(ratio)
}

func getInput(infoText string) (userInput float64) {
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		err := errors.New("Value not valid")
		println(err)
		panic("Error on input value")
	}
	return
}

func performCalculation(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func writeResultsToFile(ebt, profit, ratio float64) {
	resultsToString := fmt.Sprintf("EBT is: %f. Profit is: %f. And expenses are: %f.", ebt, profit, ratio) //We convert the data to a format string in order
	//To store the data to the file
	os.WriteFile("results.txt", []byte(resultsToString), 0644) //We create a file where to store the results
}
