package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64 = 0.0, 0.0, 0.0 //We declare and initialize the variables

	fmt.Println("Revenue: ") //We ask the user for the inpiut values
	fmt.Scan(&revenue)
	fmt.Println("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("Tax rate")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses //We perform the appropiate calculations
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("EBT: ") //We print the result to console
	fmt.Println(ebt)
	fmt.Println("Profit: ")
	fmt.Println(profit)
	fmt.Println("Ratio: ")
	fmt.Println(ratio)
}
