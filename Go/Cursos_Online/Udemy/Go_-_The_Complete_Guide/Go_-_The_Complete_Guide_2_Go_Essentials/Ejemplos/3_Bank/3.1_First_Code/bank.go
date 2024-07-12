package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome")
	fmt.Println("What do you want to do?")
	fmt.Println("1.- Check Balance")
	fmt.Println("2.- Deposit Money")
	fmt.Println("3.- Withdraw Money")
	fmt.Println("4.- Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsCheckBalance := choice == 1 //We declare and assign the variable "wantsCheckBalance" true only if the variable "choice" equals "1"

	if wantsCheckBalance { //If true, executes the code
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 { //We use else if to assure that only one option will be considered
		var depositAmount float64

		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmount)

		accountBalance = accountBalance + depositAmount

		fmt.Println("Balance: ", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64

		fmt.Print("Amount to withdraw<: ")
		fmt.Scan(&withdrawAmount)

		accountBalance = accountBalance - withdrawAmount

		fmt.Println("Balance: ", accountBalance)
	}

	fmt.Println("Your choice: ", choice)
}
