package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64 //We define the float variables
	revenue = getInput("Enter revenue: ")
	expenses = getInput("Enter expenses: ")
	taxRate = getInput("Enter tax rate: ")

	ebt, profit, ratio := performCalculation(revenue, expenses, taxRate) //We assign the results to the variables

	fmt.Println("EBT: ")
	fmt.Println(ebt)
	fmt.Println("Profit: ")
	fmt.Println(profit)
	fmt.Println("Ratio: ")
	fmt.Println(ratio)
}

func getInput(infoText string) (userInput float64) { //Function that prints a message and assigns user input value to a variable
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return
}

func performCalculation(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) { //Function that performs the calculations
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
